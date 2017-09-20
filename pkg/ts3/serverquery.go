package ts3

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net"
	"strings"
	"time"
)

// ServerQuery stores TS3 ServerQuery connection details
type ServerQuery struct {
	conn    net.Conn
	cResult chan string
	rbuf    *bufio.Reader
	wbuf    *bufio.Writer
	out     io.Writer
}

// Default Telnet Port of a TS3 Server Query
const (
	defaultPort  = "10011"
	verifyString = "TS3"
)

var (
	DialTimeout    = 1 * time.Second
	CommandTimeout = 350 * time.Millisecond
)

// NewQuery Creates a new connection to a TeamSpeak3 Server
func NewQuery(address string) (query *ServerQuery, err error) {

	query = new(ServerQuery)

	if !strings.Contains(address, ":") {
		address += ":" + defaultPort
	}

	query.conn, err = net.DialTimeout("tcp", address, DialTimeout)
	query.rbuf = bufio.NewReader(query.conn)
	query.wbuf = bufio.NewWriter(query.conn)
	query.cResult = make(chan string)

	// If the connection fails the funtion returns an error
	if err != nil {
		return nil, err
	}

	line, _ := query.rbuf.ReadString('\n')

	// Checks if the first line contains the signatur "TS3" of the ServerQuery
	if !strings.Contains(line, verifyString) {

		return nil, errors.New("This is not a TS3 Server")
	}

	line, err = query.rbuf.ReadString('\n')

	go query.sync()

	return query, nil
}

// Close the Query Connection
func (query *ServerQuery) Close() error {
	query.SendCommand("quit")

	return query.conn.Close()
}

// SendCommand sends query commands to the TS3ServerQuery
func (query *ServerQuery) SendCommand(command string) (string, error) {
	command += "\n"

	_, err := query.send(command)

	return <-query.cResult, err
}

func (query *ServerQuery) send(p string) (int, error) {
	b := []byte(p)
	// Double IAC chars
	bytes.Replace(b, []byte{0xff}, []byte{0xff, 0xff}, -1)
	return query.conn.Write(b)
}

func (query *ServerQuery) Write(p []byte) (int, error) {
	s := string(p)
	s = strings.Replace(s, "\r", "", -1)

	query.responseHandler(s)

	//query.handleResponse(s)
	return len(s), nil
}

func (query *ServerQuery) responseHandler(data string) {

	query.cResult <- data
	//lines := strings.Split(data, "\n")
	//var result string
}

// cp copies from an io.Reader to an io.Writer
func (query *ServerQuery) sync() {
	for {
		io.Copy(query, query.conn)
	}
}
