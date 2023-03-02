package counter

import (
	"fmt"
	"sync"
)

type Counter struct {
	prefix string
	suffix string
	count  int
	mutex  sync.Mutex
}

func NewCounter(prefix string, suffix string) *Counter {
	return &Counter{
		prefix: prefix,
		suffix: suffix,
		count:  0,
	}
}

func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
	fmt.Printf("\r%s%d%s", c.prefix, c.count, c.suffix)
}

func (c *Counter) End() {
	fmt.Println()
}
