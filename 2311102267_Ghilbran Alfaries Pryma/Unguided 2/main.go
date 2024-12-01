package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var data []int // Array untuk menyimpan data yang dibaca

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan bilangan bulat (akhiri dengan -5313):")

	for scanner.Scan() {
		input := scanner.Text()
		// Konversi input menjadi bilangan bulat
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid, masukkan bilangan bulat!")
			continue
		}

		if num == -5313 { // Tanda akhir input
			break
		}

		if num == 0 {
			if len(data) == 0 {
				fmt.Println("Tidak ada data untuk dihitung median.")
				continue
			}

			// Salin dan urutkan data menggunakan metode insertion sort
			sortedData := make([]int, len(data))
			copy(sortedData, data)
			sort.Slice(sortedData, func(i, j int) bool { return sortedData[i] < sortedData[j] })

			// Hitung median
			median := 0
			length := len(sortedData)
			if length%2 == 0 {
				median = (sortedData[length/2-1] + sortedData[length/2]) / 2
			} else {
				median = sortedData[length/2]
			}
			fmt.Println("Median saat ini:", median)
		} else {
			// Tambahkan bilangan ke array data
			data = append(data, num)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error saat membaca input:", err)
	}
}
 