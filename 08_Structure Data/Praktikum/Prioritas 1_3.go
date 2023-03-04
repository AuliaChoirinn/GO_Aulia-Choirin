package main

import (
	"fmt"
	"strconv"
	"strings"
)

func munculSekali(angka string) []int {
	// deklarasi variabel untuk menyimpan hasil
	var result []int

	// mengubah string angka menjadi slice of string
	splitAngka := strings.Split(angka, "")

	// deklarasi map untuk menghitung kemunculan setiap angka
	counter := make(map[string]int)
	for _, angka := range splitAngka {
		counter[angka]++
	}

	// memasukkan angka yang hanya muncul 1 kali ke dalam slice result
	for _, angka := range splitAngka {
		if counter[angka] == 1 {
			n, _ := strconv.Atoi(angka)
			result = append(result, n)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
