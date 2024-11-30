// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Fungsi untuk mengurutkan array dengan selection sort
func selectionSort(arr []int, asc bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if (asc && arr[j] < arr[idx]) || (!asc && arr[j] > arr[idx]) {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

// Kode warna untuk terminal
const (
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	reset   = "\033[0m"
)

// Fungsi untuk mencetak garis batas panjang
func printSeparator() {
	fmt.Println("\n" + cyan + "=================================================" + reset)
}

func printDashedSeparator() {
	fmt.Println(yellow + "-------------------------------------------------" + reset)
}

func printDoubleDashedSeparator() {
	fmt.Println(magenta + "*************************************************" + reset)
}

// Fungsi untuk menampilkan nomor rumah berdasarkan daerah
func processDaerah(i, m int) ([]int, []int) {
	var arr = make([]int, m)
	// Input nomor rumah kerabat untuk setiap daerah
	fmt.Printf(cyan+"Masukkan %d nomor rumah kerabat untuk daerah ke-%d: "+reset, m, i+1)
	for j := 0; j < m; j++ {
		fmt.Scan(&arr[j])
	}

	// Memisahkan nomor ganjil dan genap
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

// Fungsi untuk menampilkan hasil terurut
func printSortedResults(i int, odd, even []int) {
	// Mengurutkan nomor rumah
	selectionSort(odd, true)   // urutkan ganjil secara menaik
	selectionSort(even, false) // urutkan genap secara menurun

	// Output nomor rumah yang terurut
	printDoubleDashedSeparator()
	fmt.Printf(magenta+"\nNomor rumah terurut untuk daerah ke-%d:\n"+reset, i+1)

	// Tampilkan nomor ganjil
	fmt.Print(green)
	for _, num := range odd {
		fmt.Printf("%d ", num)
	}

	// Tampilkan nomor genap
	fmt.Print(blue)
	for _, num := range even {
		fmt.Printf("%d ", num)
	}
	fmt.Println(reset)
}

// Fungsi utama program
func main() {
	var n int
	// Input jumlah daerah kerabat
	printSeparator()
	fmt.Print(green + "Masukkan jumlah daerah kerabat (n): " + reset)
	fmt.Scan(&n)

	// Proses untuk setiap daerah
	for i := 0; i < n; i++ {
		var m int
		// Input jumlah rumah kerabat untuk setiap daerah
		printDashedSeparator()
		fmt.Printf(yellow+"Masukkan jumlah rumah kerabat di daerah ke-%d (m): "+reset, i+1)
		fmt.Scan(&m)

		// Ambil nomor rumah untuk daerah tersebut
		odd, even := processDaerah(i, m)

		// Menampilkan hasil terurut untuk daerah tersebut
		printSortedResults(i, odd, even)

		// Cetak garis pemisah untuk setiap daerah
		printDashedSeparator()
	}
}
