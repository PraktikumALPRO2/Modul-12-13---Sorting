// Haifa Zahra Azzimmi
// 2311102163
// S1 IF 11 5

package main

import "fmt"

const maxBooks = 7919

// Struktur Buku untuk menyimpan informasi buku
type Buku struct {
	kode, judul, penulis, penerbit string
	eksemplar, tahun, rating       int
}

// Array KoleksiBuku untuk menyimpan daftar buku
type KoleksiBuku [maxBooks]Buku

// Fungsi untuk menambahkan buku ke dalam koleksi
func tambahBuku(koleksi *KoleksiBuku, jumlah *int) {
	fmt.Print("Masukkan jumlah buku yang ingin didaftarkan: ")
	fmt.Scan(jumlah)

	for i := 0; i < *jumlah; i++ {
		fmt.Printf("\nMasukkan data buku ke-%d (kode, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scan(&koleksi[i].kode, &koleksi[i].judul, &koleksi[i].penulis, &koleksi[i].penerbit,
			&koleksi[i].eksemplar, &koleksi[i].tahun, &koleksi[i].rating)
	}
}

// Fungsi untuk menampilkan buku dengan rating tertinggi
func bukuFavorit(koleksi KoleksiBuku, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Tidak ada buku yang terdaftar.")
		return
	}

	favorit := koleksi[0]
	for i := 1; i < jumlah; i++ {
		if koleksi[i].rating > favorit.rating {
			favorit = koleksi[i]
		}
	}

	fmt.Println("\nBuku Favorit:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
		favorit.judul, favorit.penulis, favorit.penerbit, favorit.tahun, favorit.rating)
}

// Fungsi untuk mengurutkan buku berdasarkan rating secara menurun
func urutkanBuku(koleksi *KoleksiBuku, jumlah int) {
	for i := 1; i < jumlah; i++ {
		buku := koleksi[i]
		j := i - 1

		for j >= 0 && koleksi[j].rating < buku.rating {
			koleksi[j+1] = koleksi[j]
			j--
		}
		koleksi[j+1] = buku
	}
}

// Fungsi untuk menampilkan 5 buku dengan rating tertinggi
func tampilkanTop5(koleksi KoleksiBuku, jumlah int) {
	fmt.Println("\n5 Buku Dengan Rating Tertinggi:")
	batas := 5
	if jumlah < 5 {
		batas = jumlah
	}

	for i := 0; i < batas; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, koleksi[i].judul, koleksi[i].rating)
	}
}

// Fungsi untuk mencari buku berdasarkan rating
func cariBukuByRating(koleksi KoleksiBuku, jumlah, rating int) {
	ditemukan := false
	fmt.Printf("\nMencari Buku dengan Rating %d:\n", rating)

	for i := 0; i < jumlah; i++ {
		if koleksi[i].rating == rating {
			ditemukan = true
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
				koleksi[i].judul, koleksi[i].penulis, koleksi[i].penerbit,
				koleksi[i].tahun, koleksi[i].eksemplar, koleksi[i].rating)
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var koleksi KoleksiBuku
	var jumlahBuku, ratingPencarian int

	// Menambahkan buku ke dalam koleksi
	tambahBuku(&koleksi, &jumlahBuku)

	// Menampilkan buku dengan rating tertinggi
	bukuFavorit(koleksi, jumlahBuku)

	// Mengurutkan buku berdasarkan rating dan menampilkan hasilnya
	urutkanBuku(&koleksi, jumlahBuku)

	// Menampilkan 5 buku dengan rating tertinggi
	tampilkanTop5(koleksi, jumlahBuku)

	// Mencari buku berdasarkan rating
	fmt.Print("\nMasukkan rating yang ingin dicari: ")
	fmt.Scan(&ratingPencarian)
	cariBukuByRating(koleksi, jumlahBuku, ratingPencarian)
}
