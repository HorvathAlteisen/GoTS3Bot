package TS3ServerQuery

import (
	"bufio"
	"net"
)

// ServerQuery stores TS3 ServerQuery connection details
type ServerQuery struct {
	conn net.Conn
	rbuf *bufio.Reader
	wbuf *bufio.Writer
}

// NewQuery Creates a new connection to a TeamSpeak3 Server
func NewQuery(address string) (query *ServerQuery, err error) {

	query = new(ServerQuery)
	query.conn, err = net.Dial("tcp", address)

	// If the connection fails the funtion returns an error
	if err != nil {

		return nil, err
	}

	query.rbuf = bufio.NewReader(query.conn)
	query.wbuf = bufio.NewWriter(query.conn)

	return query, nil
}

// Close the Query Connection
func (query *ServerQuery) Close() error {

	return query.conn.Close()
}

// SendCommand sends query commands to the TS3ServerQuery
func (query *ServerQuery) SendCommand(comand string) error {

	return nil
}

// Login with a user into an existent ServerQuery Session
func (query *ServerQuery) Login(user string, password string) {

	return
}
