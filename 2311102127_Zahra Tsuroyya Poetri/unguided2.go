package main

import (
	"fmt"
)

func selectionSort(arr []int) {
	// Mengurutkan array dengan menggunakan selection sort
	for i := 0; i < len(arr)-1; i++ {
		// Temukan elemen terkecil dalam array yang belum terurut
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j 
			}
		}
		// Tukar elemen terkecil dengan elemen pertama yang belum terurut
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func cariMedian(arr []int) int {
	// Mencari median
	n := len(arr)
	if n%2 == 1 {
		// Jika jumlah elemen ganjil, median adalah elemen tengah
		return arr[n/2]
	} else {
		// Jika jumlah elemen genap, median adalah rata rata dua elemen tengah
		tengah1, tengah2 := arr[n/2-1], arr[n/2]
		return (tengah1 + tengah2) / 2
	}
}

func main() {
	var input int 
	var data []int 

	// Membaca input sampai menemukan angka -5313
	for {
		fmt.Scan(&input)

		// Jika input adalah -5313, maka berhenti 
		if input == -5313 {
			break
		}

		// Jika input adalah 0, cetak median
		if input == 0 {
			// urutkan data dengan selection sort
			selectionSort(data)
			// Cetak median
			median := cariMedian(data)
			fmt.Println(median)
		} else {
			// Jika input bukan 0, simpan data 
			data = append(data, input)
		}
	}
}