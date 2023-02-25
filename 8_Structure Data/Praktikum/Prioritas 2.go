package main

import "fmt"

func PairSum(arr []int, target int) []int {
	// inisialisasi dua indeks pointer, satu di awal dan satu di akhir array
	left, right := 0, len(arr)-1

	// lakukan perulangan selama indeks pointer kiri lebih kecil dari indeks pointer kanan
	for left < right {
		// hitung jumlah elemen pada indeks pointer kiri dan kanan
		sum := arr[left] + arr[right]

		// jika jumlah sama dengan target, kembalikan indeks dari dua angka
		if sum == target {
			return []int{left, right}
		}

		// jika jumlah lebih kecil dari target, pindahkan indeks pointer kiri ke kanan
		if sum < target {
			left++
		}

		// jika jumlah lebih besar dari target, pindahkan indeks pointer kanan ke kiri
		if sum > target {
			right--
		}
	}

	// jika tidak ada pasangan yang sesuai, kembalikan array kosong
	return []int{}
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}
