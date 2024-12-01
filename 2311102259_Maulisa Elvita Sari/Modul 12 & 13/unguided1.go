package main

import (
	"fmt"
)

func selectionSort(arr []int, descending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if descending {
				if arr[j] > arr[idx] {
					idx = j
				}
			} else {
				if arr[j] < arr[idx] {
					idx = j
				}
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

func separateOddEven(arr []int) ([]int, []int) {
	var odd, even []int
	for _, num := range arr {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	return odd, even
}

func main() {
	var t int
	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		var n int
		fmt.Printf("\nMasukkan jumlah nomor rumah untuk daerah %d: ", i)
		fmt.Scan(&n)

		arr := make([]int, n)
		fmt.Printf("Masukkan %d nomor rumah: ", n)
		for j := 0; j < n; j++ {
			fmt.Scan(&arr[j])
		}

		odd, even := separateOddEven(arr)

		selectionSort(odd, true)

		selectionSort(even, false)

		fmt.Printf("Hasil untuk daerah %d:\n", i)
		for _, num := range odd {
			fmt.Printf("%d ", num)
		}
		for _, num := range even {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
