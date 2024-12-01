package main

import (
	"fmt"
)

const maxBuku = 7919

// Struct untuk data buku
type Buku struct {
	id        string
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

// Struct untuk daftar buku
type KoleksiBuku []Buku

// Fungsi untuk menambahkan buku baru
func TambahkanBuku(koleksi *KoleksiBuku, jumlah *int) {
	var bukuBaru Buku
	fmt.Print("Masukkan ID: ")
	fmt.Scan(&bukuBaru.id)
	fmt.Print("Masukkan Judul: ")
	fmt.Scan(&bukuBaru.judul)
	fmt.Print("Masukkan Penulis: ")
	fmt.Scan(&bukuBaru.penulis)
	fmt.Print("Masukkan Penerbit: ")
	fmt.Scan(&bukuBaru.penerbit)
	fmt.Print("Masukkan Jumlah Eksemplar: ")
	fmt.Scan(&bukuBaru.eksemplar)
	fmt.Print("Masukkan Tahun: ")
	fmt.Scan(&bukuBaru.tahun)
	fmt.Print("Masukkan Rating: ")
	fmt.Scan(&bukuBaru.rating)

	*koleksi = append(*koleksi, bukuBaru)
	*jumlah++
}

// Fungsi untuk mencetak buku dengan rating tertinggi
func TampilkanFavorit(koleksi KoleksiBuku, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Tidak ada buku dalam koleksi")
		return
	}

	ratingTertinggi := koleksi[0].rating
	indeksFavorit := 0
	for i := 1; i < jumlah; i++ {
		if koleksi[i].rating > ratingTertinggi {
			ratingTertinggi = koleksi[i].rating
			indeksFavorit = i
		}
	}

	fmt.Println("\nBuku Terfavorit:")
	fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nRating: %d\n",
		koleksi[indeksFavorit].judul, koleksi[indeksFavorit].penulis,
		koleksi[indeksFavorit].penerbit, koleksi[indeksFavorit].tahun,
		koleksi[indeksFavorit].rating)
}

// Fungsi untuk mengurutkan buku berdasarkan rating menggunakan insertion sort
func SortirBuku(koleksi *KoleksiBuku, jumlah int) {
	for i := 1; i < jumlah; i++ {
		bukuSementara := (*koleksi)[i]
		j := i - 1
		for j >= 0 && (*koleksi)[j].rating < bukuSementara.rating {
			(*koleksi)[j+1] = (*koleksi)[j]
			j--
		}
		(*koleksi)[j+1] = bukuSementara
	}
}

// Fungsi untuk menampilkan 5 buku dengan rating tertinggi
func CetakTop5(koleksi KoleksiBuku, jumlah int) {
	fmt.Println("\n5 Buku dengan Rating Tertinggi:")
	jumlahTampil := 5
	if jumlah < 5 {
		jumlahTampil = jumlah
	}
	for i := 0; i < jumlahTampil; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, koleksi[i].judul, koleksi[i].rating)
	}
}

// Fungsi untuk mencari buku dengan rating tertentu menggunakan binary search
func TemukanBuku(koleksi KoleksiBuku, jumlah int, target int) {
	kiri := 0
	kanan := jumlah - 1
	ditemukan := false

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if koleksi[tengah].rating == target {
			fmt.Printf("\nBuku dengan rating %d ditemukan:\n", target)
			fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nEksemplar: %d\n",
				koleksi[tengah].judul, koleksi[tengah].penulis,
				koleksi[tengah].penerbit, koleksi[tengah].tahun,
				koleksi[tengah].eksemplar)
			ditemukan = true
			break
		}
		if koleksi[tengah].rating < target {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if !ditemukan {
		fmt.Printf("\nTidak ada buku dengan rating %d\n", target)
	}
}

func main() {
	var koleksi KoleksiBuku
	var jumlahBuku int
	var pilihan int
	var ratingTarget int

	for {
		fmt.Println("\n=== Menu Perpustakaan Buku ===")
		fmt.Println("1. Tambahkan Buku")
		fmt.Println("2. Lihat Buku Favorit")
		fmt.Println("3. Urutkan Buku berdasarkan Rating")
		fmt.Println("4. Lihat 5 Buku dengan Rating Tertinggi")
		fmt.Println("5. Cari Buku berdasarkan Rating")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu (1-6): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			TambahkanBuku(&koleksi, &jumlahBuku)
		case 2:
			TampilkanFavorit(koleksi, jumlahBuku)
		case 3:
			SortirBuku(&koleksi, jumlahBuku)
			fmt.Println("Buku berhasil diurutkan berdasarkan rating")
		case 4:
			CetakTop5(koleksi, jumlahBuku)
		case 5:
			fmt.Print("Masukkan rating yang ingin dicari: ")
			fmt.Scan(&ratingTarget)
			TemukanBuku(koleksi, jumlahBuku, ratingTarget)
		case 6:
			fmt.Println("Terima kasih telah menggunakan program koleksi buku.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Coba lagi.")
		}
	}
}
