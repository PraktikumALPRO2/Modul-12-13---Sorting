package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Buku struct {
	ID        int
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

type DaftarBuku []Buku

// Fungsi untuk membaca input angka (integer)
func bacaInt(reader *bufio.Reader, prompt string) int {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, _ := strconv.Atoi(input)
	return num
}

// Fungsi untuk membaca input string
func bacaString(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func DaftarkanBuku(pustaka *DaftarBuku, n int, reader *bufio.Reader) {
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("\nMasukkan data buku ke-%d:\n", i+1)

		buku.ID = bacaInt(reader, "ID: ")
		buku.Judul = bacaString(reader, "Judul: ")
		buku.Penulis = bacaString(reader, "Penulis: ")
		buku.Penerbit = bacaString(reader, "Penerbit: ")
		buku.Eksemplar = bacaInt(reader, "Eksemplar: ")
		buku.Tahun = bacaInt(reader, "Tahun: ")

		for {
			buku.Rating = bacaInt(reader, "Rating (0-10): ")
			if buku.Rating >= 0 && buku.Rating <= 10 {
				break
			}
			fmt.Println("Rating harus antara 0 dan 10. Coba lagi.")
		}

		*pustaka = append(*pustaka, buku)
	}
}

func CetakTerfavorit(pustaka DaftarBuku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}
	favorit := pustaka[0]
	for _, buku := range pustaka {
		if buku.Rating > favorit.Rating {
			favorit = buku
		}
	}
	fmt.Printf("\nBuku Terfavorit: %s oleh %s, %s (%d) - Rating: %d\n", favorit.Judul, favorit.Penulis, favorit.Penerbit, favorit.Tahun, favorit.Rating)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var pustaka DaftarBuku

	jumlahBuku := bacaInt(reader, "Masukkan jumlah buku: ")

	if jumlahBuku <= 0 {
		fmt.Println("Jumlah buku harus lebih dari 0.")
		return
	}

	DaftarkanBuku(&pustaka, jumlahBuku, reader)
	CetakTerfavorit(pustaka)
}
