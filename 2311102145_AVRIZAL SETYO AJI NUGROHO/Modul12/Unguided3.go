package main

import (
	"fmt"
)

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Println("Masukkan jumlah buku:")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("\nMasukkan data buku ke-%d:\n", i+1)
		fmt.Print("ID: ")
		fmt.Scan(&pustaka[i].id)
		fmt.Print("Judul: ")
		fmt.Scan(&pustaka[i].judul)
		fmt.Print("Penulis: ")
		fmt.Scan(&pustaka[i].penulis)
		fmt.Print("Penerbit: ")
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Print("Eksemplar: ")
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Print("Tahun: ")
		fmt.Scan(&pustaka[i].tahun)
		fmt.Print("Rating: ")
		fmt.Scan(&pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada buku dalam pustaka.")
		return
	}
	maxRating := pustaka[0].rating
	maxIndex := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > maxRating {
			maxRating = pustaka[i].rating
			maxIndex = i
		}
	}
	buku := pustaka[maxIndex]
	fmt.Printf("Buku terfavorit:\nJudul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n", buku.judul, buku.penulis, buku.penerbit, buku.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("5 Buku dengan rating tertinggi:")
	count := 0
	for i := 0; i < n && count < 5; i++ {
		fmt.Printf("%d. %s\n", count+1, pustaka[i].judul)
		count++
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if pustaka[mid].rating == r {
			buku := pustaka[mid]
			fmt.Printf("Buku ditemukan:\nJudul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
				buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.eksemplar, buku.rating)
			return
		} else if pustaka[mid].rating < r {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	var n int

	DaftarkanBuku(&pustaka, &n)

	CetakTerfavorit(pustaka, n)

	UrutBuku(&pustaka, n)

	Cetak5Terbaru(pustaka, n)

	var rating int
	fmt.Println("Masukkan rating buku yang ingin dicari:")
	fmt.Scan(&rating)
	CariBuku(pustaka, n, rating)
}
