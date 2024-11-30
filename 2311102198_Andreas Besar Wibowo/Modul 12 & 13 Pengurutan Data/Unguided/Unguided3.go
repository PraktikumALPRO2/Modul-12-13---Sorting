// Andreas Besar Wibowo
// 2311102198 / IF-11-05

package main

import "fmt"

const nMax = 7919

type Buku struct {
	id       string
	judul    string
	penulis  string
	penerbit string
	eksemplar, tahun, rating int
}

type DaftarBuku [nMax]Buku

// Prosedur untuk mendaftarkan buku ke perpustakaan
func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Printf("\nMasukkan data buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit,
			&pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

// Prosedur untuk menampilkan buku dengan rating tertinggi
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}

	terfavorit := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > terfavorit.rating {
			terfavorit = pustaka[i]
		}
	}

	fmt.Println("\nBuku Terfavorit:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
		terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun, terfavorit.rating)
}

// Prosedur untuk mengurutkan buku berdasarkan rating secara descending
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

// Prosedur untuk mencetak 5 buku dengan rating tertinggi
func CetakTerbaru(pustaka DaftarBuku, n int) {
	fmt.Println("\n5 Judul Buku Dengan Rating Tertinggi:")
	count := 5
	if n < 5 {
		count = n
	}

	for i := 0; i < count; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
	}
}

// Prosedur untuk mencari buku berdasarkan ratingnya
func CariBuku(pustaka DaftarBuku, n, r int) {
	found := false
	fmt.Printf("\nBuku dengan Rating %d:\n", r)

	for i := 0; i < n; i++ {
		if pustaka[i].rating == r {
			found = true
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
				pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit,
				pustaka[i].tahun, pustaka[i].eksemplar, pustaka[i].rating)
		}
	}

	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu.")
	}
}

func main() {
	var pustaka DaftarBuku
	var n int
	var ratingCari int

	// Untuk menjalankan Prosedur Daftarkan Buku
	DaftarkanBuku(&pustaka, &n)

	// Untuk menjalankan Prosedur Cetak Buku Terfavorit
	CetakTerfavorit(pustaka, n)

	// Untuk menjalankan Prosedur Urutkan Buku berdasarkan Rating
	UrutBuku(&pustaka, n)

	// Untuk menjalankan Prosedur Cetak 5 buku dengan rating tertinggi
	CetakTerbaru(pustaka, n)

	// Untuk mencari buku berdasarkan rating
	fmt.Print("\nMasukkan rating yang ingin dicari: ")
	fmt.Scan(&ratingCari)
	CariBuku(pustaka, n, ratingCari)
}
