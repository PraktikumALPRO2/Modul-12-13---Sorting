package main

import (
	"fmt"
)

// Definisi struct Buku
type Buku struct {
	ID       int
	Judul    string
	Penulis  string
	Penerbit string
	Tahun    int
	Rating   int
}

// Definisi tipe array untuk menyimpan data buku
const nMax = 7919

type DaftarBuku [nMax]Buku

// Subprogram untuk memasukkan data buku
func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (ID, Judul, Penulis, Penerbit, Tahun, Rating):\n", i+1)
		fmt.Scan(&pustaka[i].ID, &pustaka[i].Judul, &pustaka[i].Penulis, &pustaka[i].Penerbit, &pustaka[i].Tahun, &pustaka[i].Rating)
	}
}

// Subprogram untuk mencetak buku dengan rating tertinggi
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada buku yang terdaftar.")
		return
	}

	// Cari buku dengan rating tertinggi
	tertinggi := pustaka[0].Rating
	var index int
	for i := 1; i < n; i++ {
		if pustaka[i].Rating > tertinggi {
			tertinggi = pustaka[i].Rating
			index = i
		}
	}

	fmt.Println("Buku dengan rating tertinggi:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
		pustaka[index].Judul, pustaka[index].Penulis, pustaka[index].Penerbit, pustaka[index].Tahun, pustaka[index].Rating)
}

// Subprogram untuk mengurutkan data buku berdasarkan rating (Insertion Sort)
func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1

		// Geser elemen yang lebih kecil dari key ke kanan
		for j >= 0 && pustaka[j].Rating < key.Rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

// Subprogram untuk mencetak lima buku dengan rating tertinggi
func Cetak5Terbaik(pustaka DaftarBuku, n int) {
	fmt.Println("Lima buku dengan rating tertinggi:")
	for i := 0; i < 5 && i < n; i++ {
		fmt.Printf("%d. Judul: %s, Rating: %d\n", i+1, pustaka[i].Judul, pustaka[i].Rating)
	}
}

// Subprogram untuk mencari buku berdasarkan rating tertentu
func CariBuku(pustaka DaftarBuku, n int, ratingCari int) {
	fmt.Printf("Mencari buku dengan rating: %d\n", ratingCari)
	found := false
	for i := 0; i < n; i++ {
		if pustaka[i].Rating == ratingCari {
			fmt.Printf("Ditemukan buku: Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
				pustaka[i].Judul, pustaka[i].Penulis, pustaka[i].Penerbit, pustaka[i].Tahun, pustaka[i].Rating)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var pustaka DaftarBuku
	var n_2311102040 int

	// Memasukkan data buku
	DaftarkanBuku(&pustaka, &n_2311102040)

	// Mengurutkan buku berdasarkan rating
	UrutBuku(&pustaka, n_2311102040)

	// Menampilkan buku dengan rating tertinggi
	CetakTerfavorit(pustaka, n_2311102040)

	// Menampilkan lima buku dengan rating tertinggi
	Cetak5Terbaik(pustaka, n_2311102040)

	// Mencari buku dengan rating tertentu
	var ratingCari int
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&ratingCari)
	CariBuku(pustaka, n_2311102040, ratingCari)
}
