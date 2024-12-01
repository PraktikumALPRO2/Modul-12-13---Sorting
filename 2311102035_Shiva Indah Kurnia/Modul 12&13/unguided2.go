package main

import (
	"fmt"
	"sort"
)

// Fungsi untuk menghitung median
func hitungMedian(arr []int) float64 {
	n := len(arr)
	if n == 0 {
		return 0
	}

	// Jika jumlah data ganjil
	if n%2 != 0 {
		return float64(arr[n/2])
	}

	// Jika jumlah data genap
	return float64(arr[n/2-1]+arr[n/2]) / 2.0
}

func main() {
	var data []int
	var input int

	// Membaca input sampai bertemu angka 0 atau -5313
	fmt.Println("Masukkan bilangan (0 untuk menghitung median, -5313 untuk mengakhiri):")
	for {
		fmt.Scan(&input)

		if input == 0 {
			// Urutkan data dan hitung median
			if len(data) > 0 {
				sort.Ints(data) // Urutkan data menggunakan `sort.Ints`
				fmt.Printf("%.0f\n", hitungMedian(data))
			}
		} else if input == -5313 {
			break
		} else {
			// Tambahkan input ke slice data
			data = append(data, input)
		}
	}
}
