package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan angka-angka (akhiri dengan -5313):")

	var data_2311102040 []int

	// Membaca input dari pengguna
	for scanner.Scan() {
		line := scanner.Text()
		numStrings := strings.Fields(line)
		for _, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Input tidak valid!")
				return
			}
			if num == -5313 {
				// Mengakhiri program jika ditemukan -5313
				return
			}
			if num == 0 {
				// Jika input adalah 0, hitung median dari data yang sudah terkumpul
				if len(data_2311102040) == 0 {
					fmt.Println("Data kosong!")
				} else {
					selectionSort(data_2311102040)
					median := calculateMedian(data_2311102040)
					fmt.Println(median)
				}
			} else {
				// Tambahkan bilangan ke dalam data
				data_2311102040 = append(data_2311102040, num)
			}
		}
	}
}

// Fungsi untuk menghitung median
func calculateMedian(data_2311102040 []int) int {
	n := len(data_2311102040)
	if n%2 == 1 {
		// Jika jumlah data ganjil, median adalah nilai tengah
		return data_2311102040[n/2]
	}
	// Jika jumlah data genap, median adalah rata-rata dari dua nilai tengah
	return (data_2311102040[n/2-1] + data_2311102040[n/2]) / 2
}

// Fungsi untuk mengurutkan data menggunakan selection sort
func selectionSort(data_2311102040 []int) {
	n := len(data_2311102040)
	for i := 0; i < n-1; i++ {
		// Temukan indeks elemen terkecil di bagian yang belum diurutkan
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data_2311102040[j] < data_2311102040[minIdx] {
				minIdx = j
			}
		}
		// Tukar elemen terkecil dengan elemen di posisi i
		data_2311102040[i], data_2311102040[minIdx] = data_2311102040[minIdx], data_2311102040[i]
	}
}
