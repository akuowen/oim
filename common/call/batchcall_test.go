package call

import (
	"log"
	_ "net/http/pprof"
	"testing"
	"time"
)

func TestBatchProcessor(t *testing.T) {
	//go http.ListenAndServe(":6060", nil)
	processor := NewBatchProcessor(10, time.Duration(1000*time.Millisecond), func(reqs []*Req) {
		for _, req := range reqs {
			log.Printf("test %+v \n", req)
		}
		log.Printf("================batch end================")

	})
	flag := 0
	for {
		time.Sleep(10 * time.Millisecond)
		req := &Req{
			Flag: flag,
		}
		processor.AddRequest(req)
		flag++
	}

}
