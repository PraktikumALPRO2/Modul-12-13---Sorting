package main

import (
	"fmt"
	"sort"
)

func main() {
	var data []int
	var number int

	for {
		fmt.Scan(&number) // Membaca input

		if number == -5313 {
			break // Menghentikan program jika input -5313
		}

		if number == 0 {
			// Jika menemukan 0, hitung median
			if len(data) == 0 {
				fmt.Println("No data to calculate median")
				continue
			}

			// Mengurutkan data
			sort.Ints(data)

			// Menghitung median
			n := len(data)
			var median int
			if n%2 == 0 {
				median = (data[n/2-1] + data[n/2]) / 2 // Rata-rata dua nilai tengah
			} else {
				median = data[n/2] // Nilai tengah
			}

			// Mencetak median
			fmt.Println(median)
		} else {
			// Menyimpan data yang bukan 0
			data = append(data, number)
		}
	}
}
