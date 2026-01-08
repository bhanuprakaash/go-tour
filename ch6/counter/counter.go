package counter

type Counter struct {
	value int
}

func New(start int) *Counter {
	if start < 0 {
		start = 0
	}
	return &Counter{value: start}
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Reset() {
	c.value = 0
}
