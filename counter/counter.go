package counter

import "sync"

type Counter struct {
	sum int
	mu  *sync.Mutex
}

// NewCounter - конструктор Counter
func NewCounter() *Counter {
	return &Counter{mu: &sync.Mutex{}}
}

// Inc инкрементирует счетчик sum объекта Counter
func (c *Counter) Inc() {
	c.mu.Lock()
	c.sum++
	c.mu.Unlock()
}

// GetSum возвращает значение счетчика sum объекта Counter
func (c *Counter) GetSum() int {
	return c.sum
}
