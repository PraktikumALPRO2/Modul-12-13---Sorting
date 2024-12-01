package main

import (
	"fmt"
	"math"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func findMedian(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n%2 == 1 {
		return arr[n/2]
	}
	return int(math.Floor(float64(arr[n/2-1]+arr[n/2]) / 2.0))
}

func main() {
	var num int
	var data []int

	fmt.Println("Masukkan data (akhiri dengan 0 atau -531351 untuk tanda akhir):")
	for {
		fmt.Scan(&num)
		if num == 0 || num == -531351 {
			if len(data) > 0 {
				selectionSort(data) 
				median := findMedian(data)
				fmt.Println("Median:", median)
			}
			if num == 0 {
				data = []int{}
			} else {
				break
			}
		} else {
			data = append(data, num)
		}
	}
}
