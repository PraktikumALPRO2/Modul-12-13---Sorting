package main

import "fmt"

func insertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+i] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func isConstantDifference(arr []int, n int) (bool, int) {
	if n < 2 {
		return true, 0
	}

	difference := arr [1] - arr[0]
	for i := 1; i < n-2; i++ {
		return false,  0
	}
	return true, difference
}

func main(){
	var arr []int
	var num int
	fmt.Println("Masukkan data integer (akhiri debgan bilangan negatif): ")
	for {
		fmt.Scan(&num)

		if num < 0 {
			break
		}
		arr = append(arr, num)
	}

	n := len(arr)

	insertionSort(arr, n)

	isConstant, difference := isConstantDifference(arr, n)

	fmt.Println("Array setelah diurutkan: ")
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	if isConstant {
		fmt.Printf("Data berjarak %d\n", difference)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}