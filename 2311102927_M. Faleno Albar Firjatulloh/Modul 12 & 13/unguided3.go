package main

import "fmt"

const maksBuku = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type KoleksiBuku []Buku

func daftarkanBukuBaru(perpustakaan *KoleksiBuku, jumlah *int) {
	var buku Buku

	fmt.Print("Masukkan ID Buku: ")
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

	*perpustakaan = append(*perpustakaan, buku)
	*jumlah++
}

func cetakBukuTerfavorit(perpustakaan KoleksiBuku, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Tidak ada buku di perpustakaan")
		return
	}

	indeksMaksimal := 0
	for i := 1; i < jumlah; i++ {
		if perpustakaan[i].rating > perpustakaan[indeksMaksimal].rating {
			indeksMaksimal = i
		}
	}

	fmt.Println("\nBuku Terfavorit:")
	fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nRating: %d\n",
		perpustakaan[indeksMaksimal].judul, perpustakaan[indeksMaksimal].penulis,
		perpustakaan[indeksMaksimal].penerbit, perpustakaan[indeksMaksimal].tahun,
		perpustakaan[indeksMaksimal].rating)
}

func urutkanBukuBerdasarkanRating(perpustakaan *KoleksiBuku, jumlah int) {
	for i := 1; i < jumlah; i++ {
		kunci := (*perpustakaan)[i]
		j := i - 1
		for j >= 0 && (*perpustakaan)[j].rating < kunci.rating {
			(*perpustakaan)[j+1] = (*perpustakaan)[j]
			j--
		}
		(*perpustakaan)[j+1] = kunci
	}
}

func cetak5BukuTeratas(perpustakaan KoleksiBuku, jumlah int) {
	fmt.Println("\n5 Buku dengan Rating Tertinggi:")
	batasAtas := 5
	if jumlah < 5 {
		batasAtas = jumlah
	}
	for i := 0; i < batasAtas; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, perpustakaan[i].judul, perpustakaan[i].rating)
	}
}

func cariBukuBerdasarkanRating(perpustakaan KoleksiBuku, jumlah int, ratingTarget int) {
	kiri, kanan := 0, jumlah-1
	ditemukan := false

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if perpustakaan[tengah].rating == ratingTarget {
			fmt.Printf("\nBuku dengan rating %d ditemukan:\n", ratingTarget)
			fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nEksemplar: %d\n",
				perpustakaan[tengah].judul, perpustakaan[tengah].penulis,
				perpustakaan[tengah].penerbit, perpustakaan[tengah].tahun,
				perpustakaan[tengah].eksemplar)
			ditemukan = true
			break
		}
		if perpustakaan[tengah].rating < ratingTarget {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if !ditemukan {
		fmt.Printf("\nTidak ada buku dengan rating %d\n", ratingTarget)
	}
}

func main() {
	var perpustakaan KoleksiBuku
	var jumlah int
	var pilihan, ratingTarget int

	for {
		fmt.Println("\n=== Manajemen Perpustakaan ===")
		fmt.Println("1. Daftarkan Buku Baru")
		fmt.Println("2. Tampilkan Buku Terfavorit")
		fmt.Println("3. Urutkan Buku berdasarkan Rating")
		fmt.Println("4. Tampilkan 5 Buku Teratas")
		fmt.Println("5. Cari Buku berdasarkan Rating")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu (1-6): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			daftarkanBukuBaru(&perpustakaan, &jumlah)
		case 2:
			cetakBukuTerfavorit(perpustakaan, jumlah)
		case 3:
			urutkanBukuBerdasarkanRating(&perpustakaan, jumlah)
			fmt.Println("Buku telah diurutkan berdasarkan rating")
		case 4:
			cetak5BukuTeratas(perpustakaan, jumlah)
		case 5:
			fmt.Print("Masukkan rating yang dicari: ")
			fmt.Scan(&ratingTarget)
			cariBukuBerdasarkanRating(perpustakaan, jumlah, ratingTarget)
		case 6:
			fmt.Println("Terima kasih telah menggunakan program manajemen perpustakaan")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
