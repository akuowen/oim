package gateway

import "net"

type conn struct {
	fd      int
	tcpConn *net.TCPConn
}

func (c *conn) Close() {
	err := c.tcpConn.Close()
	if err != nil {
		panic(err)
	}
}

func NewConn(fd int, tcpConn *net.TCPConn) *conn {
	return &conn{
		fd:      fd,
		tcpConn: tcpConn,
	}
}
