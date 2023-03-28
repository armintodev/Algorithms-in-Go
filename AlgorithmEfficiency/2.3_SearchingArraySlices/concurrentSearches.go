package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func searchSegment[T Ordered](slice []T, target T, a, b int, ch chan<- bool) {
	//Generates boolean value put into 'ch'
	for i := a; i < b; i++ {
		if slice[i] == target {
			ch <- true
		}
	}
	ch <- false
}

func concurrentSearch[T Ordered](data []T, target T) bool {
	ch := make(chan bool)
	numSegments := runtime.NumCPU()
	segmentSize := int(float64(len(data)) / float64(numSegments))

	//Launch 'numSegments' goroutine
	for index := 0; index < numSegments; index++ {
		go searchSegment(data, target, index*segmentSize, index*segmentSize+segmentSize, ch)
	}
	num := 0 //Completed goroutines
	for {
		select {
		case value := <-ch: // Blocks until a goroutine puts a bool into the channel
			if value == true {
				return true
			}
			num += 1
			if num == numSegments { // All goroutines have completed
				return false
			}
		}
	}
	return false
}

func useConcurrentSearch() {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}

	start := time.Now()

	result := concurrentSearch[float64](data, 54.0) // Should Return false

	elapsed := time.Since(start)

	fmt.Println("Time to search slice of 100_000_000 floats using concurrentSearch= ", elapsed)
	fmt.Println("Result of search is: ", result)

	start = time.Now()

	result = concurrentSearch[float64](data, data[size/2])

	elapsed = time.Since(start)

	fmt.Println("Time to search slice of 100_000_000 floats using concurrentSearch= ", elapsed)
	fmt.Println("Result of search is: ", result)
}
