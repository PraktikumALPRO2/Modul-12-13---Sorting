// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
	"sort"
)

// Definisi Struct Buku
type Buku struct {
	id        int
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

// Daftar Buku
type DaftarBuku []Buku

// Fungsi untuk mencatat data buku
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("Masukkan data buku ke-%d:\n", i+1)
		fmt.Print("ID: ")
		fmt.Scan(&buku.id)
		fmt.Print("Judul: ")
		fmt.Scan(&buku.judul)
		fmt.Print("Penulis: ")
		fmt.Scan(&buku.penulis)
		fmt.Print("Penerbit: ")
		fmt.Scan(&buku.penerbit)
		fmt.Print("Eksemplar: ")
		fmt.Scan(&buku.eksemplar)
		fmt.Print("Tahun: ")
		fmt.Scan(&buku.tahun)
		fmt.Print("Rating: ")
		fmt.Scan(&buku.rating)
		*pustaka = append(*pustaka, buku)
	}
}

// Fungsi untuk menampilkan buku dengan rating tertinggi
func CetakTerfavorit(pustaka DaftarBuku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada buku.")
		return
	}

	// Mencari buku dengan rating tertinggi
	terfavorit := pustaka[0]
	for _, buku := range pustaka {
		if buku.rating > terfavorit.rating {
			terfavorit = buku
		}
	}

	// Menampilkan data buku
	fmt.Println("Buku dengan rating tertinggi:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
		terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun, terfavorit.rating)
}

// Fungsi untuk mengurutkan buku berdasarkan rating (Descending)
func UrutkanBuku(pustaka *DaftarBuku) {
	sort.Slice(*pustaka, func(i, j int) bool {
		return (*pustaka)[i].rating > (*pustaka)[j].rating
	})
}

// Fungsi untuk menampilkan 5 buku dengan rating tertinggi
func Cetak5Terbaru(pustaka DaftarBuku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada buku.")
		return
	}

	fmt.Println("5 Buku dengan rating tertinggi:")
	for i := 0; i < len(pustaka) && i < 5; i++ {
		buku := pustaka[i]
		fmt.Printf("%d. Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
			i+1, buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.rating)
	}
}

// Fungsi untuk mencari buku berdasarkan rating
func CariBuku(pustaka DaftarBuku, rating int) {
	left, right := 0, len(pustaka)-1
	for left <= right {
		mid := (left + right) / 2
		if pustaka[mid].rating == rating {
			buku := pustaka[mid]
			fmt.Println("Buku yang ditemukan:")
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
				buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.rating)
			return
		} else if pustaka[mid].rating > rating {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	var n int

	// Input jumlah buku
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	// Input data buku
	DaftarkanBuku(&pustaka, n)

	// Urutkan buku berdasarkan rating
	UrutkanBuku(&pustaka)

	// Cetak buku dengan rating tertinggi
	CetakTerfavorit(pustaka)

	// Cetak 5 buku dengan rating tertinggi
	Cetak5Terbaru(pustaka)

	// Cari buku berdasarkan rating
	var rating int
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, rating)
}
