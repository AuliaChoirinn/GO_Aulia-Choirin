package main

import (
    "fmt"
)

func findMinMax(numbers *[6]int) (int, int) {
    min := numbers[0]
    max := numbers[0]
    for _, number := range numbers {
        if number < min {
            min = number
        }
        if number > max {
            max = number
        }
    }
    return min, max
}

func main() {
    numbers := [6]int{4, 6, 2, 9, 1, 7}
    min, max := findMinMax(&numbers)
    fmt.Println("Numbers:", numbers)
    fmt.Println("Minimum value:", min)
    fmt.Println("Maximum value:", max)
}