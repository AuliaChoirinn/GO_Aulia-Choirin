package main

import (
	"fmt"
)

func printMultiplesOfThree(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i * 3
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go printMultiplesOfThree(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
