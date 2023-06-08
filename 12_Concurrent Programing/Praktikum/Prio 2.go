package main

import (
	"fmt"
	"strings"
	"sync"
)

func countFrequency(text string, freq map[rune]int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for _, r := range text {
		mu.Lock()
		freq[r]++
		mu.Unlock()
	}
}

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elite, sed do eiusmod tempor indidunt ut labore et dolore magna aliqua"
	text = strings.ToLower(text)

	var wg sync.WaitGroup
	var mu sync.Mutex
	freq := make(map[rune]int)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go countFrequency(text[i*len(text)/4:(i+1)*len(text)/4], freq, &wg, &mu)
	}

	wg.Wait()

	fmt.Println("Frekuensi huruf dalam teks:")
	for k, v := range freq {
		if k >= 'a' && k <= 'z' {
			fmt.Printf("%c: %d\n", k, v)
		}
	}
}
