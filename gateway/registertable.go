package gateway

import "sync"

var table sync.Map

func init() {
	table = sync.Map{}
}
