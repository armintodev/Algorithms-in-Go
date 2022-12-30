package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func outputStrings() {
	defer waitGroup.Done()
	var strings = [5]string{"One", "Two", "Three", "Four", "Five"}

	for i := 0; i < 5; i++ {
		delay := 1 + rand.Intn(3)
		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Println(strings[i])
	}
}

func outputIntegers() {
	defer waitGroup.Done()
	for i := 0; i < 5; i++ {
		delay := 1 + rand.Intn(3)
		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Println(i)
	}
}

func outputFloats() {
	defer waitGroup.Done()
	for i := 0; i < 5; i++ {
		delay := 1 + rand.Intn(3)
		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Println(float64(i*i) + 0.5)
	}
}

func useWaitGroup() {
	waitGroup.Add(3) //we have 3 delta,which of once per goroutine that use.

	go outputStrings()
	go outputIntegers()
	go outputFloats()

	waitGroup.Wait()
}
