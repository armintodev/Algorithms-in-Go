package counter

type Counter struct {
	count int
}

// Methods

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Decrement() {
	if c.count > 0 {
		c.count--
	}
}

func (c *Counter) Reset() {
	c.count = 0
}

func (c Counter) GetCount() int {
	return c.count
}
