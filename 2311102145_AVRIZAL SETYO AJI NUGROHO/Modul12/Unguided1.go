package main

import "fmt"

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

func main() {
	var n int
	fmt.Print("Masukan jumlah daerah kerabat (n): ")
	fmt.Scanln(&n)

	for daerah := 1; daerah <= n; daerah++ {
		var m int
		fmt.Printf("Masukan jumlah nomor rumah kerabat %d: ", daerah)
		fmt.Scan(&m)

		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}

		var ganjil, genap []int
		for _, num := range arr {
			if num%2 == 0 {
				genap = append(genap, num)
			} else {
				ganjil = append(ganjil, num)
			}
		}

		selectionSortAscending(ganjil)
		selectionSortDescending(genap)

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
