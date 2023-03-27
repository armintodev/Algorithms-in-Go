package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const max = 5000

func Merge[T Ordered](left, right []T) []T {
	result := make([]T, len(left)+len(right))
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}

func MergeSort[T Ordered](data []T) []T {
	if len(data) > 100 {
		middle := len(data) / 2
		left := data[:middle]
		right := data[middle:]
		data = Merge(MergeSort(left), MergeSort(right))
	} else {
		InsertSort(data)
	}

	return data
}

func useNonConcurrentMergeSort() {
	data := make([]float64, size)

	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}

	/*data2:=make([]float64,size)
	copy(data2,data)*/

	start := time.Now()

	result := MergeSort[float64](data)

	elapsed := time.Since(start)

	fmt.Println("Elapsed time for non concurrent mergesort: ", elapsed)
	fmt.Println("Is Sorted: ", IsSorted(result))
}

func ConcurrentMergeSort[T Ordered](data []T) []T {
	if len(data) > 1 {
		if len(data) <= max {
			return MergeSort(data)
		} else {
			middle := len(data) / 2
			left := data[:middle]
			right := data[middle:]

			var waitGroup sync.WaitGroup
			waitGroup.Add(2)

			var data1, data2 []T
			go func() {
				defer waitGroup.Done()
				data1 = ConcurrentMergeSort(left)
			}()

			go func() {
				defer waitGroup.Done()
				data2 = ConcurrentMergeSort(right)
			}()

			waitGroup.Wait()
			return Merge(data1, data2)
		}
	}
	return nil
}

func useConcurrentMergeSort() {
	data := make([]float64, size)

	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}

	start := time.Now()

	result := ConcurrentMergeSort[float64](data)

	elapsed := time.Since(start)

	fmt.Println("Elapsed time for concurrent mergesort: ", elapsed)
	fmt.Println("Is Sorted: ", IsSorted(result))
}
