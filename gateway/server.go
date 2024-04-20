package gateway

import (
	"fmt"
	"github.com/oim/common/tcp"
	"golang.org/x/sys/unix"
	"log"
	"net"
	"runtime"
	"syscall"
	"time"

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

func runProc(c *conn, ep *EpollDesc) {

	dataLenBuf := make([]byte, 4)
	err := tcp.ReadDataWithTimeout(c.tcpConn, dataLenBuf, time.Duration(10))
	if err != nil {
		if err.Error() == "read timeout" {
			err := unix.EpollCtl(ep.fd, syscall.EPOLL_CTL_ADD, c.fd, &unix.EpollEvent{Events: unix.EPOLLIN | unix.EPOLLHUP, Fd: int32(c.fd)})
			if err != nil {
				return
			}
			return
		}
		if err.Error() == "connection closed" {
			ep.remove(c)
		}
	} else {
		// 处理读取到的数据
		err = WorkPool.Pool.Submit(func() {
			// TODO SEND RPC

		})
		if err != nil {
			_ = fmt.Errorf("runProc:err:%+v\n", err.Error())
		}
	}

}
