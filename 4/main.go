package main

import "fmt"

/**
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

solution = 906609
*/

func main() {
	fmt.Println(getLargestPalindrome())
}

func getLargestPalindrome() int {
	largestPalindrome := 0
	for i := 100; i < 999; i++ {
		for j := 100; j < 999; j++ {
			product := i * j
			if product > largestPalindrome && isPalindrome(product) {
				largestPalindrome = product
			}
		}
	}
	return largestPalindrome
}

func isPalindrome(n int) bool {
	temp := n
	reverse := 0
	for temp != 0 {
		remainder := temp % 10
		reverse = (reverse * 10) + remainder
		temp /= 10
	}

	return (n == reverse)
}
