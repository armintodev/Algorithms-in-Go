package main

import (
	"fmt"
	"sync"
)

const (
	number = 1000
)

var countValue int

var mutex sync.Mutex

func useRaceCondition() {
	waitGroup.Add(number)

	for i := 0; i < number; i++ {
		go func() {
			countValue++
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	fmt.Printf("\n countValue = %d\n", countValue)
}

// using mutex to avoid race condition happen
func useMutexToAvoidRaceCondition() {
	waitGroup.Add(number)

	for i := 0; i < number; i++ {
		go func() {
			mutex.Lock()
			countValue++
			mutex.Unlock()
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	fmt.Printf("\n countValue = %d\n", countValue)
}
