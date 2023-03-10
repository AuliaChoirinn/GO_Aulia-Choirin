package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

func factorial(n int, wg *sync.WaitGroup, result *int) {
	defer wg.Done()

	if n == 0 {
		*result = 1
		return
	}

	mutex.Lock()
	for i := n; i > 0; i-- {
		*result *= i
	}
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	result := 1

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go factorial(i, &wg, &result)
	}

	wg.Wait()
	fmt.Printf("Hasil faktorial adalah %d\n", result)
}
