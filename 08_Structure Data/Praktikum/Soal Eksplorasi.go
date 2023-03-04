package main

import (
	"fmt"
	"math"
)

func main() {
	// contoh matriks persegi 3x3
	matrix := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	var diag1, diag2 float64

	// hitung jumlah elemen pada diagonal utama dan diagonal kedua
	for i := 0; i < len(matrix); i++ {
		diag1 += matrix[i][i]
		diag2 += matrix[i][len(matrix)-i-1]
	}

	// hitung selisih absolut antara jumlah diagonal utama dan diagonal kedua
	diff := math.Abs(diag1 - diag2)

	fmt.Printf("Selisih absolut antara jumlah diagonal utama dan diagonal kedua: %.2f\n", diff)
}
