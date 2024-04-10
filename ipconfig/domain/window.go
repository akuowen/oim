package domain

var windowNum int64

func init() {
	windowNum = 5
}

type window struct {
	stateQueue []*Status
	statusChan chan *Status
	sumStat    *Status
	idx        int64
}

func newStateWindow() *window {
	return &window{
		stateQueue: make([]*Status, windowNum),
		statusChan: make(chan *Status),
		sumStat:    &Status{},
	}
}

func (sw *window) getStat() *Status {
	res := sw.sumStat.Clone()
	res.Avg(float64(windowNum))
	return res
}

func (sw *window) appendStat(s *Status) {
	// 减去即将被删除的state
	sw.sumStat.Sub(sw.stateQueue[sw.idx%windowNum])
	// 更新最新的stat
	sw.stateQueue[sw.idx%windowNum] = s
	// 计算最新的窗口和
	sw.sumStat.Add(s)
	sw.idx++
}
