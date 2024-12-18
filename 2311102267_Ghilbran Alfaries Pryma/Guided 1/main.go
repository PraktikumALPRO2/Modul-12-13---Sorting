package main

import "fmt"

func selectionsort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxmin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxmin] {
				idxmin = j
			}
		}
		arr[i], arr[idxmin] = arr[idxmin], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	for daerah := 1; daerah <= n; daerah++ {
		var m int
		fmt.Printf("\nMasukkan jumlah nomor rumah kerabat untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}

		// Sort the house numbers
		selectionsort(arr, m)

		// Print sorted house numbers for the current daerah
		fmt.Printf("Nomor rumah terurut untuk daerah %d: ", daerah)
		for _, num := range arr {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}