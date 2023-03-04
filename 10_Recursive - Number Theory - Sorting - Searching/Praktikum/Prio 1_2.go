package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(item []string) []pair {
	counts := make(map[string]int)

	// hitung kemunculan setiap item
	for _, item := range item {
		counts[item]++
	}

	// buat slice dari pasangan item dan jumlah kemunculannya
	var pairs []pair
	for name, count := range counts {
		pairs = append(pairs, pair{name, count})
	}

	// urutkan slice berdasarkan jumlah kemunculan
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	return pairs
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "js", "js"}))
	// golang-> 1 ruby->1 js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// C-> 1 D->2 B->3 A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "golang", "tenis"}))
	// football-> 1 basketball->1 tenis->1
}
