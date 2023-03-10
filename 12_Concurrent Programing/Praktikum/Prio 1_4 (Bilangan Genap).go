package main

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
)

func printEvenNumbers() {
	for i := 2; i <= 100; i += 2 {
		mu.Lock()
		fmt.Printf("%d ", i)
		mu.Unlock()
	}
}

func main() {
	go printEvenNumbers()
	go printEvenNumbers()

	// menunggu kedua goroutine selesai
	var wg sync.WaitGroup
	wg.Add(2)
	wg.Wait()
}
