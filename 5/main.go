package main

import (
	"fmt"
	"math"
)

/**
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

Plan? 
	Find the smallest set of primes, such that all numbers 1-20 can be constructued. This set will be the prime facorization of the smallest number
	to which 1-20 are all evenly divisible
*/

func main() {

	var lcm []int

	for i := 2; i <= 20; i++ {
		var primes = getPrimes(i)
		var difference = difference(lcm, primes)
		lcm = append(lcm, difference...)
	}

	var solution int = 1
	for _, item := range lcm {
		solution *= item
	}
	fmt.Println(solution)
}

func getPrimes(n int) []int {
	var primes []int
	for n%2 == 0 {
		n = n / 2
		primes = append(primes, 2)
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			primes = append(primes, i)
			n = n / i
		}
	}

	if n > 2 {
		primes = append(primes, n)
	}

	return primes
}

// elements in b that are not in a
func difference(a, b []int) (diff []int) {

	m := make(map[int]int)

	for _, item := range a {
		m[item] = m[item] + 1
	}

	for _, item := range b {
		if count := m[item]; count > 0 {
			m[item] = m[item] - 1
		} else {
			diff = append(diff, item)
		}
	}
	return
}
