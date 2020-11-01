package main

import (
	"fmt"
	"math"
)

/**
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

solution = 6857
*/
func main() {
	n := 600851475143
	var largestPrimeFactor int

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		for n%i == 0 {
			if i > largestPrimeFactor {
				largestPrimeFactor = i
			}
			n = n / i
		}
	}

	if n > 2 && n > largestPrimeFactor {
		largestPrimeFactor = n
	}

	fmt.Println(largestPrimeFactor)
}
