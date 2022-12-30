package main

import "fmt"

func MyMap(input []int, f func(int) int) []int {
	result := make([]int, len(input))
	for index, value := range input {
		result[index] = f(value)
	}

	return result
}

func MyGenericMap[T1, T2 any](input []T1, f func(T1) T2) []T2 {
	var result = make([]T2, len(input))
	for index, value := range input {
		result[index] = f(value)
	}

	return result
}

func useMap() {
	slice := []int{1, 5, 2, 7, 4}
	result := MyMap(slice, func(i int) int {
		return i * i
	})

	fmt.Println(result)
}

func useGenericMap() {
	slice := []int{1, 5, 2, 7, 4}
	result := MyGenericMap(slice, func(i int) int {
		return i * i
	})

	fmt.Println(result)
}
