package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan selection sort
func selectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			// Cari elemen terkecil
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		// Tukar elemen terkecil dengan elemen di posisi i
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	for daerah := 1; daerah <= n; daerah++ {
		var m int
		fmt.Printf("\nMasukkan jumlah nomor rumah kerabat untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat : ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}
		selectionSort(arr, m)

		fmt.Printf("Nomor rumah terurut untuk daerah %d : ", daerah)
		for _, num := range arr {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
