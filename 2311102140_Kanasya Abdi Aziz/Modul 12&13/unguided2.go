package main

import (
	"fmt"
)

// Fungsi untuk melakukan insertion sort
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Fungsi untuk menghitung median
func hitungMedian(arr []int) float64 {
	n := len(arr)
	if n == 0 {
		return 0
	}

	// Jika jumlah data ganjil
	if n%2 != 0 {
		return float64(arr[n/2])
	}

	// Jika jumlah data genap
	return float64(arr[(n-1)/2]+arr[n/2]) / 2.0
}

func main() {
	var data []int
	var input int

	// Membaca input sampai bertemu angka 0 atau -5313
	fmt.Println("Masukkan bilangan (0 untuk menghitung median, -5313 untuk mengakhiri):")
	for {
		fmt.Scan(&input)

		if input == 0 {
			// Urutkan data yang ada dan hitung median
			if len(data) > 0 {
				tempData := make([]int, len(data))
				copy(tempData, data)
				insertionSort(tempData)
				median := hitungMedian(tempData)
				fmt.Printf("%.0f\n", median)
			}
		} else if input == -5313 {
			break
		} else {
			// Tambahkan input ke slice data
			data = append(data, input)
		}
	}
}
