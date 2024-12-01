package main

import (
	"fmt"
)

// Fungsi untuk melakukan selection sort menaik
func SelectionSortMembesar(arr []int, n_158 int) {
	for i := 0; i < n_158-1; i++ {
		idxMin := i
		for j := i + 1; j < n_158; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

// Fungsi untuk melakukan selection sort menurun
func SelectionSortMengecil(arr []int, n_158 int) {
	for i := 0; i < n_158-1; i++ {
		idxMax := i
		for j := i + 1; j < n_158; j++ {
			if arr[j] > arr[idxMax] {
				idxMax = j
			}
		}
		arr[i], arr[idxMax] = arr[idxMax], arr[i]
	}
}

func main() {
	var n_158 int
	fmt.Print("n: ")
	fmt.Scan(&n_158)

	// List untuk menyimpan semua hasil terurut
	var hasil [][]int

	for i := 0; i < n_158; i++ {
		var m int
		fmt.Scan(&m)

		arr := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		if m%2 == 0 {
			SelectionSortMengecil(arr, m)
		} else {
			SelectionSortMembesar(arr, m)
		}

		hasil = append(hasil, arr)
	}

	// Cetak semua hasil terurut setelah input selesai
	fmt.Println("\nHasil Terurut :")
	for _, arr := range hasil {
		for _, num := range arr {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
