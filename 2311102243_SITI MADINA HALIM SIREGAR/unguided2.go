package main

import (
	"fmt"
)
func insertionSort(arr []int, num int) []int {
	arr = append(arr, num)

	for i := len(arr) - 1; i > 0 && arr[i] < arr[i-1]; i-- {
		arr[i], arr[i-1] = arr[i-1], arr[i]
	}

	return arr
}

func calculateMedian(arr []int) int {
	n := len(arr)
	if n%2 == 1 {

		return arr[n/2]
	}
	return (arr[n/2-1] + arr[n/2]) / 2
}

func main() {
	var input int
	data := []int{}

	fmt.Println("Masukkan rangkaian bilangan bulat (akhiri dengan -5313):")
	for {
		fmt.Scan(&input)
		if input == -5313 {
			break
		}
		if input == 0 {
			if len(data) > 0 {
				median := calculateMedian(data)
				fmt.Println(median)
			}
		} else {

			data = insertionSort(data, input)
		}
	}
}
