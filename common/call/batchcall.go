package call

import (
	"sync"
	"time"
)

type Req struct {
}

type CallProcess[T any] interface {
	AddRequest(T *Req)
	Close()
}

var _ CallProcess[Req] = &BatchProcessor[Req]{}

type BatchProcessor[T Req] struct {
	mu          sync.Mutex
	requests    []*T
	batchSize   int
	timeout     time.Duration
	dataHandler func([]*T)
	done        chan *T
}

func NewBatchProcessor(batchSize int, timeout time.Duration, dataHandler func([]*Req)) *BatchProcessor[Req] {
	bp := &BatchProcessor[Req]{
		requests:    make([]*Req, 0, batchSize),
		timeout:     timeout,
		dataHandler: dataHandler,
		done:        make(chan *Req),
	}
	go bp.processLoop()

	return bp
}

func (bp *BatchProcessor[Req]) AddRequest(req *Req) {
	bp.mu.Lock()
	defer bp.mu.Unlock()
	bp.requests = append(bp.requests, req)

	if len(bp.requests) >= bp.batchSize {
		bp.sendRequests()
	}
}

func (bp *BatchProcessor[Req]) Close() {
	bp.done <- &Req{}
}

func (bp *BatchProcessor[Req]) processLoop() {
	ticker := time.NewTicker(bp.timeout)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			bp.sendRequests()
		case <-bp.done:
			return
		}
	}
}

func (bp *BatchProcessor[Req]) sendRequests() {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	if len(bp.requests) == 0 {
		return
	}
	bp.dataHandler(bp.requests)
	bp.requests = make([]*Req, 0, bp.batchSize)
}
