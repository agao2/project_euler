package main

import "fmt"

/**
Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

func main() {
	
	// just brute force, it'll be O(n) and runs fast enough for this problem

	n := 100
	sum := 0
	squaredSum := 0

	for i := 1; i <= n; i++ {
		sum += i
		squaredSum += (i * i)
	}

	solution := (sum * sum) - squaredSum
	fmt.Println(solution)
}
