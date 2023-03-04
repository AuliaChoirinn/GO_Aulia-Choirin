package main

import (
	"fmt"
	"strings"
)

func caesar(offset int, input string) string {
	// handle offset yang lebih besar dari 26
	offset = offset % 26

	// handle input string yang kosong
	if input == "" {
		return ""
	}

	var result strings.Builder

	for _, ch := range input {
		if ch == ' ' {
			result.WriteRune(' ')
		} else if ch >= 'a' && ch <= 'z' {
			newCh := ch + rune(offset)
			if newCh > 'z' {
				newCh = newCh - 26
			}
			result.WriteRune(newCh)
		}
	}

	return result.String()
}

func main() {
	fmt.Println(caesar(3, "abc")) // def
	fmt.Println(caesar(2, "alta")) // cnvc
	fmt.Println(caesar(10, "alterraacademy")) // kvdabbkkmknowy
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
