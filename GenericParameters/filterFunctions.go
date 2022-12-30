package main

import "fmt"

func MyFilter(input []float64, f func(float64) bool) []float64 {
	var result []float64
	for _, value := range input {
		if f(value) == true {
			result = append(result, value)
		}
	}

	return result
}

func MyGenericFilter[T any](input []T, f func(T) bool) []T {
	var result []T
	for _, value := range input {
		if f(value) {
			result = append(result, value)
		}
	}

	return result
}

func useFilter() {
	input := []float64{17.3, 11.1, 9.9, 4.3, 12.6}
	result := MyFilter(input, func(number float64) bool {
		if number <= 10.0 {
			return true
		}

		return false
	})

	fmt.Println(result)
}

func useGenericFilterAndMap() {
	var input = []float64{-5.0, -2.0, 4.0, 8.0}
	var result1 = MyGenericMap[float64, float64](input, func(n float64) float64 {
		return n * n
	})
	fmt.Println("Generic Map :", result1)

	greaterThanFive := MyGenericFilter[int]([]int{4, 6, 5, 2, 20, 1, 7},
		func(i int) bool {
			return i > 5
		})
	fmt.Println("Generic Filter - GreaterThanFive :", greaterThanFive)

	oddNumbers := MyGenericFilter[int]([]int{4, 6, 5, 2, 20, 1, 7},
		func(i int) bool {
			return i%2 == 1
		})
	fmt.Println("Generic Filter - OddNumbers :", oddNumbers)

	lengthGreaterThanThree := MyGenericFilter[string]([]string{"hello", "or", "the", "maybe"},
		func(s string) bool {
			return len(s) > 3
		})
	fmt.Println("Generic Filter - LengthGreaterThanThree :", lengthGreaterThanThree)
}
