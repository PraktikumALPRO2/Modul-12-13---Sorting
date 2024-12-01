package main

import "fmt"

// Fungsi untuk mengurutkan array dengan selection sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Cari elemen terkecil di subarray yang belum terurut
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Tukar elemen terkecil dengan elemen pertama subarray
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// Fungsi untuk menghitung nilai median dari array yang sudah diurutkan
func cariMedian(arr []int) int {
	n := len(arr)
	if n%2 == 1 {
		// Jika jumlah elemen bernilai ganjil, maka ambil elemen tengah
		return arr[n/2]
	}
	// Jika jumlah elemen bernilai genap, maka ambil rata-rata dari dua elemen tengah dan dibulatkan
	return (arr[(n/2)-1] + arr[n/2]) / 2
}

func main() {
	var data []int
	fmt.Print("Masukkan angka-angka: ")

	for {
		var num int
		_, err := fmt.Scan(&num)

		if num == -5313 {
			break
		}

		if err != nil {
			break
		}

		if num == 0 {
			// untuk mengurutkan data dan mencetak median
			selectionSort(data)
			fmt.Println(cariMedian(data))
		} else {
			// menambahkan bilangan ke array
			data = append(data, num)
		}
	}
}