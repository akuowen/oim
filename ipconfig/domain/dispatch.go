package domain

import (
	"github.com/oim/ipconfig/source"
)

var dp *Dispatcher

func Init() {
	dp = &Dispatcher{}
	// 开协程 进行调度
	go func() {
		eventChan := source.EventChan()
		for event := range eventChan {
			switch event.Type {
			case source.AddNodeEvent:
				//调度器 更新的 操作
				Add(event)
			case source.DelNodeEvent:
				//调度器 删除的操作
			}
		}
	}()

}

type Dispatcher struct {
	local map[string]*Point
}

func Add(event *source.Event) {
	point, ok := dp.local[event.Key()]
	if !ok {
		newPoint := NewPoint(event.Ip, event.Port)
		dp.local[event.Key()] = newPoint
	}
	point.UpdateStat(NewStatus(event.ConnectNum, event.MessageBytes))

}
