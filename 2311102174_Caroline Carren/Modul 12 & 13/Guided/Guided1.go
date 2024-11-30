// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan Selection Sort
func selectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		// Tukar elemen terkecil dengan elemen pada indeks i
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func main() {
	var n int

	// Meminta jumlah daerah kerabat
	fmt.Println("Masukkan jumlah daerah kerabat (n):")
	fmt.Scan(&n)

	// Validasi input jumlah daerah harus lebih besar dari 0
	if n <= 0 {
		fmt.Println("Jumlah daerah kerabat harus lebih besar dari 0.")
		return
	}

	// Loop untuk setiap daerah
	for i := 0; i < n; i++ {
		var m int

		// Meminta jumlah rumah kerabat di daerah ke-i
		fmt.Printf("Masukkan jumlah rumah kerabat di daerah ke-%d (m): ", i+1)
		fmt.Scan(&m)

		// Validasi input jumlah rumah kerabat harus lebih besar dari 0
		if m <= 0 {
			fmt.Printf("Jumlah rumah kerabat di daerah ke-%d harus lebih besar dari 0.\n", i+1)
			continue
		}

		// Membuat slice untuk menyimpan nomor rumah kerabat
		arr := make([]int, m)

		// Meminta input nomor rumah kerabat
		fmt.Printf("Masukkan %d nomor rumah kerabat untuk daerah ke-%d: ", m, i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		// Mengurutkan nomor rumah menggunakan Selection Sort
		selectionSort(arr, m)

		// Menampilkan hasil pengurutan
		fmt.Printf("Nomor rumah terurut untuk daerah ke-%d: ", i+1)
		for _, num := range arr {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}

	fmt.Println("Pengurutan selesai.")
}
