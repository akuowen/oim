package prpc

import (
	"context"
	"testing"

	"github.com/oim/common/config"

	"github.com/oim/common/prpc/example/helloservice"

	ptrace "github.com/oim/common/prpc/trace"
	"google.golang.org/grpc"
)

const (
	testIp   = "127.0.0.1"
	testPort = 8867
)

func TestNewPServer(t *testing.T) {
	config.Init("../../oim.yaml")

	ptrace.StartAgent()
	defer ptrace.StopAgent()

	s := NewPServer(WithServiceName("oim_server"), WithIP(testIp), WithPort(testPort), WithWeight(100))
	s.RegisterService(func(server *grpc.Server) {
		helloservice.RegisterGreeterServer(server, helloservice.HelloServer{})
	})
	s.Start(context.TODO())
}
