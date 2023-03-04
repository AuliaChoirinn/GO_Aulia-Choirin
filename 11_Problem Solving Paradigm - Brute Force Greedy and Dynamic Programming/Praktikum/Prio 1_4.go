package main

import "fmt"

func SimpleEquations(a, b, c int) {
    found := false
    for x := -100; x <= 100 && !found; x++ {
        for y := -100; y <= 100 && !found; y++ {
            for z := -100; z <= 100 && !found; z++ {
                if x != y && y != z && x != z && x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
                    fmt.Println(x, y, z)
                    found = true
                }
            }
        }
    }
    if !found {
        fmt.Println("no solution")
    }
}

func main() {
    SimpleEquations(1, 2, 3)       // no solution
    SimpleEquations(6, 6, 14)      // 1 2 3
    SimpleEquations(138, 138, 297) // 6 9 23
}
