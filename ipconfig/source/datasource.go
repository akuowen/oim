package source

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/oim/common/config"
	"github.com/oim/common/discovery"
)

func Init() {
	eventChan = make(chan *Event)
	ctx := context.Background()
	go DataHandler(&ctx)
	if config.IsDebug() {
		//ctx := context.Background()
		//testServiceRegister(&ctx, "7896", "node1")
		//testServiceRegister(&ctx, "7897", "node2")
		//testServiceRegister(&ctx, "7898", "node3")
	}
}

// DataHandler 主要处理从存储中拿到的数据 然后转给计算
func DataHandler(ctx *context.Context) {
	serviceDiscovery := discovery.NewServiceDiscovery(ctx)
	defer serviceDiscovery.Close()

	serviceDiscovery.WatchService(config.GetServicePathForIPConf(), func(key, value string) {
		// 新增 写入计算chan  计算的逻辑监听这个 eventChan
		endpointInfoModel, err := discovery.UnMarshal([]byte(value))
		if err == nil {
			if endpointInfoModel != nil {
				event := NewEvent(endpointInfoModel)
				event.Type = AddNodeEvent
				eventChan <- event
			}
		} else {
			logger.CtxErrorf(*ctx, "DataHandler.setFunc.err :%s", err.Error())
		}
	}, func(key, value string) {
		// 删除
		endpointInfoModel, err := discovery.UnMarshal([]byte(value))
		if err == nil {
			if endpointInfoModel != nil {
				event := NewEvent(endpointInfoModel)
				event.Type = DelNodeEvent
				eventChan <- event
			}
		} else {
			logger.CtxErrorf(*ctx, "DataHandler.delFund.err :%s", err.Error())
		}
	})
}
