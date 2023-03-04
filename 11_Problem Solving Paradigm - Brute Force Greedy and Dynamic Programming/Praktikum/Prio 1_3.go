package main

import "fmt"

var memo map[int]int

func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if _, ok := memo[n]; !ok {
		memo[n] = fibo(n-1) + fibo(n-2)
	}
	return memo[n]
}

func main() {
	memo = make(map[int]int)
	fmt.Println(fibo(0))  // 0
	fmt.Println(fibo(1))  // 1
	fmt.Println(fibo(2))  // 1
	fmt.Println(fibo(3))  // 2
	fmt.Println(fibo(5))  // 5
	fmt.Println(fibo(6))  // 8
	fmt.Println(fibo(7))  // 13
	fmt.Println(fibo(9))  // 34
	fmt.Println(fibo(10)) // 55
}
