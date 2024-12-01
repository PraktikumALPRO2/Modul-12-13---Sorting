// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int

	// Meminta jumlah daerah kerabat
	fmt.Println("Masukkan jumlah daerah kerabat (n):")
	fmt.Scan(&n)

	// Validasi input jumlah daerah kerabat harus lebih besar dari 0
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
		houses := make([]int, m)

		// Meminta input nomor rumah kerabat
		fmt.Printf("Masukkan %d nomor rumah kerabat untuk daerah ke-%d: ", m, i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}

		// Mengurutkan nomor rumah menggunakan sort.Ints
		sort.Ints(houses)

		// Menampilkan hasil pengurutan
		fmt.Printf("Nomor rumah terurut untuk daerah ke-%d: ", i+1)
		for _, house := range houses {
			fmt.Printf("%d ", house)
		}
		fmt.Println()
	}

	fmt.Println("Pengurutan selesai.")
}
