package main

import (
	"sync/atomic"
)

type Counter struct {
	count int64
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.count, 1)
}

func (c *Counter) GetCount() int64 {
	return atomic.LoadInt64(&c.count)
}
