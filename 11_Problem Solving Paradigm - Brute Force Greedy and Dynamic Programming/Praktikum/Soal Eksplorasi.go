package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	// slice of possible Roman numeral symbols
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	// slice of corresponding integer values
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	// initialize result string
	result := strings.Builder{}

	// loop through the symbols and values slices
	for i := 0; i < len(symbols); i++ {
		// repeat the current symbol until the value is less than the corresponding integer value
		for num >= values[i] {
			result.WriteString(symbols[i])
			num -= values[i]
		}
	}

	// return the Roman numeral string
	return result.String()
}

func main() {
	// input an integer to convert to Roman numerals
	var input int
	fmt.Print("Enter an integer to convert to Roman numerals: ")
	fmt.Scan(&input)

	// convert the input integer to Roman numerals
	romanNumeral := intToRoman(input)

	// print the Roman numeral string
	fmt.Println(romanNumeral)
}
