package counter

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
)

// Counter is a simple counter that can be incremented, decremented, or set to a specific value.
type Counter struct {
	prefix string
	suffix string
	count  int
	mutex  sync.Mutex
}

// NewDefaultCounter creates a counter with the default (and colored) prefix and suffix.
func NewDefaultCounter(task string) *Counter {
	return NewCounter(color.CyanString("Running %s ", task)+color.RedString("["), "x"+color.RedString("]"))
}

// NewCounter creates a counter with the given prefix and suffix.
func NewCounter(prefix string, suffix string) *Counter {
	return &Counter{
		prefix: prefix,
		suffix: suffix,
		count:  0,
	}
}

// Inc increments the counter by 1.
func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
	fmt.Printf("\r%s%d%s", c.prefix, c.count, c.suffix)
}

// Dec decrements the counter by 1.
func (c *Counter) Dec() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count--
	fmt.Printf("\r%s%d%s", c.prefix, c.count, c.suffix)
}

// To sets the counter to the given value.
func (c *Counter) To(n int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count = n
	fmt.Printf("\r%s%d%s", c.prefix, c.count, c.suffix)
}

// End prints a newline to the console. This should be called after the counter is no longer needed.
func (c *Counter) End() {
	fmt.Println()
}
