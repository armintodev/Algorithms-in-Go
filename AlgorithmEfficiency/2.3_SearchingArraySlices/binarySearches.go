package main

import (
	"fmt"
	"time"
)

func binarySearch[T Ordered](slice []T, target T) bool {
	low := 0
	high := len(slice) - 1

	for low <= high {
		median := (low + high) / 2

		if slice[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(slice) || slice[low] != target {
		return false
	}

	return true
}

func useBinarySearch() {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = float64(i) // is sorted
	}

	start := time.Now()

	result := binarySearch[float64](data, -10.0) // Should Return false

	elapsed := time.Since(start)

	fmt.Println("Time to search slice of 100_000_000 floats using binarySearch= ", elapsed)
	fmt.Println("Result of search is: ", result)

	start = time.Now()

	result = binarySearch[float64](data, float64(size/2))

	elapsed = time.Since(start)

	fmt.Println("Time to search slice of 100_000_000 floats using binarySearch= ", elapsed)
	fmt.Println("Result of search is: ", result)
}
