// Haifa Zahra Azzimmi
// 2311102163
// S1 IF 11 5

package main

import (
	"fmt"
	"sort"
)

func main() {
	var jumlahDaerah int
	fmt.Print("Masukkan jumlah daerah kerabat: ")
	fmt.Scan(&jumlahDaerah)

	for i := 1; i <= jumlahDaerah; i++ {
		var jumlahRumah int
		fmt.Printf("\nMasukkan jumlah rumah di daerah %d: ", i)
		fmt.Scan(&jumlahRumah)

		nomorRumah := make([]int, jumlahRumah)
		fmt.Printf("Masukkan %d nomor rumah: ", jumlahRumah)
		for j := 0; j < jumlahRumah; j++ {
			fmt.Scan(&nomorRumah[j])
		}

		// Pisahkan nomor rumah ganjil dan genap
		var ganjil []int
		var genap []int
		for _, nomor := range nomorRumah {
			if nomor%2 == 0 {
				genap = append(genap, nomor)
			} else {
				ganjil = append(ganjil, nomor)
			}
		}

		// Urutkan ganjil secara membesar dan genap secara mengecil
		sort.Ints(ganjil)
		sort.Sort(sort.Reverse(sort.IntSlice(genap)))

		// Tampilkan hasil
		fmt.Printf("\nNomor rumah terurut untuk daerah %d:\n", i)
		for _, nomor := range ganjil {
			fmt.Printf("%d ", nomor)
		}
		for _, nomor := range genap {
			fmt.Printf("%d ", nomor)
		}
		fmt.Println()
	}
}
