package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Print("Masukkan Jumlah Daerah (n) : ")
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		var n int
		fmt.Printf("\nMasukkan Jumlah Rumah Daerah %d : ", i)
		fmt.Scan(&n)

		rumah := make([]int, n)
		fmt.Printf("Masukkan %d Nomor Rumah Daerah %d : ", n, i)
		for j := 0; j < n; j++ {
			fmt.Scan(&rumah[j])
		}

		var ganjil, genap []int
		for _, nomor := range rumah {
			if nomor%2 == 0 {
				genap = append(genap, nomor)
			} else {
				ganjil = append(ganjil, nomor)
			}
		}

		sort.Ints(ganjil)
		sort.Sort(sort.Reverse(sort.IntSlice(genap)))

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