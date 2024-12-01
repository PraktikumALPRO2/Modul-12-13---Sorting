package main

import (
	"fmt"
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

type DaftarBuku struct {
	Pustaka  []Buku
	NPustaka int
}

func TambahBuku(pustaka *DaftarBuku) {
	var jumlah int
	fmt.Print("Masukkan Jumlah Buku : ")
	fmt.Scan(&jumlah)

	pustaka.NPustaka = jumlah
	pustaka.Pustaka = make([]Buku, jumlah)

	for i := 0; i < jumlah; i++ {
		fmt.Printf("\nMasukkan data untuk buku ke-%d\n", i+1)
		fmt.Print("ID : ")
		fmt.Scan(&pustaka.Pustaka[i].ID)
		fmt.Print("Judul : ")
		fmt.Scan(&pustaka.Pustaka[i].Judul)
		fmt.Print("Penulis : ")
		fmt.Scan(&pustaka.Pustaka[i].Penulis)
		fmt.Print("Penerbit : ")
		fmt.Scan(&pustaka.Pustaka[i].Penerbit)
		fmt.Print("Eksemplar : ")
		fmt.Scan(&pustaka.Pustaka[i].Eksemplar)
		fmt.Print("Tahun : ")
		fmt.Scan(&pustaka.Pustaka[i].Tahun)
		fmt.Print("Rating : ")
		fmt.Scan(&pustaka.Pustaka[i].Rating)
	}
}

func CetakBuku(pustaka DaftarBuku) {
	fmt.Println("\nDaftar Buku :")
	fmt.Printf("%-5s %-25s %-20s %-20s %-10s %-10s %-5s\n", "ID", "Judul", "Penulis", "Penerbit", "Eksemplar", "Tahun", "Rating")
	for _, buku := range pustaka.Pustaka {
		fmt.Printf("%-5s %-25s %-20s %-20s %-10d %-10d %-5d\n",
			buku.ID, buku.Judul, buku.Penulis, buku.Penerbit, buku.Eksemplar, buku.Tahun, buku.Rating)
	}
}

func UrutkanBuku(pustaka *DaftarBuku) {
	n := pustaka.NPustaka
	for i := 1; i < n; i++ {
		key := pustaka.Pustaka[i]
		j := i - 1
		for j >= 0 && pustaka.Pustaka[j].Rating < key.Rating {
			pustaka.Pustaka[j+1] = pustaka.Pustaka[j]
			j--
		}
		pustaka.Pustaka[j+1] = key
	}
}

func CetakBukuTeratas(pustaka DaftarBuku) {
	fmt.Println("\n5 Buku dengan Rating Tertinggi :")
	fmt.Printf("%-25s %-5s\n", "Judul", "Rating")
	for i := 0; i < 5 && i < pustaka.NPustaka; i++ {
		fmt.Printf("%-25s %-5d\n", pustaka.Pustaka[i].Judul, pustaka.Pustaka[i].Rating)
	}
}

func CariBuku(pustaka DaftarBuku) {
	var rating int
	fmt.Print("\nMasukkan rating yang ingin dicari : ")
	fmt.Scan(&rating)

	kiri, kanan := 0, pustaka.NPustaka-1
	ditemukan := false
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if pustaka.Pustaka[tengah].Rating == rating {
			ditemukan = true
			buku := pustaka.Pustaka[tengah]
			fmt.Printf("\nBuku dengan Rating %d :\n", rating)
			fmt.Printf("%-5s %-25s %-20s %-20s %-10s %-10s %-5s\n", "ID", "Judul", "Penulis", "Penerbit", "Eksemplar", "Tahun", "Rating")
			fmt.Printf("%-5s %-25s %-20s %-20s %-10d %-10d %-5d\n",
				buku.ID, buku.Judul, buku.Penulis, buku.Penerbit, buku.Eksemplar, buku.Tahun, buku.Rating)
			break
		} else if pustaka.Pustaka[tengah].Rating < rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	if !ditemukan {
		fmt.Println("\nBuku dengan rating tersebut tidak ditemukan.")
	}
}

func main() {
	var pustaka DaftarBuku

	TambahBuku(&pustaka)

	UrutkanBuku(&pustaka)

	CetakBuku(pustaka)

	CetakBukuTeratas(pustaka)

	CariBuku(pustaka)
}