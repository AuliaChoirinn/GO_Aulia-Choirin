package main

import (
	"fmt"
)

func getPrime(n int) int {
	if n <= 0 {
		return -1
	}
	if n == 1 {
		return 2
	}

	count := 1
	current := 3

	for {
		if isPrime(current) {
			count++
			if count == n {
				return current
			}
		}
		current += 2
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeX(number int) int {
	return getPrime(number)
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
