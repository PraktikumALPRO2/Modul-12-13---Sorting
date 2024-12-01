package main

import (
	"fmt"
	"sort"
)

// Fungsi untuk menghitung median
func findMedian(numbers []int) float64 {
	sort.Ints(numbers)
	n := len(numbers)
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		return float64(numbers[n/2-1]+numbers[n/2]) / 2
	} else {
		return float64(numbers[n/2])
	}
}

func main() {
	var numbers []int
	var input int

	fmt.Println("Enter numbers (0 to calculate median, -5313 to stop):")

	for {
		fmt.Scan(&input)
		if input == 0 {
			if len(numbers) > 0 {
				median := findMedian(numbers)
				fmt.Printf("Median: %.0f\n", median)
			}
		} else if input == -5313 {
			break
		} else {
			numbers = append(numbers, input)
		}
	}
}
