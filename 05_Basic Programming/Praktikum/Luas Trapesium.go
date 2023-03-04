package main

import "fmt"

func main() {
	var a, b, t float64

	fmt.Print("Masukkan panjang sisi atas trapesium : ")
	fmt.Scan(&a)

	fmt.Print("Masukkan panjang sisi bawah trapesium : ")
	fmt.Scan(&b)

	fmt.Print("Masukkan tinggi trapesium : ")
	fmt.Scan(&t)

	luas := 0.5 * (a + b) * t

	fmt.Print("Luas trapesium adalah: ", luas)
}
