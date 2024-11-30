// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Fungsi Selection Sort untuk mengurutkan angka dalam urutan tertentu
func selectionSort(arr []int, ascending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if (ascending && arr[j] < arr[idx]) || (!ascending && arr[j] > arr[idx]) {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

// Fungsi untuk memisahkan dan mengurutkan nomor ganjil dan genap
func pisahDanUrutkan(arr []int) ([]int, []int) {
	var ganjil, genap []int
	for _, num := range arr {
		if num%2 == 0 {
			genap = append(genap, num)
		} else {
			ganjil = append(ganjil, num)
		}
	}

	// Urutkan ganjil membesar dan genap mengecil
	selectionSort(ganjil, true) // Ganjil membesar
	selectionSort(genap, false) // Genap mengecil

	return ganjil, genap
}

func main() {
	fmt.Println("===== Program Pengurutan Nomor Rumah =====")

	var n int
	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Jumlah daerah harus lebih besar dari 0.")
		return
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("\n-- Daerah ke-%d --\n", i)

		var m int
		fmt.Print("Masukkan jumlah rumah: ")
		fmt.Scan(&m)

		if m <= 0 {
			fmt.Println("Jumlah rumah harus lebih besar dari 0.")
			continue
		}

		// Input nomor rumah
		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah (pisahkan dengan spasi): ", m)
		for j := range arr {
			fmt.Scan(&arr[j])
		}

		// Pisah dan urutkan nomor rumah
		ganjil, genap := pisahDanUrutkan(arr)

		// Tampilkan hasil
		fmt.Printf("Nomor rumah terurut: Ganjil (%v) + Genap (%v)\n", ganjil, genap)
	}

	fmt.Println("\n===== Program Selesai =====")
}
