package main

import (
	"fmt"
)

func urutkanAngka(daftarAngka []int) {
	for i := 1; i < len(daftarAngka); i++ {
		angkaKunci := daftarAngka[i]
		j := i - 1
		for j >= 0 && daftarAngka[j] > angkaKunci {
			daftarAngka[j+1] = daftarAngka[j]
			j--
		}
		daftarAngka[j+1] = angkaKunci
	}
}

func hitungMedian(daftarAngka []int) float64 {
	jumlahData := len(daftarAngka)
	if jumlahData == 0 {
		return 0
	}

	if jumlahData%2 != 0 {
		return float64(daftarAngka[jumlahData/2])
	}

	return float64(daftarAngka[(jumlahData-1)/2]+daftarAngka[jumlahData/2]) / 2.0
}

func main() {
	var koleksiAngka []int
	var inputAngka int

	fmt.Println("Masukkan angka (0 untuk menghitung median dan -5313 untuk mengakhiri):")
	for {
		fmt.Scan(&inputAngka)

		if inputAngka == 0 {

			if len(koleksiAngka) > 0 {
				koleksiSementara := make([]int, len(koleksiAngka))
				copy(koleksiSementara, koleksiAngka)
				urutkanAngka(koleksiSementara)
				median := hitungMedian(koleksiSementara)
				fmt.Printf("%.0f\n", median)
			}
		} else if inputAngka == -5313 {
			break
		} else {

			koleksiAngka = append(koleksiAngka, inputAngka)
		}
	}
}
