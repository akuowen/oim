package gateway

import (
	"context"
	"fmt"
	"github.com/oim/common/tcp"
	"log"
	"net"
	"runtime"

	"github.com/oim/common/config"
)

func init() {

}

func RunGateway(path string) {
	config.InitGateway(path)
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: config.GetGatewayTCPServerPort()})
	if err != nil {
		log.Printf("StartTCPEPollServer err:%s", err.Error())
		panic(err)
	}
	NewPool(runtime.NumCPU())
	InitEpoll(listener, runProc)
	select {}
}

func runProc(c *conn) {
	_ = context.Background() // 起始的contenxt
	// step1: 读取一个完整的消息包
	_, err := tcp.ReadData(c.tcpConn)
	err = WorkPool.Pool.Submit(func() {
		//TODO SEND RPC

	})
	if err != nil {
		fmt.Errorf("runProc:err:%+v\n", err.Error())
	}
}
