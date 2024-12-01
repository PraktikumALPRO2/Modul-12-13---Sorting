package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Buku struct {
	ID        string
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

type DaftarBuku []Buku

// Prosedur untuk mendaftarkan buku
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating):\n", i+1)
		scanner.Scan()
		input := strings.Split(scanner.Text(), ",")
		if len(input) != 7 {
			fmt.Println("Data buku tidak lengkap, masukkan ulang!")
			i--
			continue
		}
		eksemplar, _ := strconv.Atoi(strings.TrimSpace(input[4]))
		tahun, _ := strconv.Atoi(strings.TrimSpace(input[5]))
		rating, _ := strconv.Atoi(strings.TrimSpace(input[6]))
		*pustaka = append(*pustaka, Buku{
			ID:        strings.TrimSpace(input[0]),
			Judul:     strings.TrimSpace(input[1]),
			Penulis:   strings.TrimSpace(input[2]),
			Penerbit:  strings.TrimSpace(input[3]),
			Eksemplar: eksemplar,
			Tahun:     tahun,
			Rating:    rating,
		})
	}
}

// Prosedur untuk mencetak buku terfavorit
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}
	maxRating := pustaka[0].Rating
	terfavorit := pustaka[0]
	for _, buku := range pustaka {
		if buku.Rating > maxRating {
			maxRating = buku.Rating
			terfavorit = buku
		}
	}
	fmt.Printf("Buku Terfavorit: ID: %s, %s oleh %s (Penerbit: %s, Tahun: %d)\n", terfavorit.ID, terfavorit.Judul, terfavorit.Penulis, terfavorit.Penerbit, terfavorit.Tahun)
}

// Prosedur untuk mengurutkan buku berdasarkan rating (descending)
func UrutBuku(pustaka *DaftarBuku, n int) {
	if n < 2 {
		return
	}
	sort.Slice(*pustaka, func(i, j int) bool {
		return (*pustaka)[i].Rating > (*pustaka)[j].Rating
	})
}

// Prosedur untuk mencetak 5 buku terbaru
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 0; i < n && i < 5; i++ {
		fmt.Printf("%d. ID: %s, %s (Rating: %d)\n", i+1, pustaka[i].ID, pustaka[i].Judul, pustaka[i].Rating)
	}
}

// Prosedur untuk mencari buku berdasarkan rating menggunakan pencarian biner
func CariBuku(pustaka DaftarBuku, n, targetRating int) {
	low, high := 0, n-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].Rating == targetRating {
			fmt.Printf("Ditemukan: ID: %s, %s oleh %s (Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d)\n",
				pustaka[mid].ID, pustaka[mid].Judul, pustaka[mid].Penulis, pustaka[mid].Penerbit, pustaka[mid].Tahun, pustaka[mid].Eksemplar, pustaka[mid].Rating)
			return
		} else if pustaka[mid].Rating > targetRating {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	var n, targetRating int

	fmt.Println("Masukkan jumlah buku:")
	fmt.Scanln(&n)
	DaftarkanBuku(&pustaka, n)

	// Mengurutkan buku berdasarkan rating
	fmt.Println("\nMengurutkan buku berdasarkan rating...")
	UrutBuku(&pustaka, n)

	// Menampilkan buku terfavorit
	fmt.Println("\nBuku Terfavorit:")
	CetakTerfavorit(pustaka, n)

	// Menampilkan 5 buku terbaru berdasarkan rating
	fmt.Println("\n5 Buku Terbaru Berdasarkan Rating:")
	Cetak5Terbaru(pustaka, n)

	// Mencari buku berdasarkan rating
	fmt.Println("\nMasukkan rating buku yang ingin dicari:")
	fmt.Scanln(&targetRating)
	CariBuku(pustaka, n, targetRating)
}
