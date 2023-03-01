package counter

import (
	"fmt"
	"sync"
)

type Counter struct {
	text  string
	count int
	mutex sync.Mutex
}

func NewCounter(text string) *Counter {
	return &Counter{
		text:  text,
		count: 0,
	}
}

func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
	fmt.Printf("\033[%s %d", c.text, c.count)
}
