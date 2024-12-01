package main

import (
	"fmt"
)

func selectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := 0; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}

		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	for daerah := 1; daerah <= n; daerah++ {
		var m int
		fmt.Printf("\n Masukkan jumlah nomor kerabat untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}
		selectionSort(arr, m)

		fmt.Printf("Nomor rumah terurut untuk daerah %d: ", daerah)
		for _, num := range arr {
			fmt.Printf("%d", num)
		}
		fmt.Println()
	}

}