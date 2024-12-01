package main

import "fmt"

const nMax235 = 7919

type Buku235 struct {
	id235       string
	judul235    string
	penulis235  string
	penerbit235 string
	eksemplar235, tahun235, rating235 int
}

type DaftarBuku235 [nMax235]Buku235

// Prosedur untuk mendaftarkan buku ke perpustakaan
func DaftarkanBuku235(pustaka235 *DaftarBuku235, n235 *int) {
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(n235)

	for i235 := 0; i235 < *n235; i235++ {
		fmt.Printf("\nMasukkan data buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i235+1)
		fmt.Scan(&pustaka235[i235].id235, &pustaka235[i235].judul235, &pustaka235[i235].penulis235, &pustaka235[i235].penerbit235,
			&pustaka235[i235].eksemplar235, &pustaka235[i235].tahun235, &pustaka235[i235].rating235)
	}
}

// Prosedur untuk menampilkan buku dengan rating tertinggi
func CetakTerfavorit235(pustaka235 DaftarBuku235, n235 int) {
	if n235 == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}

	terfavorit235 := pustaka235[0]
	for i235 := 1; i235 < n235; i235++ {
		if pustaka235[i235].rating235 > terfavorit235.rating235 {
			terfavorit235 = pustaka235[i235]
		}
	}

	fmt.Println("\nBuku Terfavorit:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
		terfavorit235.judul235, terfavorit235.penulis235, terfavorit235.penerbit235, terfavorit235.tahun235, terfavorit235.rating235)
}

// Prosedur untuk mengurutkan buku berdasarkan rating secara descending
func UrutBuku235(pustaka235 *DaftarBuku235, n235 int) {
	for i235 := 1; i235 < n235; i235++ {
		key235 := pustaka235[i235]
		j235 := i235 - 1

		for j235 >= 0 && pustaka235[j235].rating235 < key235.rating235 {
			pustaka235[j235+1] = pustaka235[j235]
			j235--
		}
		pustaka235[j235+1] = key235
	}
}

// Prosedur untuk mencetak 5 buku dengan rating tertinggi
func CetakTerbaru235(pustaka235 DaftarBuku235, n235 int) {
	fmt.Println("\n5 Judul Buku Dengan Rating Tertinggi:")
	count235 := 5
	if n235 < 5 {
		count235 = n235
	}

	for i235 := 0; i235 < count235; i235++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i235+1, pustaka235[i235].judul235, pustaka235[i235].rating235)
	}
}

// Prosedur untuk mencari buku berdasarkan ratingnya
func CariBuku235(pustaka235 DaftarBuku235, n235, r235 int) {
	found235 := false
	fmt.Printf("\nBuku dengan Rating %d:\n", r235)

	for i235 := 0; i235 < n235; i235++ {
		if pustaka235[i235].rating235 == r235 {
			found235 = true
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
				pustaka235[i235].judul235, pustaka235[i235].penulis235, pustaka235[i235].penerbit235,
				pustaka235[i235].tahun235, pustaka235[i235].eksemplar235, pustaka235[i235].rating235)
		}
	}

	if !found235 {
		fmt.Println("Tidak ada buku dengan rating seperti itu.")
	}
}

func main() {
	var pustaka235 DaftarBuku235
	var n235 int
	var ratingCari235 int

	// Untuk menjalankan Prosedur Daftarkan Buku
	DaftarkanBuku235(&pustaka235, &n235)

	// Untuk menjalankan Prosedur Cetak Buku Terfavorit
	CetakTerfavorit235(pustaka235, n235)

	// Untuk menjalankan Prosedur Urutkan Buku berdasarkan Rating
	UrutBuku235(&pustaka235, n235)

	// Untuk menjalankan Prosedur Cetak 5 buku dengan rating tertinggi
	CetakTerbaru235(pustaka235, n235)

	// Untuk mencari buku berdasarkan rating
	fmt.Print("\nMasukkan rating yang ingin dicari: ")
	fmt.Scan(&ratingCari235)
	CariBuku235(pustaka235, n235, ratingCari235)
}
