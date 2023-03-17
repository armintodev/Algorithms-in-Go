package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

var output1 float64
var output2 float64
var output3 float64
var output4 float64

func worker1() {
	defer waitGroup.Done()
	var output []float64
	sum := 0.0

	for index := 0; index < 100_000_000; index++ {
		output = append(output, 89.6)
		sum += 89.6
	}
	output1 = sum
}

func worker2() {
	defer waitGroup.Done()
	var output []float64
	sum := 0.0

	for index := 0; index < 100_000_000; index++ {
		output = append(output, 64.8)
		sum += 64.8
	}
	output2 = sum
}

func worker3() {
	defer waitGroup.Done()
	var output []float64
	sum := 0.0

	for index := 0; index < 100_000_000; index++ {
		output = append(output, 956.8)
		sum += 956.8
	}
	output3 = sum
}

func worker4() {
	defer waitGroup.Done()
	var output []float64
	sum := 0.0

	for index := 0; index < 100_000_000; index++ {
		output = append(output, 1235.8)
		sum += 1235.8
	}
	output4 = sum
}

func useBenchmarkingConcurrentApplications() {
	waitGroup.Add(8)

	//Compute time with no concurrent processing
	start := time.Now()

	worker1()
	worker2()
	worker3()
	worker4()

	elapsed := time.Since(start)

	fmt.Println("\nTime for 4 workers in series:", elapsed)
	fmt.Printf("Output1: %f \nOutput2: %f \nOutput3: %f \nOutput4: %f\n", output1, output2, output3, output4)

	//Compute time with concurrent processing
	start = time.Now()

	go worker1()
	go worker2()
	go worker3()
	go worker4()
	waitGroup.Wait()

	elapsed = time.Since(start)

	fmt.Println("\nTime for 4 workers in parallel:", elapsed)
	fmt.Printf("Output1: %f \nOutput2: %f \nOutput3: %f \nOutput4: %f\n", output1, output2, output3, output4)
}

const NumbersToSum = 10_000_000 // can be 15_000_000 or whatever

func sum(s []float64, c chan<- float64) {
	//A generator that puts data into channel
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	c <- sum // blocks until c is taken out of the channel
}

func plainSum(s []float64) float64 {
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	return sum
}

func useBenchmarkingNumbersToSum() {
	var s []float64
	for i := 0; i < NumbersToSum; i++ {
		s = append(s, 1.0)
	}

	c := make(chan float64)

	start := time.Now()

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	//go sum(s[:len(s)/2], c)
	//first, second, third := <-c, <-c, <-c // receive from each c
	first, second := <-c, <-c // receive from each c

	elapsed := time.Since(start)
	//fmt.Printf("first: %f  second: %f  third: %f elapsed time: %v", first, second, third, elapsed)
	fmt.Printf("first: %f  second: %f elapsed time: %v", first, second, elapsed)

	start = time.Now()

	answer := plainSum(s)

	elapsed = time.Since(start)
	fmt.Printf("\nplain sum: %f elapsed time: %v", answer, elapsed)
}
