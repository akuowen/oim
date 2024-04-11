package domain

import (
	"github.com/oim/ipconfig/source"
	"sort"
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

func (d *Dispatcher) getCandidatePoint(ctx *GetIpContext) []*Point {
	candidateList := make([]*Point, 0, len(dp.local))
	for _, ed := range dp.local {
		candidateList = append(candidateList, ed)
	}
	return candidateList
}

func Add(event *source.Event) {
	point, ok := dp.local[event.Key()]
	if !ok {
		newPoint := NewPoint(event.Ip, event.Port)
		dp.local[event.Key()] = newPoint
	}
	point.UpdateStat(NewStatus(event.ConnectNum, event.MessageBytes))

}

func Dispatch(ctx *GetIpContext) []*Point {
	points := dp.getCandidatePoint(ctx)
	for _, point := range points {
		point.CalculateScore(ctx)
	}
	// 排序
	sort.Slice(points, func(i, j int) bool {
		if points[i].ActiveScore > points[j].ActiveScore {
			return true
		}
		if points[i].ActiveScore == points[j].ActiveScore {
			if points[i].StaticScore > points[j].StaticScore {
				return true
			}
			return false
		}
		return false
	})
	return points
}
