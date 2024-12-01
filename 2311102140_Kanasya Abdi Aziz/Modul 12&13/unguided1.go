package main

import (
	"fmt"
	"sort"
)

// Struct untuk menyimpan data daerah
type Daerah struct {
	ganjil []int
	genap  []int
}

// Fungsi untuk memisahkan dan mengurutkan nomor rumah
func prosesNomorRumah(nomor []int) Daerah {
	var hasil Daerah

	// Pemisahan nomor ganjil dan genap
	for _, n := range nomor {
		if n%2 == 0 {
			hasil.genap = append(hasil.genap, n)
		} else {
			hasil.ganjil = append(hasil.ganjil, n)
		}
	}

	// Pengurutan menggunakan package sort
	sort.Ints(hasil.ganjil)                             // Urutkan ganjil naik
	sort.Sort(sort.Reverse(sort.IntSlice(hasil.genap))) // Urutkan genap turun

	return hasil
}

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

		// Membuat slice untuk nomor rumah
		rumah := make([]int, 0, jumlahRumah)
		fmt.Printf("Masukkan %d nomor rumah: ", jumlahRumah)

		// Input nomor rumah
		for j := 0; j < jumlahRumah; j++ {
			var nomor int
			fmt.Scan(&nomor)
			rumah = append(rumah, nomor)
		}

		// Proses pengurutan
		hasil := prosesNomorRumah(rumah)

		// Output hasil
		fmt.Printf("\nHasil pengurutan daerah %d:\n", i)
		// Cetak nomor ganjil
		for _, n := range hasil.ganjil {
			fmt.Printf("%d ", n)
		}
		// Cetak nomor genap
		for _, n := range hasil.genap {
			fmt.Printf("%d ", n)
		}
		fmt.Println()
	}
}
