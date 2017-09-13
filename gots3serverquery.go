package GoTS3ServerQuery

import (
	"net"
)

type ServerQuery struct {
	conn	net.Conn
}

func NewQuery(address string) (serverQuery *ServerQuery, err error) {

	serverQuery = new(ServerQuery)
	serverQuery.conn = net.Dial("tcp", address)



	return 
}