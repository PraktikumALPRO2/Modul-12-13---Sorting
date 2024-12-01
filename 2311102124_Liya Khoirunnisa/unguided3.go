/* Liya Khoirunisa - 2311102124 */
package main

import (
	"fmt"
)

// Nilai maksimum
const nMax int = 7919

// Struct buku
type Buku struct {
	id        string
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

// Slice dari struct buku
type DaftarBuku []Buku

// Fungsi untuk mendaftarkan buku baru ke pustaka
func daftarkanBuku_124(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		var buku Buku

		// Deklarasi buku
		fmt.Printf("Masukkan data buku ke-%d:\n", i+1)
		fmt.Print("ID buku          : ")
		fmt.Scan(&buku.id)
		fmt.Print("Judul buku       : ")
		fmt.Scan(&buku.judul)
		fmt.Print("Penulis buku     : ")
		fmt.Scan(&buku.penulis)
		fmt.Print("Penerbit buku    : ")
		fmt.Scan(&buku.penerbit)
		fmt.Print("Eksemplar buku   : ")
		fmt.Scan(&buku.eksemplar)
		fmt.Print("Tahun terbit buku: ")
		fmt.Scan(&buku.tahun)
		fmt.Print("Rating buku      : ")
		fmt.Scan(&buku.rating)
		*pustaka = append(*pustaka, buku)
	}
}

// Fungsi untuk mencetak buku dengan rating tertinggi
func cetakTerfavorit_124(pustaka DaftarBuku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}
	maxRating := pustaka[0]

	// Mencari buku dengan rating tertinggi
	for _, buku := range pustaka {
		if buku.rating > maxRating.rating {
			maxRating = buku
		}
	}

	// Menampilkan buku terfavorit
	fmt.Println("\n///// Buku Terfavorit /////")
	fmt.Printf("Judul buku       : %s\n", maxRating.judul)
	fmt.Printf("Penulis buku     : %s\n", maxRating.penulis)
	fmt.Printf("Penerbit buku    : %s\n", maxRating.penerbit)
	fmt.Printf("Tahun terbit buku: %d\n", maxRating.tahun)
}

// Fungsi untuk mengurutkan pustaka berdasarkan rating
func urutBuku_124(pustaka *DaftarBuku) {
	for i := 1; i < len(*pustaka); i++ {
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
func cetak5Terbaru_124(pustaka DaftarBuku) {
	fmt.Println("\n///// 5 Buku dengan Rating Tertinggi /////")
	for i := 0; i < len(pustaka) && i < 5; i++ {
		fmt.Printf("%d. Judul: %s, Rating: %d\n", i+1, pustaka[i].judul, pustaka[i].rating)
	}
}

// FUngsi untuk mencari buku berdasarkan rating
func cariBuku_124(pustaka DaftarBuku, r int) {
	low, high := 0, len(pustaka)-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			buku := pustaka[mid]
			fmt.Printf("\nBuku ditemukan:\n")
			fmt.Printf("Judul buku       : %s\n", buku.judul)
			fmt.Printf("Penulis buku     : %s\n", buku.penulis)
			fmt.Printf("Penerbit buku    : %s\n", buku.penerbit)
			fmt.Printf("Tahun terbit buku: %d\n", buku.tahun)
			fmt.Printf("Eksemplar buku   : %d\n", buku.eksemplar)
			fmt.Printf("Rating buku      : %d\n", buku.rating)
			return
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("\nTidak ada buku dengan rating seperti itu.")
}

func main() {
	// Deklarasi variabel
	var pustaka DaftarBuku
	var n, rating int

	// Meminta input jumlah buku
	fmt.Print("Masukkan jumlah buku yang ingin didaftarkan: ")
	fmt.Scan(&n)

	// Panggil fungsi untuk mendaftarkan buku
	daftarkanBuku_124(&pustaka, n)

	// Panggil fungsi untuk mengurutkan buku
	urutBuku_124(&pustaka)

	// Panggil fungsi untuk menampilkan buku dengan rating tertinggi
	cetakTerfavorit_124(pustaka)

	// Panggil fungsi untuk menampilkan 5 buku dengan rating tertinggi
	cetak5Terbaru_124(pustaka)

	// Meminta input rating yang ingin dicari
	fmt.Print("\nMasukkan rating yang ingin dicari: ")
	fmt.Scan(&rating)

	// Panggil fungsi untuk mencari buku berdasarkan rating
	cariBuku_124(pustaka, rating)
}
