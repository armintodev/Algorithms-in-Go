package maincounter

import (
	"GenericDataStructuresAndAlgorithms/counter"
	"fmt"
)

func main() {
	myCounter := counter.Counter{}

	myCounter.Increment()
	myCounter.Increment()
	myCounter.Reset()
	myCounter.Increment()
	myCounter.Increment()
	myCounter.Increment()
	myCounter.Increment()
	myCounter.Decrement()
	countValue := myCounter.GetCount()
	fmt.Println(countValue)
}
