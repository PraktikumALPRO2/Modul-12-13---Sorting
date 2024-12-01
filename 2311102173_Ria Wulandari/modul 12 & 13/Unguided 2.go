package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var data []int

	fmt.Println("Masukkan angka satu per satu. Input '0' untuk mencetak median. Input '-5313' untuk keluar.")

	for {
		// Membaca input
		scanner.Scan()
		input := scanner.Text()

		// Konversi input menjadi integer
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid, masukkan angka!")
			continue
		}

		// Periksa apakah input adalah -5313 untuk keluar
		if num == -5313 {
			fmt.Println("Program selesai.")
			break
		}

		// Jika input adalah 0, hitung median
		if num == 0 {
			if len(data) == 0 {
				fmt.Println("Tidak ada data untuk menghitung median.")
				continue
			}

			// Salin data untuk pengurutan
			tempData := make([]int, len(data))
			copy(tempData, data)

			// Urutkan data
			sort.Ints(tempData)

			// Hitung median
			median := calculateMedian(tempData)

			// Tampilkan hasil median
			fmt.Printf("Data terurut: %v\n", tempData)
			fmt.Printf("Median: %d\n\n", median)
			continue
		}

		// Tambahkan input ke data
		data = append(data, num)
	}
}

// Fungsi untuk menghitung median dari array integer terurut
func calculateMedian(sortedData []int) int {
	n := len(sortedData)
	if n%2 == 1 {
		// Jika jumlah elemen ganjil, ambil elemen tengah
		return sortedData[n/2]
	} else {
		// Jika jumlah elemen genap, ambil rata-rata dari dua elemen tengah
		return (sortedData[(n/2)-1] + sortedData[n/2]) / 2
	}
}
