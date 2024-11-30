package main

import "fmt"

// Fungsi untuk mengurutkan angka dalam array menggunakan metode selection sort
func urutkanArray(arr []int) {
	panjang := len(arr)
	for i := 0; i < panjang-1; i++ {
		// Menentukan elemen terkecil di subarray yang belum terurut
		indeksTerkecil := i
		for j := i + 1; j < panjang; j++ {
			if arr[j] < arr[indeksTerkecil] {
				indeksTerkecil = j
			}
		}
		// Tukar posisi elemen terkecil dengan elemen pertama yang belum terurut
		arr[i], arr[indeksTerkecil] = arr[indeksTerkecil], arr[i]
	}
}

// Fungsi untuk menghitung dan mengembalikan nilai median dari array yang sudah terurut
func hitungMedian(arr []int) int {
	totalElemen := len(arr)
	if totalElemen%2 == 1 {
		// Jika jumlah elemen ganjil, ambil elemen tengah
		return arr[totalElemen/2]
	}
	// Jika jumlah elemen genap, rata-rata dua elemen tengah
	return (arr[(totalElemen/2)-1] + arr[totalElemen/2]) / 2
}

func main() {
	var angkaList []int
	fmt.Print("Masukkan angka-angka (akhir dengan -5313): ")

	// Loop untuk menerima input angka
	for {
		var angka int
		_, err := fmt.Scan(&angka)

		// Cek jika angka yang dimasukkan adalah -5313 untuk keluar
		if angka == -5313 {
			break
		}

		// Jika terjadi error dalam input, keluar dari loop
		if err != nil {
			fmt.Println("Terjadi kesalahan input.")
			break
		}

		// Jika angka adalah 0, urutkan dan tampilkan median
		if angka == 0 {
			urutkanArray(angkaList)
			fmt.Printf("Median: %d\n", hitungMedian(angkaList))
		} else {
			// Menambahkan angka ke list
			angkaList = append(angkaList, angka)
		}
	}
}
