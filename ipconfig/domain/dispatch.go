package domain

import "github.com/oim/ipconfig/source"

var dp *Dispatcher

func Init() {
	dp = &Dispatcher{}
	// 开携程 进行调度
	go func() {
		eventChan := source.EventChan()
		for event := range eventChan {
			switch event.Type {
			case source.AddNodeEvent:
				//调度器 更新的 操作
			case source.DelNodeEvent:
				//调度器 删除的操作
			}
		}
	}()

}

type Dispatcher struct {
}
