package main

import (
	"fmt"
)

func selectionSort(arr []int, ascending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if ascending {
				if arr[j] < arr[idx] {
					idx = j
				}
			} else {
				if arr[j] > arr[idx] {
					idx = j
				}
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	for daerah := 1; daerah <= n; daerah++ {
		var m int
		fmt.Printf("\nMasukkan jumlah nomor rumah kerabat untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}

		// Pisahkan nomor ganjil dan genap
		var ganjil, genap []int
		for _, num := range arr {
			if num%2 == 0 {
				genap = append(genap, num)
			} else {
				ganjil = append(ganjil, num)
			}
		}

		// Urutkan nomor ganjil secara membesar
		selectionSort(ganjil, true)
		// Urutkan nomor genap secara mengecil
		selectionSort(genap, false)

		// Tampilkan hasil
		fmt.Printf("Nomor rumah untuk daerah %d:\n", daerah)
		fmt.Print("Ganjil terurut membesar: ")
		for _, num := range ganjil {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
		fmt.Print("Genap terurut mengecil: ")
		for _, num := range genap {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
