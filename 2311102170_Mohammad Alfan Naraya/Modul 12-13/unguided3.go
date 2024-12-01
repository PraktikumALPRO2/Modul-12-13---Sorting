package main

import (
	"fmt"
)

const nMax = 7919

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

// Struct untuk menyimpan daftar buku
type DaftarBuku []Buku

// Fungsi untuk mendaftarkan buku baru
func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
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
	*n++
}

// Fungsi untuk menampilkan buku terfavorit
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada buku dalam perpustakaan")
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

	fmt.Println("\nBuku Terfavorit:")
	fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nRating: %d\n",
		pustaka[maxIndex].judul, pustaka[maxIndex].penulis,
		pustaka[maxIndex].penerbit, pustaka[maxIndex].tahun,
		pustaka[maxIndex].rating)
}

// Fungsi untuk mengurutkan buku berdasarkan rating (Insertion Sort)
func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := (*pustaka)[i]
		j := i - 1
		for j >= 0 && (*pustaka)[j].rating < key.rating {
			(*pustaka)[j+1] = (*pustaka)[j]
			j--
		}
		(*pustaka)[j+1] = key
	}
}

// Fungsi untuk mencetak 5 buku dengan rating tertinggi
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("\n5 Buku dengan Rating Tertinggi:")
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
	}
}

// Fungsi untuk mencari buku berdasarkan rating (Binary Search)
func CariBuku(pustaka DaftarBuku, n int, targetRating int) {
	left := 0
	right := n - 1
	found := false

	for left <= right {
		mid := (left + right) / 2
		if pustaka[mid].rating == targetRating {
			fmt.Printf("\nBuku dengan rating %d ditemukan:\n", targetRating)
			fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nEksemplar: %d\n",
				pustaka[mid].judul, pustaka[mid].penulis,
				pustaka[mid].penerbit, pustaka[mid].tahun,
				pustaka[mid].eksemplar)
			found = true
			break
		}
		if pustaka[mid].rating < targetRating {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if !found {
		fmt.Printf("\nTidak ada buku dengan rating %d\n", targetRating)
	}
}

func main() {
	var pustaka DaftarBuku
	var n int
	var pilihan int
	var targetRating int

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
			DaftarkanBuku(&pustaka, &n)
		case 2:
			CetakTerfavorit(pustaka, n)
		case 3:
			UrutBuku(&pustaka, n)
			fmt.Println("Buku telah diurutkan berdasarkan rating")
		case 4:
			Cetak5Terbaru(pustaka, n)
		case 5:
			fmt.Print("Masukkan rating yang dicari: ")
			fmt.Scan(&targetRating)
			CariBuku(pustaka, n, targetRating)
		case 6:
			fmt.Println("Terima kasih telah menggunakan program perpustakaan")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
