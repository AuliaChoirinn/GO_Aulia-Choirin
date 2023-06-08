package main

import (
    "fmt"
    "strings"
)

func Compare(a, b string) string {
    // Find the shorter string
    var short, long string
    if len(a) < len(b) {
        short, long = a, b
    } else {
        short, long = b, a
    }

    // Check for common substrings of decreasing length
    for i := len(short); i >= 1; i-- {
        for j := 0; j <= len(short)-i; j++ {
            if strings.Contains(long, short[j:j+i]) {
                return short[j : j+i]
            }
        }
    }

    // If no common substring is found, return an empty string
    return ""
}

func main() {
    fmt.Println(Compare("AKA", "AKASHI"))    // AKA
    fmt.Println(Compare("KANGOORO", "KANG")) // KANG
    fmt.Println(Compare("KI", "KIJANG"))     // KI
    fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
    fmt.Println(Compare("ILALANG", "ILA"))   // ILA
}

