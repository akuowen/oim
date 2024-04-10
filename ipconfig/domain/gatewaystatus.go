package domain

type Status struct {
	ConnectNums float64
	MessageByte float64
}

func NewStatus(c, m float64) *Status {
	return &Status{
		ConnectNums: c,
		MessageByte: m,
	}
}

func (s *Status) Sub(status *Status) {
	if status == nil {
		return
	}
	s.ConnectNums -= status.ConnectNums
	s.MessageByte -= status.MessageByte
}

func (s *Status) Add(status *Status) {
	if status == nil {
		return
	}
	s.ConnectNums += status.ConnectNums
	s.MessageByte += status.MessageByte
}

func (s *Status) Clone() *Status {
	newStat := &Status{
		MessageByte: s.MessageByte,
		ConnectNums: s.ConnectNums,
	}
	return newStat
}

func (s *Status) Avg(num float64) {
	s.ConnectNums /= num
	s.MessageByte /= num
}
