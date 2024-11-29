package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n) // Membaca jumlah daerah

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m) // Membaca jumlah nomor rumah di daerah ini

		oddNumbers := []int{}
		evenNumbers := []int{}

		for j := 0; j < m; j++ {
			var houseNumber int
			fmt.Scan(&houseNumber) // Membaca nomor rumah
			if houseNumber%2 == 0 {
				evenNumbers = append(evenNumbers, houseNumber) // Jika genap
			} else {
				oddNumbers = append(oddNumbers, houseNumber) // Jika ganjil
			}
		}

		// Mengurutkan nomor rumah ganjil secara menaik
		sort.Ints(oddNumbers)

		// Mengurutkan nomor rumah genap secara menurun
		sort.Sort(sort.Reverse(sort.IntSlice(evenNumbers)))

		// Menampilkan hasil untuk nomor ganjil
		for _, num := range oddNumbers {
			fmt.Print(num, " ")
		}
		// Menampilkan hasil untuk nomor genap
		for _, num := range evenNumbers {
			fmt.Print(num, " ")
		}
		fmt.Println() // Baris baru setelah setiap daerah
	}
}