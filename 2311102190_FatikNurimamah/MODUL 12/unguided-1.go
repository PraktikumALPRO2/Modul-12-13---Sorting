package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi selection sort untuk mengurutkan array
func selectionSort(angka []int, ascending bool) {
	panjang := len(angka)
	for i := 0; i < panjang-1; i++ {
		indeksEkstrem := i
		for j := i + 1; j < panjang; j++ {
			if (ascending && angka[j] < angka[indeksEkstrem]) || (!ascending && angka[j] > angka[indeksEkstrem]) {
				indeksEkstrem = j
			}
		}
		angka[i], angka[indeksEkstrem] = angka[indeksEkstrem], angka[i]
	}
}

func pisahDanUrutkan(angka []int) ([]int, []int) {
	var ganjil, genap []int

	for _, nilai := range angka {
		if nilai%2 == 0 {
			genap = append(genap, nilai)
		} else {
			ganjil = append(ganjil, nilai)
		}
	}

	selectionSort(ganjil, false) // Ganjil descending
	selectionSort(genap, true)  // Genap ascending

	return ganjil, genap
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Memasukkan jumlah daerah
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	jumlahDaerah, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error: Input jumlah daerah tidak valid!")
		return
	}

	for i := 1; i <= jumlahDaerah; i++ {
		// Memasukkan jumlah nomor rumah untuk daerah ini
		fmt.Printf("\nMasukkan jumlah nomor kerabat untuk daerah %d: ", i)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		jumlahNomor, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: Input jumlah nomor tidak valid!")
			return
		}

		// Memasukkan nomor rumah
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", jumlahNomor)
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		stringAngka := strings.Split(input, " ")

		var angkaDaerah []int
		for _, stringNilai := range stringAngka {
			angka, err := strconv.Atoi(stringNilai)
			if err == nil {
				angkaDaerah = append(angkaDaerah, angka)
			} else {
				fmt.Printf("Error: '%s' bukan angka yang valid\n", stringNilai)
			}
		}

		// Memastikan jumlah input sesuai dengan jumlah nomor yang dimasukkan
		if len(angkaDaerah) != jumlahNomor {
			fmt.Printf("Error: jumlah nomor yang dimasukkan (%d) tidak sesuai dengan jumlah yang diharapkan (%d)\n", len(angkaDaerah), jumlahNomor)
			continue
		}

	
		// Mengurutkan dan memisahkan angka ganjil dan genap
		ganjil, genap := pisahDanUrutkan(angkaDaerah)

		// Menampilkan hasil urutan langsung di bawah input
		fmt.Printf("Ganjil (descending): %v\n", ganjil)
		fmt.Printf("Genap (ascending): %v\n", genap)
	}
}
