package c14

import "sync"

func NewCounter() *Counter {
	return &Counter{}
}

type Counter struct {
	mu  sync.Mutex
	num int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num++
}

func (c *Counter) Value() int {
	return c.num
}
