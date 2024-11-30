package main

import (
	"fmt"
	"sort"
)

func main() {
	var totalDaerah int
	fmt.Print("Masukkan jumlah daerah kerabat: ")
	fmt.Scan(&totalDaerah)

	for i := 1; i <= totalDaerah; i++ {
		var totalRumah int

		fmt.Printf("\nMasukkan jumlah rumah di daerah %d: ", i)
		fmt.Scan(&totalRumah)

		nomorRumah := make([]int, totalRumah)

		fmt.Printf("Masukkan %d nomor rumah: ", totalRumah)
		for j := 0; j < totalRumah; j++ {
			fmt.Scan(&nomorRumah[j])
		}

		var rumahGanjil, rumahGenap []int
		for _, nomor := range nomorRumah {
			if nomor%2 == 0 {
				rumahGenap = append(rumahGenap, nomor)
			} else {
				rumahGanjil = append(rumahGanjil, nomor)
			}
		}

		sort.Ints(rumahGanjil) 
		sort.Ints(rumahGenap) 
		fmt.Printf("\nNomor rumah terurut untuk daerah %d:\n", i)
		if len(rumahGanjil) > 0 { 
			for _, nomor := range rumahGanjil {
				fmt.Printf("%d ", nomor)
			}
		}
		if len(rumahGenap) > 0 { 
			for _, nomor := range rumahGenap {
				fmt.Printf("%d ", nomor)
			}
		}
		fmt.Println()
	}
}
