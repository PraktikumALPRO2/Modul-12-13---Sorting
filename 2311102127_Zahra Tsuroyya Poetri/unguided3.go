package main

import (
	"fmt"
)

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

var Pustaka DaftarBuku
var nPustaka int

// DaftarkanBuku untuk memasukkan data buku
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	fmt.Println("Masukkan data buku (id, judul, penulis, penerbit, tahun, rating, eksemplar):")
	for i := 0; i < n; i++ {
		fmt.Printf("Buku ke-%d:\n", i+1)
		// Pastikan memasukkan rating yang valid
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit,
			&pustaka[i].tahun, &pustaka[i].rating, &pustaka[i].eksemplar)
	}
}

// CetakTerfavorit untuk menampilkan buku dengan rating tertinggi
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	maxRating := -1
	var terfavorit Buku
	for i := 0; i < n; i++ {
		if pustaka[i].rating > maxRating {
			maxRating = pustaka[i].rating
			terfavorit = pustaka[i]
		}
	}
	fmt.Println("\nBuku Terfavorit:")
	fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\n", terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun)
}

// UrutBuku untuk mengurutkan buku berdasarkan rating menggunakan Insertion Sort
func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		// Urutkan buku berdasarkan rating dengan metode Insertion Sort
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

// Cetak5Terbaru untuk menampilkan 5 buku dengan rating tertinggi
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("\n5 Buku dengan Rating Tertinggi:")
	for i := 0; i < 5 && i < n; i++ {
		fmt.Printf("%d. %s\n", i+1, pustaka[i].judul)
	}
}

// CariBuku untuk mencari buku berdasarkan rating menggunakan pencarian biner
func CariBuku(pustaka DaftarBuku, n int, r int) {
	// Pencarian biner
	idxAwal, idxAkhir := 0, n-1
	var hasil []Buku
	for idxAwal <= idxAkhir {
		idxTengah := (idxAwal + idxAkhir) / 2
		if pustaka[idxTengah].rating == r {
			// Cari ke kiri
			i := idxTengah
			for i >= 0 && pustaka[i].rating == r {
				hasil = append(hasil, pustaka[i])
				i--
			}
			// Cari ke kanan
			i = idxTengah + 1
			for i < n && pustaka[i].rating == r {
				hasil = append(hasil, pustaka[i])
				i++
			}
			break
		} else if pustaka[idxTengah].rating > r {
			idxAkhir = idxTengah - 1
		} else {
			idxAwal = idxTengah + 1
		}
	}

	// Periksa apakah ada buku dengan rating tersebut
	if len(hasil) > 0 {
		fmt.Printf("\nBuku dengan rating %d ditemukan:\n", r)
		for _, buku := range hasil {
			fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nEksemplar: %d\nRating: %d\n\n",
				buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.eksemplar, buku.rating)
		}
	} else {
		fmt.Println("\nTidak ada buku dengan rating seperti itu.")
	}
}

	
func main() {
	// Input jumlah buku
	fmt.Print("Masukkan jumlah buku di perpustakaan: ")
	fmt.Scan(&nPustaka)

	// Daftarkan buku-buku
	DaftarkanBuku(&Pustaka, nPustaka)

	// Cetak buku terfavorit (rating tertinggi)
	CetakTerfavorit(Pustaka, nPustaka)

	// Urutkan buku berdasarkan rating menggunakan Insertion Sort
	UrutBuku(&Pustaka, nPustaka)

	// Cetak 5 buku dengan rating tertinggi
	Cetak5Terbaru(Pustaka, nPustaka)

	// Input rating yang dicari
	var rating int
	fmt.Print("\nMasukkan rating yang ingin dicari: ")
	fmt.Scan(&rating)

	// Cari buku dengan rating tersebut
	CariBuku(Pustaka, nPustaka, rating)
}
