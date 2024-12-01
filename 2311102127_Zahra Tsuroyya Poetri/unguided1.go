package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan Selection Sort
func selectionSort(arr []int, ascending bool) {
	for i := 0; i < len(arr)-1; i++ {
		idx := i
		for j := i + 1; j < len(arr); j++ {
			// Cari elemen terkecil (untuk ascending) atau terbesar (untuk descending)
			if (ascending && arr[j] < arr[idx]) || (!ascending && arr[j] > arr[idx]) {
				idx = j
			}
		}
		// Tukar elemen terkecil/terbesar dengan elemen di posisi i
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	// Proses tiap daerah
	for daerah := 1; daerah <= n; daerah++ {
		var m int
		fmt.Printf("\nMasukkan jumlah nomor rumah kerabat untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		// Membaca nomor rumah untuk daerah ini
		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}

		// Pisahkan nomor ganjil dan genap
		ganjil := []int{}
		genap := []int{}
		for _, num := range arr {
			if num%2 != 0 {
				ganjil = append(ganjil, num)
			} else {
				genap = append(genap, num)
			}
		}

		// Urutkan nomor ganjil secara ascending dan nomor genap secara descending
		selectionSort(ganjil, true)  // Urutkan ganjil secara menaik (ascending)
		selectionSort(genap, false)   // Urutkan genap secara menurun (descending)

		// Tampilkan hasil
		fmt.Printf("Nomor rumah terurut untuk daerah %d: ", daerah)
		// Tampilkan nomor ganjil
		for _, num := range ganjil {
			fmt.Printf("%d ", num)
		}
		// Tampilkan nomor genap
		for _, num := range genap {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
