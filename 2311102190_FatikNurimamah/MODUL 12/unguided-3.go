package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

// Struct untuk menyimpan data buku
type Buku struct {
	KodeBuku      string
	NamaBuku      string
	Pengarang     string
	Penerbit      string
	JumlahEksemplar int
	TahunTerbit   int
	NilaiRating   int
}

// Struct untuk menyimpan daftar buku
type KoleksiBuku struct {
	KumpulanBuku []Buku
	TotalBuku    int
}

// Procedure DaftarkanBuku
func DaftarkanBuku(pustaka *KoleksiBuku, n int) {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("\nMasukkan data untuk buku ke-%d:\n", i+1)

		// Input kode buku
		fmt.Print("Kode Buku: ")
		buku.KodeBuku, _ = reader.ReadString('\n')
		buku.KodeBuku = strings.TrimSpace(buku.KodeBuku)

		// Input nama buku
		fmt.Print("Nama Buku: ")
		buku.NamaBuku, _ = reader.ReadString('\n')
		buku.NamaBuku = strings.TrimSpace(buku.NamaBuku)

		// Input pengarang
		fmt.Print("Pengarang: ")
		buku.Pengarang, _ = reader.ReadString('\n')
		buku.Pengarang = strings.TrimSpace(buku.Pengarang)

		// Input penerbit
		fmt.Print("Penerbit: ")
		buku.Penerbit, _ = reader.ReadString('\n')
		buku.Penerbit = strings.TrimSpace(buku.Penerbit)

		// Input jumlah eksemplar
		fmt.Print("Jumlah Eksemplar: ")
		eksemplar, _ := reader.ReadString('\n')
		buku.JumlahEksemplar, _ = strconv.Atoi(strings.TrimSpace(eksemplar))

		// Input tahun terbit
		fmt.Print("Tahun Terbit: ")
		tahun, _ := reader.ReadString('\n')
		buku.TahunTerbit, _ = strconv.Atoi(strings.TrimSpace(tahun))

		// Input nilai rating
		fmt.Print("Nilai Rating: ")
		rating, _ := reader.ReadString('\n')
		buku.NilaiRating, _ = strconv.Atoi(strings.TrimSpace(rating))

		// Tambahkan buku ke dalam koleksi
		pustaka.KumpulanBuku = append(pustaka.KumpulanBuku, buku)
	}
	pustaka.TotalBuku = n
}

// Procedure CetakTerfavorit
func CetakTerfavorit(pustaka KoleksiBuku) {
	if len(pustaka.KumpulanBuku) == 0 {
		fmt.Println("Tidak ada buku dalam koleksi.")
		return
	}

	maxRating := -1
	var favorit Buku
	for _, buku := range pustaka.KumpulanBuku {
		if buku.NilaiRating > maxRating {
			maxRating = buku.NilaiRating
			favorit = buku
		}
	}

	fmt.Printf("Buku terfavorit:\nNama: %s, Pengarang: %s, Penerbit: %s, Tahun: %d\n",
		favorit.NamaBuku, favorit.Pengarang, favorit.Penerbit, favorit.TahunTerbit)
}

// Procedure UrutBuku (menggunakan Insertion Sort)
func UrutBuku(pustaka *KoleksiBuku) {
	sort.Slice(pustaka.KumpulanBuku, func(i, j int) bool {
		return pustaka.KumpulanBuku[i].NilaiRating > pustaka.KumpulanBuku[j].NilaiRating
	})
}

// Procedure Cetak5Terbaru
func Cetak5Terbaru(pustaka KoleksiBuku) {
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	limit := 5
	if len(pustaka.KumpulanBuku) < 5 {
		limit = len(pustaka.KumpulanBuku)
	}
	for i := 0; i < limit; i++ {
		buku := pustaka.KumpulanBuku[i]
		fmt.Printf("%d. Nama: %s, Rating: %d\n", i+1, buku.NamaBuku, buku.NilaiRating)
	}
}

// Procedure CariBuku
func CariBuku(pustaka KoleksiBuku, r int) {
	for _, buku := range pustaka.KumpulanBuku {
		if buku.NilaiRating == r {
			fmt.Printf("Buku ditemukan:\nNama: %s, Pengarang: %s, Penerbit: %s, Tahun: %d\n",
				buku.NamaBuku, buku.Pengarang, buku.Penerbit, buku.TahunTerbit)
			return
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka KoleksiBuku
	var jumlahBuku int

	// Masukkan data buku
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scanln(&jumlahBuku)
	DaftarkanBuku(&pustaka, jumlahBuku)

	// Cetak buku terfavorit
	CetakTerfavorit(pustaka)

	// Urutkan buku berdasarkan rating
	UrutBuku(&pustaka)

	// Cetak 5 buku dengan rating tertinggi
	Cetak5Terbaru(pustaka)

	// Cari buku berdasarkan rating
	var nilaiRating int
	fmt.Print("Masukkan rating untuk mencari buku: ")
	fmt.Scanln(&nilaiRating)
	CariBuku(pustaka, nilaiRating)
}
