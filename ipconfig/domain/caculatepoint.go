package domain

import (
	"sync/atomic"
	"unsafe"
)

type Point struct {
	Ip          string `json:"ip"`
	Port        string `json:"port"`
	ActiveScore float64
	StaticScore float64
	Status      *Status
	window      *window
}

func (p *Point) CalculateScore(ctx *GetIpContext) {
	if p.Status != nil {
		p.ActiveScore = p.Status.CalculateActiveScore()
		p.StaticScore = p.Status.CalculateStaticScore()
	}
}

func (ed *Point) UpdateStat(s *Status) {
	ed.window.statusChan <- s
}

func NewPoint(ip, port string) *Point {
	point := Point{
		Ip:   ip,
		Port: port,
	}
	point.window = newStateWindow()
	point.Status = point.window.getStat()
	go func() {
		for status := range point.window.statusChan {
			point.window.appendStat(status)
			newStat := point.window.getStat()
			atomic.SwapPointer((*unsafe.Pointer)((unsafe.Pointer)(point.Status)), unsafe.Pointer(newStat))
		}
	}()
	return &point
}
