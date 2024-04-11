package domain

import "math"

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

func decimal(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}

func getGB(m float64) float64 {
	return decimal(m / (1 << 30))
}

func (s *Status) CalculateActiveScore() float64 {
	return getGB(s.MessageByte)
}
func (s *Status) CalculateStaticScore() float64 {
	return s.ConnectNums
}
