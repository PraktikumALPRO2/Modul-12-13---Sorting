package main

import (
	"fmt"
	"sort"
)

// Struct untuk menyimpan data buku
type Buku struct {
	id        string
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

// Fungsi untuk mendaftarkan buku baru
func DaftarkanBuku(pustaka *[]Buku) {
	var buku Buku
	fmt.Print("Masukkan ID: ")
	fmt.Scan(&buku.id)
	fmt.Print("Masukkan Judul: ")
	fmt.Scan(&buku.judul)
	fmt.Print("Masukkan Penulis: ")
	fmt.Scan(&buku.penulis)
	fmt.Print("Masukkan Penerbit: ")
	fmt.Scan(&buku.penerbit)
	fmt.Print("Masukkan Jumlah Eksemplar: ")
	fmt.Scan(&buku.eksemplar)
	fmt.Print("Masukkan Tahun: ")
	fmt.Scan(&buku.tahun)
	fmt.Print("Masukkan Rating: ")
	fmt.Scan(&buku.rating)

	*pustaka = append(*pustaka, buku)
}

// Fungsi untuk menampilkan buku terfavorit
func CetakTerfavorit(pustaka []Buku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada buku dalam perpustakaan")
		return
	}

	terfavorit := pustaka[0]
	for _, buku := range pustaka {
		if buku.rating > terfavorit.rating {
			terfavorit = buku
		}
	}

	fmt.Println("\nBuku Terfavorit:")
	fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nRating: %d\n",
		terfavorit.judul, terfavorit.penulis, terfavorit.penerbit,
		terfavorit.tahun, terfavorit.rating)
}

// Fungsi untuk mengurutkan buku berdasarkan rating
func UrutBuku(pustaka *[]Buku) {
	sort.Slice(*pustaka, func(i, j int) bool {
		return (*pustaka)[i].rating > (*pustaka)[j].rating
	})
}

// Fungsi untuk mencetak 5 buku dengan rating tertinggi
func Cetak5Terbaru(pustaka []Buku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada buku dalam perpustakaan")
		return
	}

	fmt.Println("\n5 Buku dengan Rating Tertinggi:")
	for i, buku := range pustaka {
		if i == 5 {
			break
		}
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, buku.judul, buku.rating)
	}
}

// Fungsi untuk mencari buku berdasarkan rating
func CariBuku(pustaka []Buku, targetRating int) {
	idx := sort.Search(len(pustaka), func(i int) bool {
		return pustaka[i].rating <= targetRating
	})

	if idx < len(pustaka) && pustaka[idx].rating == targetRating {
		buku := pustaka[idx]
		fmt.Printf("\nBuku dengan rating %d ditemukan:\n", targetRating)
		fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nEksemplar: %d\n",
			buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.eksemplar)
	} else {
		fmt.Printf("\nTidak ada buku dengan rating %d\n", targetRating)
	}
}

func main() {
	var pustaka []Buku
	var pilihan int

	for {
		fmt.Println("\n=== Menu Perpustakaan ===")
		fmt.Println("1. Daftarkan Buku")
		fmt.Println("2. Tampilkan Buku Terfavorit")
		fmt.Println("3. Urutkan Buku berdasarkan Rating")
		fmt.Println("4. Tampilkan 5 Buku Rating Tertinggi")
		fmt.Println("5. Cari Buku berdasarkan Rating")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu (1-6): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			DaftarkanBuku(&pustaka)
		case 2:
			CetakTerfavorit(pustaka)
		case 3:
			UrutBuku(&pustaka)
			fmt.Println("Buku telah diurutkan berdasarkan rating")
		case 4:
			Cetak5Terbaru(pustaka)
		case 5:
			if len(pustaka) == 0 {
				fmt.Println("Tidak ada buku dalam perpustakaan")
				continue
			}
			fmt.Print("Masukkan rating yang dicari: ")
			var targetRating int
			fmt.Scan(&targetRating)
			UrutBuku(&pustaka) // Pastikan buku terurut sebelum pencarian
			CariBuku(pustaka, targetRating)
		case 6:
			fmt.Println("Terima kasih telah menggunakan program perpustakaan")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
