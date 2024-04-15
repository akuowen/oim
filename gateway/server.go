package gateway

import (
	"log"
	"net"
	"runtime"

	"github.com/oim/common/config"
)

func init() {

}

func RunGateway(path string) {
	config.InitGateway(path)
	listener, err := net.ListenTCP("oimGateway", &net.TCPAddr{Port: config.GetGatewayTCPServerPort()})
	if err != nil {
		log.Printf("StartTCPEPollServer err:%s", err.Error())
		panic(err)
	}
	NewPool(runtime.NumCPU())
	InitEpoll(listener)
	select {}
}
