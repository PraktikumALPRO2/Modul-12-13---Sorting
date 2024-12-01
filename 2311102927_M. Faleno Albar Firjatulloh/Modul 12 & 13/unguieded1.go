package main

import (
	"fmt"
	"sort"
)

type Daerah struct {
	ganjil []int
	genap  []int
}

func prosesNomorRumah(nomor []int) Daerah {
	var hasil Daerah

	for _, n := range nomor {
		if n%2 == 0 {
			hasil.genap = append(hasil.genap, n)
		} else {
			hasil.ganjil = append(hasil.ganjil, n)
		}
	}

	sort.Ints(hasil.ganjil)
	sort.Sort(sort.Reverse(sort.IntSlice(hasil.genap)))

	return hasil
}

func main() {
	var jumlahDaerah int

	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scan(&jumlahDaerah)

	for i := 1; i <= jumlahDaerah; i++ {
		var jumlahRumah int

		fmt.Printf("\nJumlah rumah di daerah %d: ", i)
		fmt.Scan(&jumlahRumah)

		rumah := make([]int, 0, jumlahRumah)
		fmt.Printf("Masukkan %d nomor rumah: ", jumlahRumah)

		for j := 0; j < jumlahRumah; j++ {
			var nomor int
			fmt.Scan(&nomor)
			rumah = append(rumah, nomor)
		}

		hasil := prosesNomorRumah(rumah)

		fmt.Printf("\nHasil pengurutan daerah %d:\n", i)

		for _, n := range hasil.ganjil {
			fmt.Printf("%d ", n)
		}

		for _, n := range hasil.genap {
			fmt.Printf("%d ", n)
		}
		fmt.Println()
	}
}
