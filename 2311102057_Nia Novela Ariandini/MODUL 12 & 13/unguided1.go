package main

import (
	"fmt"
	"sort"
)

func main() {
	// Masukkan jumlah daerah
	var t int
	fmt.Print("Masukkan Jumlah Daerah (n) : ")
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		// Input jumlah rumah untuk daerah ke-i
		var n int
		fmt.Printf("\nMasukkan Jumlah Rumah Daerah %d : ", i)
		fmt.Scan(&n)

		// Input nomor rumah untuk daerah ke-i
		rumah := make([]int, n)
		fmt.Printf("Masukkan %d Nomor Rumah Daerah %d : ", n, i)
		for j := 0; j < n; j++ {
			fmt.Scan(&rumah[j])
		}

		// Memisahkan nomor ganjil dan genap
		var ganjil, genap []int
		for _, nomor := range rumah {
			if nomor%2 == 0 {
				genap = append(genap, nomor)
			} else {
				ganjil = append(ganjil, nomor)
			}
		}

		// Mengurutkan ganjil membesar dan genap mengecil
		sort.Ints(ganjil)
		sort.Sort(sort.Reverse(sort.IntSlice(genap)))

		// Menampilkan hasil dalam format seperti gambar
		fmt.Printf("Nomor Rumah Terurut (Ganjil - Genap) Daerah %d : ", i)
		for _, nomor := range ganjil {
			fmt.Print(nomor, " ")
		}
		for _, nomor := range genap {
			fmt.Print(nomor, " ")
		}
		fmt.Println()
	}
}
