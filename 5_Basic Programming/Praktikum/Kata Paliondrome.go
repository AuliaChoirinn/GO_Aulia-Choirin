package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Scan("Input kata yang ingin dimasukkan: ")

	// Membaca inputan dari console
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Menghapus karakter whitespace dan mengubah huruf menjadi lowercase
	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))

	// Memeriksa apakah kata adalah palindrome atau tidak
	isPalindrome := true
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			isPalindrome = false
			break
		}
	}

	// Menampilkan hasil
	if isPalindrome {
		fmt.Scan("%s adalah palindrome\n", input)
	} else {
		fmt.Scan("%s bukan palindrome\n", input)
	}
}
