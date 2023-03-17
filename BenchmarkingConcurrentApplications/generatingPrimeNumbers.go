package main

import (
	"fmt"
	"math"
	"time"
)

const LargestPrime = 10_000_000

var primes []int

var cores int

func SieveOfEratosthenes(n int) []int {
	//Finds all primes up to 'n'
	primes := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		primes[i] = true
	}

	//The sieve logic for removing non-prime indices
	for p := 2; p*p <= n; p++ {
		if primes[p] == true {
			//Update all multiples of 'p'
			for i := p * 2; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	//return all prime numbers <= n
	var primeNumbers []int
	for p := 2; p <= n; p++ {
		if primes[p] == true {
			primeNumbers = append(primeNumbers, p)
		}
	}

	return primeNumbers
}

func useNonConcurrentSieveOfEratosthenes() {
	start := time.Now()

	sieve := SieveOfEratosthenes(LargestPrime)

	elapsed := time.Since(start)

	fmt.Println("\nComputation time: ", elapsed)
	fmt.Println("Largest prime: ", sieve[len(sieve)-1])
}

// Generate Send the sequence 3,5, ... to channel 'ch'.
func Generate(prime1 chan<- int) {
	for i := 3; ; i += 2 {
		prime1 <- i // Send 'i' to channel 'prime1'
	}
}

// Filter Copy the values from channel 'in' to channel 'out'
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'
		}
	}
}

func useConcurrentSieveOfEratosthenes() {
	start := time.Now()

	prime1 := make(chan int) // Create a new Channel

	go Generate(prime1) // Launch goroutine
	for {
		prime := <-prime1 // Take prime1 from out of channel
		if prime > LargestPrime {
			break
		}

		primes = append(primes, prime)
		prime2 := make(chan int)

		go Filter(prime1, prime2, prime)
		prime1 = prime2
	}

	elapsed := time.Since(start)
	fmt.Println("Computation time: ", elapsed)
	fmt.Println("Number of primes= ", len(primes))
}

func primesBetween(prime []int, low, high int) []int {
	//Computes the prime numbers between low and high
	//given the initial set of primes from the SieveOfEratosthenes
	limit := high - low
	var result []int
	segment := make([]bool, limit+1)

	for i := 0; i < len(segment); i++ {
		segment[i] = true
	}

	//Find the primes in the current segment based on initial primes
	for i := 0; i < len(prime); i++ {
		lowLimit := int(math.Floor(float64(low)/float64(prime[i])) * float64(prime[i]))

		if lowLimit < low {
			lowLimit += prime[i]
		}

		for j := lowLimit; j < high; j += prime[i] {
			segment[j-low] = false
		}
	}

	for i := low; i < high; i++ {
		if segment[i-low] == true {
			result = append(result, i)
		}
	}

	return result
}

func SegmentedSieve(n int) []int {
	//Each segment is of size square root of 'n'
	//Finds all primes up to 'n'
	var primeNumbers []int
	limit := (int)(math.Floor(math.Sqrt(float64(n))))

	prime := SieveOfEratosthenes(limit)
	for i := 0; i < len(prime); i++ {
		primeNumbers = append(primeNumbers, prime[i])
	}

	low := limit
	high := 2 * limit
	if high >= n {
		high = n
	}

	for {
		if low < n {
			next := primesBetween(prime, low, high)
			//fmt.Printf("\nprimesBetween(%d,%d) = %v", low, high, next)

			for i := 0; i < len(next); i++ {
				primeNumbers = append(primeNumbers, next[i])
			}

			low = low + limit
			high = high + limit
			if high >= n {
				high = n
			}
		} else {
			break
		}
	}

	return primeNumbers
}

func useNonConcurrentSegmentedSieveOfEratosthenes() {
	start := time.Now()

	primeNumbers := SegmentedSieve(LargestPrime)

	elapsed := time.Since(start)

	fmt.Println("\nComputation time: ", elapsed)
	fmt.Println("Number of primes: ", len(primeNumbers))
}
