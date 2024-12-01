package main

import (
	"fmt"
	"sort"
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

func UrutBuku(buku []Buku) {
	sort.Slice(buku, func(i, j int) bool {
		return buku[i].Rating > buku[j].Rating
	})
}

func CetakTerfavorit(buku Buku) {
	fmt.Println("Buku Terfavorit:")
	fmt.Printf("%+v\n", buku)
}

func Cetak5Terbaru(buku []Buku) {
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 0; i < 5 && i < len(buku); i++ {
		fmt.Println(buku[i].Judul)
	}
}

func CariBuku(buku []Buku, ratingDicari int) int {
	left, right := 0, len(buku)-1
	for left <= right {
		mid := (left + right) / 2
		if buku[mid].Rating == ratingDicari {
			return mid
		} else if buku[mid].Rating < ratingDicari {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	var n int
	var pustaka []Buku

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	pustaka = make([]Buku, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (ID Judul Penulis Penerbit Eksemplar Tahun Rating): ", i+1)
		fmt.Scan(&pustaka[i].ID, &pustaka[i].Judul, &pustaka[i].Penulis, &pustaka[i].Penerbit, &pustaka[i].Eksemplar, &pustaka[i].Tahun, &pustaka[i].Rating)
	}

	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&n)

	UrutBuku(pustaka)

	CetakTerfavorit(pustaka[0])

	Cetak5Terbaru(pustaka)

	index := CariBuku(pustaka, n)
	if index != -1 {
		fmt.Printf("Buku dengan rating %d ditemukan: %+v\n", n, pustaka[index])
	} else {
		fmt.Printf("Buku dengan rating %d tidak ditemukan\n", n)
	}
}
