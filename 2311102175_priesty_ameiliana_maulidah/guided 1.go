package main

import (
	"fmt"
)

// fungsi untuk mengurutkan array menggunakan selection sort
func selectionsort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxmin := i
		for j := i + 1; j < n; j++ {
			// mencari elemen terkecil 
			if arr[j] < arr[idxmin] {
				idxmin = j
			}
		}
		// tukar elemen terkecil dengan elemen di posisi i
		arr[i], arr[idxmin] = arr[idxmin], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	// proses tiap daerah 
	for daerah := 1; daerah <= n; daerah++ {
		var m int // gunakan m untuk jumlah nomor rumah di daerah ini
		fmt.Printf("\nmasukkan jumlah nomor rumah kerabat untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		// membaca nomor rumah untuk daerah ini 
		arr := make([]int, m)
		fmt.Printf("masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}

		// urutkan array dari terkecil ke terbesar
		selectionsort(arr, m)

		// tampilkan hasil
		fmt.Printf("nomor rumah terurut untuk daerah %d: ", daerah)
		for _, num := range arr {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}