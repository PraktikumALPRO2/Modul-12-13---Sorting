/* Liya Khoirunnisa - 2311102124 */
package main

import (
	"fmt"
)

// Fungsi untuk Selection Sort
func selectionSort_124(arr []int, ascending bool) {
	// Mendapatkan panjang array
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Temukan indeks elemen terkecil atau terbesar
		index := i
		for j := i + 1; j < n; j++ {
			// Jika ascending, cri elemen terkecil namun jika descending, cari elemen terbesar
			if (ascending && arr[j] < arr[index]) || (!ascending && arr[j] > arr[index]) {
				index = j
			}
		}
		// Tukar elemen
		arr[i], arr[index] = arr[index], arr[i]
	}
}

func main() {
	// Deklarasi variabel
	var jumlahDaerah, jumlahNomor int

	// Input jumlah daerah
	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scan(&jumlahDaerah)

	// Perulangan untuk setiap daerah
	for i := 0; i < jumlahDaerah; i++ {
		// Input jumlah nomor rumah di daerah
		fmt.Print("\nMasukkan jumlah nomor rumah di daerah ", i+1, ": ")
		fmt.Scan(&jumlahNomor)

		// Slice untuk menyimpan nomor rumah
		nomorRumah := make([]int, jumlahNomor)

		// Input nomor rumah
		fmt.Print("Masukkan nomor rumah                   : ")
		for j := 0; j < jumlahNomor; j++ {
			fmt.Scan(&nomorRumah[j])
		}

		// Pisahkan nomor ganjil dan genap
		var nomorGanjil []int
		var nomorGenap []int

		// Perulangan untuk memisahkan nomor ganjil dan genap
		for _, nomor := range nomorRumah {
			if nomor%2 == 0 {
				nomorGenap = append(nomorGenap, nomor)
			} else {
				nomorGanjil = append(nomorGanjil, nomor)
			}
		}

		// Urutkan nomor ganjil dan genap menggunakan Selection Sort
		selectionSort_124(nomorGanjil, true) // Urutkan ganjil secara ascending
		selectionSort_124(nomorGenap, false) // Urutkan genap secara descending

		// Tampilkan hasil
		fmt.Print("Nomor rumah di daerah ", i+1, "                : ")
		for _, nomor := range nomorGanjil {
			fmt.Print(nomor, " ")
		}
		for _, nomor := range nomorGenap {
			fmt.Print(nomor, " ")
		}
		fmt.Println() // Untuk pindah ke baris baru setelah semua nomor ditampilkan
	}
}
