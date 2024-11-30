// Andreas Besar Wibowo
// 2311102198 / IF-11-05

package main

import "fmt"

// Fungsi untuk mengurutkan array dengan selection sort
func selectionSort(arr []int, asc bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if (asc && arr[j] < arr[idx]) || (!asc && arr[j] > arr[idx]) {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

func main() {
	var n int
	// Input jumlah daerah kerabat
	fmt.Print("Masukkan jumlah daerah kerabat (n) : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		// Input jumlah rumah kerabat untuk setiap daerah
		fmt.Printf("\nMasukkan jumlah rumah kerabat di daerah ke-%d (m): ", i+1)
		fmt.Scan(&m)

		arr := make([]int, m)
		// Input nomor rumah kerabat
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		// untuk memisahkan nomor ganjil dan genap
		var odd, even []int
		for _, num := range arr {
			if num%2 == 0 {
				even = append(even, num)
			} else {
				odd = append(odd, num)
			}
		}

		// untuk mengurutkan ganjil dengan menaik
		selectionSort(odd, true)

		// untuk mengurutkan genap dengan menurun
		selectionSort(even, false)

		// Output nomor rumah yang terurut
		fmt.Printf("\nNomor rumah terurut untuk daerah ke-%d:\n", i+1)

		// Tampilkan nomor ganjil
		for _, num := range odd {
			fmt.Printf("%d ", num)
		}

		// Tampilkan nomor genap
		for _, num := range even {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
