package main

import (
	"fmt"
	"time"
)

func printMultiplesOfX(x, n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", i*x)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	go printMultiplesOfX(5, 10)

	// menunggu goroutine selesai
	time.Sleep(3 * time.Second)
}