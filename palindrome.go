package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}

func largestPalindromeProduct() (int, int, int) {
	largestPalindrome := 0
	var mul1, mul2 int
	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			product := i * j
			if product < largestPalindrome {
				break
			}
			if isPalindrome(product) && product > largestPalindrome {
				largestPalindrome = product
				mul1, mul2 = i, j
			}
		}
	}

	return largestPalindrome, mul1, mul2
}

func main() {
	result, mul1, mul2 := largestPalindromeProduct()
	fmt.Printf("Largest palindrome number is: %d\n", result)
	fmt.Printf("Multiplicands are %d and %d\n", mul1, mul2)
}
