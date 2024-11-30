package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n) // Membaca jumlah daerah kerabat

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m) // Membaca jumlah nomor rumah di daerah tersebut

		// Inisialisasi slice untuk menyimpan nomor rumah
		arr := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j]) // Memasukkan nomor rumah
		}

		// Pisahkan nomor rumah menjadi ganjil dan genap
		var ganjil, genap []int
		for _, num := range arr {
			if num%2 == 0 {
				genap = append(genap, num)
			} else {
				ganjil = append(ganjil, num)
			}
		}

		// Urutkan nomor rumah ganjil secara menaik
		sort.Ints(ganjil)

		// Urutkan nomor rumah genap secara menurun
		sort.Sort(sort.Reverse(sort.IntSlice(genap)))

		// Cetak hasil sesuai format
		for _, num := range ganjil {
			fmt.Printf("%d ", num)
		}
		for _, num := range genap {
			fmt.Printf("%d ", num)
		}
		fmt.Println() // Baris baru untuk tiap daerah
	}
}
