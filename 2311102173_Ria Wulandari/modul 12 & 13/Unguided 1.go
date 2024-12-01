package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi untuk insertion sort (ascending)
func insertionSortAsc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

// Fungsi untuk insertion sort (descending)
func insertionSortDesc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] < key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Masukkan jumlah baris input
	fmt.Print("Masukkan jumlah baris masukan: ")
	scanner.Scan()
	t, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Input jumlah baris tidak valid.")
		return
	}

	// Proses setiap baris masukan
	for i := 1; i <= t; i++ {
		fmt.Printf("Masukkan baris ke-%d (pisahkan dengan spasi): ", i)
		scanner.Scan()
		input := scanner.Text()

		// Pisahkan input ke dalam array bilangan
		numbers := strings.Fields(input)
		var ganjil, genap []int
		for _, num := range numbers {
			val, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Bilangan '%s' tidak valid, diabaikan.\n", num)
				continue
			}
			if val%2 == 0 {
				genap = append(genap, val)
			} else {
				ganjil = append(ganjil, val)
			}
		}

		// Urutkan bilangan ganjil (ascending) dan genap (descending)
		ganjil = insertionSortAsc(ganjil)
		genap = insertionSortDesc(genap)

		// Tampilkan hasil
		fmt.Printf("Hasil untuk baris ke-%d:\n", i)
		fmt.Println("Ganjil terurut membesar:", ganjil)
		fmt.Println("Genap terurut mengecil:", genap)
		fmt.Println()
	}
}
