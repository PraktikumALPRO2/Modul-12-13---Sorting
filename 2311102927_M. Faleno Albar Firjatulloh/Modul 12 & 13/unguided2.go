package main

import (
	"fmt"
)

// Fungsi pengurutan menggunakan metode insertion
func urutkanData(daftarBilangan []int) {
	for i := 1; i < len(daftarBilangan); i++ {
		kunci := daftarBilangan[i]
		j := i - 1
		for j >= 0 && daftarBilangan[j] > kunci {
			daftarBilangan[j+1] = daftarBilangan[j]
			j--
		}
		daftarBilangan[j+1] = kunci
	}
}

// Fungsi perhitungan median
func tentanganMedian(daftarBilangan []int) float64 {
	panjangData := len(daftarBilangan)
	if panjangData == 0 {
		return 0
	}

	// Kondisi data ganjil
	if panjangData%2 != 0 {
		return float64(daftarBilangan[panjangData/2])
	}

	// Kondisi data genap
	return float64(daftarBilangan[(panjangData-1)/2]+daftarBilangan[panjangData/2]) / 2.0
}

func main() {
	var dataBilangan []int
	var masukan int

	// Instruksi input
	fmt.Println("Masukkan bilangan (0 untuk menghitung median, -5313 untuk mengakhiri):")
	for {
		fmt.Scan(&masukan)

		if masukan == 0 {
			// Proses perhitungan median
			if len(dataBilangan) > 0 {
				dataSementara := make([]int, len(dataBilangan))
				copy(dataSementara, dataBilangan)
				urutkanData(dataSementara)
				median := tentanganMedian(dataSementara)
				fmt.Printf("%.0f\n", median)
			}
		} else if masukan == -5313 {
			break
		} else {
			// Tambahkan input ke slice
			dataBilangan = append(dataBilangan, masukan)
		}
	}
}
