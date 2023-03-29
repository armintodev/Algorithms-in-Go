package main

import "fmt"

type Counter struct {
	count int
}

// Methods

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Decrement() {
	c.count--
}

func (c *Counter) Reset() {
	c.count = 0
}

func (c Counter) GetCount() int {
	return c.count
}

func main() {
	myCounter := new(Counter)

	//myCounter.count = 100 // Defeats the encapsulation of Counter
	fmt.Println(myCounter.GetCount())

	for i := 1; i <= 10; i++ {
		myCounter.Increment()
	}
	myCounter.Decrement()
	//myCounter.count -= 6 // Defeats the encapsulation of Counter
	fmt.Println(myCounter.GetCount())
}
