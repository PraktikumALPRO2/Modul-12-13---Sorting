/* Liya Khoirunisa - 2311102124 */
package main

import (
	"fmt"
)

// Fungsi untuk menghitung median
func hitungMedian_124(data []int) int {
	n := len(data)
	if n%2 == 1 {
		// Jika jumlah data ganjil, ambil nilai tengah
		return data[n/2]
	} else {
		// Jika jumlah data genap, ambil rata-rata dua nilai tengah dibulatkan ke bawah
		return (data[n/2-1] + data[n/2]) / 2
	}
}

// Fungsi untuk mengurutkan data menggunakan selection sort
func selectionSort_124(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		// Menentukan indeks elemen terkecil di sisa array
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Menukar elemen terkecil dengan elemen di posisi i
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	// Deklarasi variabel
	var data []int
	var num int

	// Meminta input rangkaian bilangan bulat
	fmt.Println("Masukkan rangkaian bilangan bulat (ketik -5313 untuk mengakhiri): ")
	for {
		// Membaca input dari pengguna
		fmt.Scan(&num)

		// Jika menginputkan -5313, maka keluar dari loop
		if num == -5313 {
			break
		}

		// Jika input 0 maka data akan diproses
		if num == 0 {
			if len(data) > 0 {
				// Urutkan data menggunakan selection sort
				selectionSort_124(data)

				// Tampilkan data yang sudah tersusun
				fmt.Println("Data yang sudah tersusun:", data)

				// Hitung median
				median := hitungMedian_124(data)
				fmt.Println("Median:", median)
			}
			// Tidak mengosongkan data untuk pembacaan berikutnya
		} else {
			data = append(data, num)
		}
	}
}
