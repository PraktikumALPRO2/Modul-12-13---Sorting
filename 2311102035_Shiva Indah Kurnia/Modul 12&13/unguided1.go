package main

import (
	"fmt"
	"sort"
)

func main() {
	var jumlahDaerah int

	// Input jumlah daerah
	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scan(&jumlahDaerah)

	// Proses untuk setiap daerah
	for i := 1; i <= jumlahDaerah; i++ {
		var jumlahRumah int
		fmt.Printf("\nJumlah rumah di daerah %d: ", i)
		fmt.Scan(&jumlahRumah)

		// Input nomor rumah
		rumah := make([]int, jumlahRumah)
		fmt.Printf("Masukkan %d nomor rumah: ", jumlahRumah)
		for j := 0; j < jumlahRumah; j++ {
			fmt.Scan(&rumah[j])
		}

		// Pemisahan nomor ganjil dan genap
		ganjil, genap := []int{}, []int{}
		for _, n := range rumah {
			if n%2 == 0 {
				genap = append(genap, n)
			} else {
				ganjil = append(ganjil, n)
			}
		}

		// Pengurutan
		sort.Ints(ganjil)                             // Urutkan ganjil naik
		sort.Sort(sort.Reverse(sort.IntSlice(genap))) // Urutkan genap turun

		// Output hasil
		fmt.Printf("\nHasil pengurutan daerah %d:\n", i)
		fmt.Println(append(ganjil, genap...)) // Gabung dan cetak ganjil diikuti genap
	}
}
