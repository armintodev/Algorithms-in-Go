package main

import "fmt"

type Ordered interface {
	~float64 | ~int | ~string
}

func bubbleSort[T Ordered](data []T) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func useBubbleSort() {
	numbers := []float64{3.5, -2.4, 12.8, 9.1}
	names := []string{"Armin", "Parsa", "Mohammad", "John", "Ali"}

	bubbleSort[float64](numbers)
	fmt.Println(numbers)

	bubbleSort[string](names)
	fmt.Println(names)
}
