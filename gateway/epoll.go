package gateway

import (
	"fmt"
	"golang.org/x/sys/unix"
	"log"
	"net"
	"runtime"
	"sync"
	"syscall"
)

var epoll *Epoll

type Epoll struct {
	connChan        chan *conn
	registerTable   sync.Map
	thSize          int
	connProcessSize int
	listener        *net.TCPListener
}

type epollDesc struct {
	id int64
	fd int
}

func NewEpollDesc() (*epollDesc, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		log.Fatalf("create epoll fail %s", err)
		return nil, err
	}
	return &epollDesc{
		id: 0,
		fd: fd,
	}, nil
}

func InitEpoll(listener *net.TCPListener) {
	epoll = NewEpoll(listener)
	epoll.handleConnEventProcess()
	epoll.start()
}

func NewEpoll(listener *net.TCPListener) *Epoll {

	return &Epoll{
		connChan:        make(chan *conn),
		registerTable:   sync.Map{},
		thSize:          runtime.NumCPU(),
		connProcessSize: runtime.NumCPU(),
		listener:        listener,
	}
}

func (c *Epoll) handleConnEventProcess() {
	for i := 0; i < c.connProcessSize; i++ {
		go func() {
			for {
				tcpconn, err := c.listener.AcceptTCP()
				if err != nil {
					continue
				}
				f, err := tcpconn.File()
				if err != nil {
					continue
				}
				newConn := NewConn(int(f.Fd()), tcpconn)
				c.addBizConnTask(newConn)
			}
		}()

	}
}

func (c *Epoll) addBizConnTask(newConn *conn) {
	c.connChan <- newConn
}

func (c *Epoll) start() {
	for i := 0; i < c.thSize; i++ {
		go c.startEpollHandle()
	}
}

func (c *Epoll) startEpollHandle() {
	desc, err := NewEpollDesc()
	if err != nil {
		panic(err)
	}

	//将当前conn 订阅 可读和 挂起事件
	go func() {
		for {
			select {

			case conn := <-c.connChan:
				if err := desc.addEpollTask(conn); err != nil {
					fmt.Printf("failed to add connection %v\n", err)
					conn.Close()
					continue
				}
			}
		}
	}()

	// 监听epoll 事件
	for {
		connections, err := desc.wait(200)
		if err != nil && err != syscall.EINTR {
			fmt.Printf("failed to epoll wait %v\n", err)
			continue
		}
		for _, conn := range connections {
			if conn == nil {
				continue
			}
			//WorkPool.pool.Submit()
		}
	}

}

func (e *epollDesc) wait(msec int) ([]*conn, error) {
	events := make([]unix.EpollEvent, 100)
	n, err := unix.EpollWait(e.fd, events, msec)
	if err != nil {
		return nil, err
	}
	var connections []*conn
	for i := 0; i < n; i++ {
		if connTmp, ok := epoll.registerTable.Load(int(events[i].Fd)); ok {
			connections = append(connections, connTmp.(*conn))
		}
	}
	return connections, nil
}

func (c *epollDesc) addEpollTask(co *conn) error {
	fd := co.fd
	err := unix.EpollCtl(c.fd, unix.EPOLL_CTL_ADD, fd, &unix.EpollEvent{
		Events: unix.EPOLLIN | unix.EPOLLHUP, // 监听可读事件
		Fd:     int32(fd),
	})
	if err != nil {
		log.Fatalf("%d call EpollCtl fail  %v\n", c.id, err)
		return err
	}
	epoll.registerTable.Store(fd, co)
	return nil
}
