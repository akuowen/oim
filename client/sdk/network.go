package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/oim/common/tcp"
	"net"
)

type connect struct {
	serverAddr string
	sendChan   chan *MessageBody
	revChan    chan *MessageBody
	con        *net.TCPConn
}

func newConnect(serverAddr string) *connect {
	clientConn := &connect{
		serverAddr: serverAddr,
		sendChan:   make(chan *MessageBody),
		revChan:    make(chan *MessageBody),
	}
	tcpAddr := &net.TCPAddr{IP: net.ParseIP(serverAddr), Port: 8900}

	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Printf("DialTCP.err=%+v", err)
		return nil
	}
	clientConn.con = tcpConn

	go func() {
		for {

			data, err := tcp.ReadData(tcpConn)
			if err != nil {
				fmt.Printf("ReadData err:%+v", err)
			}
			msg := &MessageBody{}
			json.Unmarshal(data, msg)
			clientConn.revChan <- msg
		}
	}()
	return clientConn
}

func (c *connect) sendMessage(body *MessageBody) {
	c.revChan <- body
}

func (c *connect) rev() <-chan *MessageBody {
	return c.revChan
}

func (c *connect) close() {
	// 目前没啥值得回收的
	err := c.con.Close()
	if err != nil {
		return
	}
}
