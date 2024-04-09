package ipconfig

import (
	"github.com/oim/common/config"
	"github.com/oim/ipconfig/source"
)

func RunIpConfigServer(path string) {
	config.InitConfig(path)
	//链接存储
	source.Init()
	//启动调度

	//开启http服务
}
