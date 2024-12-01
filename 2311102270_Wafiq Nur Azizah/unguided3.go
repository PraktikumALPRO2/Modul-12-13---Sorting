package main

import "fmt"

const kapasitasMaksimal = 7919

type Buku struct {
	kode, judul, penulis, penerbit string
	stok, tahun, rating            int
}

type DaftarBuku []Buku

func tambahBukuBaru(koleksi *DaftarBuku, total *int) {
	var bukuBaru Buku

	fmt.Print("Masukkan Kode Buku: ")
	fmt.Scan(&bukuBaru.kode)
	fmt.Print("Masukkan Judul: ")
	fmt.Scan(&bukuBaru.judul)
	fmt.Print("Masukkan Penulis: ")
	fmt.Scan(&bukuBaru.penulis)
	fmt.Print("Masukkan Penerbit: ")
	fmt.Scan(&bukuBaru.penerbit)
	fmt.Print("Masukkan Jumlah Stok: ")
	fmt.Scan(&bukuBaru.stok)
	fmt.Print("Masukkan Tahun: ")
	fmt.Scan(&bukuBaru.tahun)
	fmt.Print("Masukkan Rating: ")
	fmt.Scan(&bukuBaru.rating)

	*koleksi = append(*koleksi, bukuBaru)
	*total++
}

func tampilkanBukuFavorit(koleksi DaftarBuku, total int) {
	if total == 0 {
		fmt.Println("Tidak ada buku dalam koleksi")
		return
	}

	indeksTerbaik := 0
	for i := 1; i < total; i++ {
		if koleksi[i].rating > koleksi[indeksTerbaik].rating {
			indeksTerbaik = i
		}
	}

	fmt.Println("\nBuku Favorit:")
	fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nRating: %d\n",
		koleksi[indeksTerbaik].judul, koleksi[indeksTerbaik].penulis,
		koleksi[indeksTerbaik].penerbit, koleksi[indeksTerbaik].tahun,
		koleksi[indeksTerbaik].rating)
}

func urutkanBukuRating(koleksi *DaftarBuku, total int) {
	for i := 1; i < total; i++ {
		kunci := (*koleksi)[i]
		j := i - 1
		for j >= 0 && (*koleksi)[j].rating < kunci.rating {
			(*koleksi)[j+1] = (*koleksi)[j]
			j--
		}
		(*koleksi)[j+1] = kunci
	}
}

func tampilkan5BukuTerbaik(koleksi DaftarBuku, total int) {
	fmt.Println("\n5 Buku dengan Rating Tertinggi:")
	batas := 5
	if total < 5 {
		batas = total
	}
	for i := 0; i < batas; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, koleksi[i].judul, koleksi[i].rating)
	}
}

func cariBukuRating(koleksi DaftarBuku, total int, ratingDicari int) {
	kiri, kanan := 0, total-1
	ditemukan := false

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if koleksi[tengah].rating == ratingDicari {
			fmt.Printf("\nBuku dengan rating %d ditemukan:\n", ratingDicari)
			fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nStok: %d\n",
				koleksi[tengah].judul, koleksi[tengah].penulis,
				koleksi[tengah].penerbit, koleksi[tengah].tahun,
				koleksi[tengah].stok)
			ditemukan = true
			break
		}
		if koleksi[tengah].rating < ratingDicari {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if !ditemukan {
		fmt.Printf("\nTidak ada buku dengan rating %d\n", ratingDicari)
	}
}

func main() {
	var koleksiBuku DaftarBuku
	var totalBuku int
	var pilihanMenu, ratingDicari int

	for {
		fmt.Println("\n=== Sistem Manajemen Buku ===")
		fmt.Println("1. Tambah Buku Baru")
		fmt.Println("2. Tampilkan Buku Favorit")
		fmt.Println("3. Urutkan Buku berdasarkan Rating")
		fmt.Println("4. Tampilkan 5 Buku Terbaik")
		fmt.Println("5. Cari Buku berdasarkan Rating")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu (1-6): ")
		fmt.Scan(&pilihanMenu)

		switch pilihanMenu {
		case 1:
			tambahBukuBaru(&koleksiBuku, &totalBuku)
		case 2:
			tampilkanBukuFavorit(koleksiBuku, totalBuku)
		case 3:
			urutkanBukuRating(&koleksiBuku, totalBuku)
			fmt.Println("Buku telah diurutkan berdasarkan rating")
		case 4:
			tampilkan5BukuTerbaik(koleksiBuku, totalBuku)
		case 5:
			fmt.Print("Masukkan rating yang dicari: ")
			fmt.Scan(&ratingDicari)
			cariBukuRating(koleksiBuku, totalBuku, ratingDicari)
		case 6:
			fmt.Println("Terima kasih telah menggunakan sistem manajemen buku")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
