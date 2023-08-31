package main

import (
	"fmt"
	"math"
)

func largestPrimeFactor(n int) int {
	largest := 2
	for n%2 == 0 {
		n /= 2
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			largest = i
			n /= i
		}
	}
	if n > 2 {
		largest = n
	}
	return largest
}

func main() {
	var n int
	fmt.Scan(&n)
	if n <= 1 {
		fmt.Println("It is not a prime number.")
	} else {
		fmt.Printf("The largest prime factor of %d is %d.\n", n, largestPrimeFactor(n))
	}
}
