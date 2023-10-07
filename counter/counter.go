package counter

import "sync"

type Counter struct {
	sum int
	mu  sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{mu: sync.Mutex{}}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.sum++
	c.mu.Unlock()
}

func (c *Counter) GetSum() int {
	return c.sum
}
