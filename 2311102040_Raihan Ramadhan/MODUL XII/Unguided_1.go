package main

import "fmt"

// Fungsi untuk mengurutkan array secara ascending menggunakan Selection Sort
func selectionSortAscending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// Fungsi untuk mengurutkan array secara descending menggunakan Selection Sort
func selectionSortDescending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

// Fungsi untuk memisahkan bilangan ganjil dan genap
func pisahkanGanjilGenap(arr []int) ([]int, []int) {
	var ganjil, genap []int
	for _, num := range arr {
		if num%2 == 0 {
			genap = append(genap, num)
		} else {
			ganjil = append(ganjil, num)
		}
	}
	return ganjil, genap
}

func main() {
	var t_2311102040 int
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&t_2311102040 )

	// Iterasi untuk setiap daerah
	for daerah := 1; daerah <= t_2311102040; daerah++ {
		var m int
		fmt.Printf("\nMasukkan jumlah nomor rumah untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}

		// Pisahkan nomor ganjil dan genap
		ganjil, genap := pisahkanGanjilGenap(arr)

		// Urutkan ganjil secara naik dan genap secara turun
		selectionSortAscending(ganjil)
		selectionSortDescending(genap)

		// Tampilkan hasil
		fmt.Printf("Nomor rumah terurut untuk daerah %d: ", daerah)
		for _, num := range ganjil {
			fmt.Printf("%d ", num)
		}
		for _, num := range genap {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
