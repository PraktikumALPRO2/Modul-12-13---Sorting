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

	// Meminta jumlah daerah
	fmt.Print("Masukkan jumlah daerah: ")
	scanner.Scan()
	jumlahDaerah, _ := strconv.Atoi(scanner.Text())

	for i := 1; i <= jumlahDaerah; i++ {
		// Meminta jumlah rumah kerabat
		fmt.Printf("Masukkan jumlah rumah kerabat daerah %d: ", i)
		scanner.Scan()
		jumlahRumah, _ := strconv.Atoi(scanner.Text())

		// Meminta nomor rumah
		fmt.Printf("Masukkan nomor rumah daerah %d: ", i)
		scanner.Scan()
		baris := scanner.Text()

		// Memisahkan angka-angka dari input
		angkaStr := strings.Fields(baris)
		if len(angkaStr) != jumlahRumah {
			fmt.Println("Jumlah rumah harus sesuai dengan input!")
			continue
		}

		// Memisahkan angka menjadi ganjil dan genap
		var rumahGanjil, rumahGenap []int
		for _, angka := range angkaStr {
			nomor, _ := strconv.Atoi(angka)
			if nomor%2 == 0 {
				rumahGenap = append(rumahGenap, nomor)
			} else {
				rumahGanjil = append(rumahGanjil, nomor)
			}
		}

		// Mengurutkan ganjil secara menaik
		sort.Ints(rumahGanjil)

		// Mengurutkan genap secara menurun
		sort.Sort(sort.Reverse(sort.IntSlice(rumahGenap)))

		// Mencetak output sesuai format
		fmt.Print("Urutan rumah: ")
		for _, nomor := range rumahGanjil {
			fmt.Printf("%d ", nomor)
		}
		for _, nomor := range rumahGenap {
			fmt.Printf("%d ", nomor)
		}
		fmt.Println()
	}
}
