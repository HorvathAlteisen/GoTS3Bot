package GoTS3ServerQuery

import (
	"net"
)

type ServerQuery struct {
	conn net.Conn
}

// Creates a new connection to a TeamSpeak3 Server
func NewQuery(address string) (serverQuery *ServerQuery, err error) {

	serverQuery = new(ServerQuery)
	serverQuery.conn, err = net.Dial("tcp", address)

	return
}

/*
func (query *ServerQuery) Write() (

	return
)*/

/* Login with a user into an existent ServerQuery
func (query *ServerQuery) Login(user string, password string) {

	query.conn.Write

	return
}*/
