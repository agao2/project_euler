package main

import (
	"fmt"
	"math"
)

/**
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10001st prime number?
*/

func main() {

	primeIndex := 1
	prime := 2

	for i := 1; primeIndex < 10001; i += 2 {
		if isPrime(i) {
			primeIndex++
			prime = i
		}
	}

	fmt.Println(prime)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
