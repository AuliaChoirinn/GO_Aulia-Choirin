package main

import (
	"fmt"
	"strconv"
)

func getBinaryOneCount(n int) int {
	// convert integer to binary string
	binaryStr := strconv.FormatInt(int64(n), 2)

	// count number of ones in binary string
	count := 0
	for _, c := range binaryStr {
		if c == '1' {
			count++
		}
	}

	return count
}

func getBinaryOneCountArray(n int) []int {
	// initialize result array with length n+1
	ans := make([]int, n+1)

	// loop through all integers from 0 to n
	for i := 0; i <= n; i++ {
		ans[i] = getBinaryOneCount(i)
	}

	return ans
}

func main() {
	// input an integer
	var input int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&input)

	// get array of binary one counts
	binaryOneCountArray := getBinaryOneCountArray(input)

	// print the array
	fmt.Println(binaryOneCountArray)

	// input another integer
	var input2 int
	fmt.Print("Enter another integer: ")
	fmt.Scan(&input2)

	// get array of binary one counts for second input
	binaryOneCountArray2 := getBinaryOneCountArray(input2)

	// print the array for second input
	fmt.Println(binaryOneCountArray2)
}
