package main

import (
	"fmt"
	"math"
)

func sort(arr []int) {
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

func median(arr []int) int {
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
	var n int
	var nums []int

	fmt.Println("Masukkan angka (akhiri dengan 0 atau -5313 sebagai penanda akhir):")
	for {
		fmt.Scan(&n)
		if n == 0 || n == -5313 {
			if len(nums) > 0 {
				sort(nums)
				fmt.Println("Median:", median(nums))
			}
			if n == 0 {
				nums = []int{}
			} else {
				break
			}
		} else {
			nums = append(nums, n)
		}
	}
}
