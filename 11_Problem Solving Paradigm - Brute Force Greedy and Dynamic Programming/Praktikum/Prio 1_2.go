package main

import "fmt"

func generatePascalTriangle(numRows int) [][]int {
	triangle := make([][]int, numRows)

	for i := range triangle {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

func main() {
	n := 5 // jumlah baris yang diinginkan
	pascalTriangle := generatePascalTriangle(n)

	// print segitiga pascal
	for i := range pascalTriangle {
		fmt.Println(pascalTriangle[i])
	}

	// print baris pertama
	fmt.Println("Baris pertama segitiga pascal:", pascalTriangle[0])
}
