package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

//const size = 100_000

const size = 50_000_000
const threshold = 5000

func quicksort[T Ordered](data []T, low, high int) {
	if low < high {
		var pivot = partition(data, low, high)
		quicksort(data, low, pivot)
		quicksort(data, pivot+1, high)
	}
}

func partition[T Ordered](data []T, low int, high int) int {
	//Pick a lowest bound element as a pivot value
	var pivot = data[low]

	var i = low
	var j = high

	for i < j {
		for data[i] <= pivot && i < high {
			i++
		}

		for data[j] > pivot && j > low {
			j--
		}

		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}

	data[low] = data[j]
	data[j] = pivot

	return j
}

func useQuickSort() {
	numbers := []float64{3.5, -2.4, 12.8, 9.1}
	names := []string{"Armin", "Parsa", "Mohammad", "John", "Ali"}

	quicksort[float64](numbers, 0, len(numbers)-1)
	fmt.Println(numbers)

	quicksort[string](names, 0, len(names)-1)
	fmt.Println(names)
}

func compareBubbleSortToQuickSort() {
	data := make([]float64, size)

	for i := 0; i < size; i++ {
		data[i] = math.Sin(float64(i * i))
	}

	start := time.Now()

	quicksort[float64](data, 0, len(data)-1)

	elapsed := time.Since(start)
	fmt.Println("Elapsed sort time for sine wave using quickSort: ", elapsed)

	start = time.Now()

	bubbleSort[float64](data)

	elapsed = time.Since(start)
	fmt.Println("Elapsed sort time for sine wave using bubbleSort: ", elapsed)
}

func InsertSort[T Ordered](data []T) {
	i := 1
	for i < len(data) {
		h := data[i]
		j := i - 1

		for j >= 0 && h < data[j] { // 'j' isn't always true because we reduce 'j' to 1 in step before
			data[j+1] = data[j]
			j -= 1
		}

		data[j+1] = h
		i += 1
	}
}

func Partition[T Ordered](data []T) int {
	data[len(data)/2], data[0] = data[0], data[len(data)/2]
	pivot := data[0]
	mid := 0
	i := 1
	for i < len(data) {
		if data[i] < pivot {
			mid += 1
			data[i], data[mid] = data[mid], data[i]
		}
		i += 1
	}

	data[0], data[mid] = data[mid], data[0]
	return mid
}

func IsSorted[T Ordered](data []T) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func concurrentQuickSort[T Ordered](data []T, waitGroup *sync.WaitGroup) {
	for len(data) >= 30 {
		mid := Partition(data)
		var portion []T

		if mid < len(data)/2 {
			portion = data[:mid]
			data = data[mid+1:]
		} else {
			portion = data[mid+1:]
			data = data[:mid]
		}
		if len(portion) > threshold {
			waitGroup.Add(1)
			go func(data []T) {
				defer waitGroup.Done()
				concurrentQuickSort(data, waitGroup)
			}(portion)
		} else {
			concurrentQuickSort(portion, waitGroup)
		}
	}

	InsertSort(data)
}

func QSort[T Ordered](data []T) {
	var waitGroup sync.WaitGroup
	concurrentQuickSort(data, &waitGroup)
	waitGroup.Wait()
}

func compareConcurrentAndRegularQuickSort() {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}

	data2 := make([]float64, size)
	copy(data2, data)

	start := time.Now()
	QSort[float64](data)
	elapsed := time.Since(start)

	fmt.Println("Elapsed time for concurrent quicksort: ", elapsed)
	fmt.Println("Is Sorted: ", IsSorted(data))

	start = time.Now()
	quicksort[float64](data2, 0, len(data2)-1)
	elapsed = time.Since(start)

	fmt.Println("Elapsed time for regular quicksort: ", elapsed)
	fmt.Println("Is Sorted: ", IsSorted(data2))
}
