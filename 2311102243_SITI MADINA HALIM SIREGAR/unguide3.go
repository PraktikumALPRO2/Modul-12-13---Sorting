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
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku %d (format: ID Judul Penulis Penerbit Eksemplar Tahun Rating):\n", i+1)
		data, _ := reader.ReadString('\n')
		data = strings.TrimSpace(data)
		fields := strings.Split(data, " ")

		id, _ := strconv.Atoi(fields[0])
		eksemplar, _ := strconv.Atoi(fields[4])
		tahun, _ := strconv.Atoi(fields[5])
		rating, _ := strconv.Atoi(fields[6])

		buku := Buku{
			ID:        id,
			Judul:     fields[1],
			Penulis:   fields[2],
			Penerbit:  fields[3],
			Eksemplar: eksemplar,
			Tahun:     tahun,
			Rating:    rating,
		}

		*pustaka = append(*pustaka, buku)
	}
}
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}

	terfavorit := pustaka[0]
	for _, buku := range pustaka {
		if buku.Rating > terfavorit.Rating {
			terfavorit = buku
		}
	}

	fmt.Println("Buku terfavorit:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n",
		terfavorit.Judul, terfavorit.Penulis, terfavorit.Penerbit, terfavorit.Tahun)
}
func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := (*pustaka)[i]
		j := i - 1
		for j >= 0 && (*pustaka)[j].Rating < key.Rating {
			(*pustaka)[j+1] = (*pustaka)[j]
			j--
		}
		(*pustaka)[j+1] = key
	}
}
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("Lima buku dengan rating tertinggi:")
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("%s\n", pustaka[i].Judul)
	}
}
func CariBuku(pustaka DaftarBuku, n int, r int) {
	low, high := 0, n-1

	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].Rating == r {
			buku := pustaka[mid]
			fmt.Println("Buku ditemukan:")
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
				buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Eksemplar, buku.Rating)
			return
		} else if pustaka[mid].Rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	var n, rating int

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan jumlah buku: ")
	input, _ := reader.ReadString('\n')
	n, _ = strconv.Atoi(strings.TrimSpace(input))
	DaftarkanBuku(&pustaka, n)
	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	fmt.Print("Masukkan rating yang ingin dicari: ")
	input, _ = reader.ReadString('\n')
	rating, _ = strconv.Atoi(strings.TrimSpace(input))
	CariBuku(pustaka, n, rating)
}
