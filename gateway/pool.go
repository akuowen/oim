package gateway

import (
	"fmt"
	"github.com/panjf2000/ants"
)

var WorkPool *bizPool

type bizPool struct {
	pool *ants.Pool
}

func NewPool(size int) {
	pool, err := ants.NewPool(size)
	if err != nil {
		fmt.Printf("InitWorkPoll.err :%s num:%d\n", err.Error(), size)

	}
	WorkPool = &bizPool{
		pool: pool,
	}

}
