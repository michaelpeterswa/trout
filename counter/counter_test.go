package counter_test

import (
	"time"

	"676f.dev/trout/counter"
	"github.com/fatih/color"
)

func ExampleCounter() {
	c := counter.NewCounter(color.HiCyanString("Hello")+"(", "x)")
	time.Sleep(1 * time.Second)
	c.Inc()
	time.Sleep(1 * time.Second)
	c.Inc()
	time.Sleep(1 * time.Second)
	c.Inc()
	time.Sleep(1 * time.Second)

	c.End()
}
