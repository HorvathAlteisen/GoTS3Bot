package ts3

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

// ServerQuery stores TS3 ServerQuery connection details
type ServerQuery struct {
	conn             net.Conn
	scan             *bufio.Scanner
	cLine            chan string
	cNotify          chan string
	cErr             chan string
	res              string
	notifyHandler    func(string, string)
	connErrorHandler func(error)
	WelcomeMsg       string
}

type Error struct {
	id  int
	msg string
}

// Default Telnet Port of a TS3 Server Query
const (
	defaultIP    = "127.0.0.1"
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

	// If the connection fails the funtion returns an error
	query.conn, err = net.DialTimeout("tcp", address, DialTimeout)
	if err != nil {
		return nil, err
	}

	query.cLine = make(chan string)

	query.scan = bufio.NewScanner(query.conn)
	query.scan.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		/*for b := range data {

			fmt.Println(data[b])

		}
		fmt.Println("End of Data")*/

		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.Index(data, []byte("\n\r")); i >= 0 {
			return i + 2, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	go func() {
		for {
			if !query.scan.Scan() && query.connErrorHandler != nil {
				err = query.scan.Err()
				if err != nil {
					query.connErrorHandler(err)
				} else {
					query.connErrorHandler(errors.New("EOF"))
				}
			}
			query.cLine <- query.scan.Text()
		}
	}()

	// Checks if the first line contains the signature "TS3" of the ServerQuery
	if !strings.Contains(<-query.cLine, verifyString) {
		query.Close()
		return nil, errors.New("This is not a TS3 Server")
	}

	// Get Welcome Msg, dunno I want to show it on my project
	query.WelcomeMsg = <-query.cLine

	query.cErr = make(chan string)
	query.cNotify = make(chan string)

	// Notify Handler Loop

	go func() {
		for {
			notification := <-query.cNotify
			fmt.Println(notification) /*
				if query.notifyHandler != nil {
					client.notifyHandler(ParseNotification(notification))
				} else if client.notifyHandlerString != nil {
					client.notifyHandlerString(notification)
				} else {
					// discard
				}*/
		}
	}()

	// Line Handler
	go func() {

		for {
			line := <-query.cLine

			if strings.Index(line, "error") == 0 {
				query.cErr <- line
			} else if strings.Index(line, "notify") == 0 {
				query.cNotify <- line
			} else {
				query.res = line
			}
		}

	}()

	return query, nil
}

// Close the Query Connection
func (query *ServerQuery) Close() error {
	query.SendCommand("quit")

	return query.conn.Close()
}

// SendCommand sends query commands to the TS3ServerQuery
func (query *ServerQuery) SendCommand(command string) (string, string) {
	fmt.Fprintf(query.conn, "%s\n\r", command)
	err := <-query.cErr
	res := query.res
	query.res = ""

	return res, err
}
