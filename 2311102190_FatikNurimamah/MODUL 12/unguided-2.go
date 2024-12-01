package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan data (akhiri dengan -5313): ")
	scanner.Scan()
	inputData := scanner.Text()

	data := strings.Fields(inputData)
	var processedNumbers []int
	for _, value := range data {
		number, _ := strconv.Atoi(value)
		if number == -5313 {
			break
		}
		if number == 0 {
			// Ketika menemukan angka 0, hitung median dari angka yang telah terkumpul
			if len(processedNumbers) > 0 {
				sort.Ints(processedNumbers)
				medianValue := CariMedian(processedNumbers)
				fmt.Printf("Median dari data saat ini (%v) adalah: %d\n", processedNumbers, medianValue)
			}
		} else {
			processedNumbers = append(processedNumbers, number)
		}
	}
}

func CariMedian(numbers []int) int {
	length := len(numbers)
	if length%2 == 1 {
		return numbers[length/2]
	}
	// Untuk jumlah data genap, ambil rata-rata dua nilai tengah
	return (numbers[length/2-1] + numbers[length/2]) / 2
}
