package ipconfig

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/oim/common/config"
	"github.com/oim/ipconfig/domain"
	"github.com/oim/ipconfig/source"
)

func RunIpConfigServer(path string) {
	config.InitConfig(path)
	//链接存储
	source.Init()
	//启动调度
	domain.Init()
	//开启http服务
	s := server.Default(server.WithHostPorts(":6789"))
	s.GET("/ip/list", getIp)
	s.Spin()
}
