package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan insertion sort
func urutkanDenganInsertion(data []int) {
	for i := 1; i < len(data); i++ {
		elemen := data[i]
		j := i - 1
		for j >= 0 && data[j] > elemen {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = elemen
	}
}

// Fungsi untuk menghitung median dari array yang sudah terurut
func cariMedian(data []int) float64 {
	ukuran := len(data)
	if ukuran == 0 {
		return 0
	}

	// Jika jumlah elemen ganjil
	if ukuran%2 != 0 {
		return float64(data[ukuran/2])
	}

	// Jika jumlah elemen genap
	tengah := ukuran / 2
	return float64(data[tengah-1]+data[tengah]) / 2.0
}

func main() {
	var angka int
	nilai := []int{}

	fmt.Println("Masukkan angka (0 untuk menampilkan median, -5313 untuk keluar):")
	for {
		fmt.Scan(&angka)

		if angka == 0 {
			if len(nilai) > 0 {
				// Salin dan urutkan data sebelum menghitung median
				salinan := append([]int{}, nilai...)
				urutkanDenganInsertion(salinan)
				median := cariMedian(salinan)
				fmt.Printf("%.0f\n", median)
			}
		} else if angka == -5313 {
			// Keluar dari program
			break
		} else {
			// Tambahkan angka ke daftar nilai
			nilai = append(nilai, angka)
		}
	}
}
