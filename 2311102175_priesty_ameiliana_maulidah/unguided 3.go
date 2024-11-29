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

type DaftarBuku [7919]Buku

func main() {
	var n int
	fmt.Println("Masukkan jumlah buku:")
	fmt.Scanln(&n)

	if n < 1 || n > 7919 {
		fmt.Println("Jumlah buku harus di antara 1 hingga 7919.")
		return
	}

	var pustaka DaftarBuku

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating):\n", i+1)
		fmt.Scanf("%s %s %s %s %d %d %d\n",
			&pustaka[i].ID,
			&pustaka[i].Judul,
			&pustaka[i].Penulis,
			&pustaka[i].Penerbit,
			&pustaka[i].Eksemplar,
			&pustaka[i].Tahun,
			&pustaka[i].Rating)
	}

	var cariRating int
	fmt.Println("Masukkan rating yang akan dicari:")
	fmt.Scanln(&cariRating)

	fmt.Println("Buku dengan rating", cariRating, ":")
	found := false
	for i := 0; i < n; i++ {
		if pustaka[i].Rating == cariRating {
			fmt.Printf("ID: %s, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
				pustaka[i].ID, pustaka[i].Judul, pustaka[i].Penulis, pustaka[i].Penerbit, pustaka[i].Eksemplar, pustaka[i].Tahun, pustaka[i].Rating)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}
