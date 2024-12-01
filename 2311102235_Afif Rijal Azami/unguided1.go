package main

import "fmt"

// Fungsi untuk mengurutkan array dengan selection sort
func selectionSort235(arr235 []int, asc235 bool) {
	n235 := len(arr235)
	for i235 := 0; i235 < n235-1; i235++ {
		idx235 := i235
		for j235 := i235 + 1; j235 < n235; j235++ {
			if (asc235 && arr235[j235] < arr235[idx235]) || (!asc235 && arr235[j235] > arr235[idx235]) {
				idx235 = j235
			}
		}
		arr235[i235], arr235[idx235] = arr235[idx235], arr235[i235]
	}
}

func main() {
	var n235 int
	// Input jumlah daerah kerabat
	fmt.Print("Masukkan jumlah daerah kerabat (n) : ")
	fmt.Scan(&n235)

	for i235 := 0; i235 < n235; i235++ {
		var m235 int
		// Input jumlah rumah kerabat untuk setiap daerah
		fmt.Printf("\nMasukkan jumlah rumah kerabat di daerah ke-%d (m): ", i235+1)
		fmt.Scan(&m235)

		arr235 := make([]int, m235)
		// Input nomor rumah kerabat
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m235)
		for j235 := 0; j235 < m235; j235++ {
			fmt.Scan(&arr235[j235])
		}

		// untuk memisahkan nomor ganjil dan genap
		var odd235, even235 []int
		for _, num235 := range arr235 {
			if num235%2 == 0 {
				even235 = append(even235, num235)
			} else {
				odd235 = append(odd235, num235)
			}
		}

		// untuk mengurutkan ganjil dengan menaik
		selectionSort235(odd235, true)

		// untuk mengurutkan genap dengan menurun
		selectionSort235(even235, false)

		// Output nomor rumah yang terurut
		fmt.Printf("\nNomor rumah terurut untuk daerah ke-%d:\n", i235+1)

		// Tampilkan nomor ganjil
		for _, num235 := range odd235 {
			fmt.Printf("%d ", num235)
		}

		// Tampilkan nomor genap
		for _, num235 := range even235 {
			fmt.Printf("%d ", num235)
		}
		fmt.Println()
	}
}
