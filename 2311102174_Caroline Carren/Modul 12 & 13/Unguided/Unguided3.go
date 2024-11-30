// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

const nMax = 10000

type Buku struct {
	id                       string
	judul                    string
	penulis                  string
	penerbit                 string
	eksemplar, tahun, rating int
}

type DaftarBuku [nMax]Buku

// Warna ANSI
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	white  = "\033[37m"
)

// Prosedur untuk mendaftarkan buku ke perpustakaan
func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Print(yellow + "Masukkan jumlah buku yang ingin didaftarkan: " + reset)
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Printf("\n"+blue+"Masukkan data buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n"+reset, i+1)
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit,
			&pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

// Prosedur untuk menampilkan buku dengan rating tertinggi
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println(red + "\nTidak ada data buku." + reset)
		return
	}

	terfavorit := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > terfavorit.rating {
			terfavorit = pustaka[i]
		}
	}

	fmt.Println("\n" + green + "=============================================" + reset)
	fmt.Println(green + "Buku Terfavorit:" + reset)
	fmt.Println(green + "=============================================" + reset)
	fmt.Printf("Judul    : "+cyan+"%s\n"+reset, terfavorit.judul)
	fmt.Printf("Penulis  : "+cyan+"%s\n"+reset, terfavorit.penulis)
	fmt.Printf("Penerbit : "+cyan+"%s\n"+reset, terfavorit.penerbit)
	fmt.Printf("Tahun    : "+cyan+"%d\n"+reset, terfavorit.tahun)
	fmt.Printf("Rating   : "+green+"%d\n"+reset, terfavorit.rating)
	fmt.Println(green + "=============================================" + reset)
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
	fmt.Println("\n" + green + "=============================================" + reset)
	fmt.Println(green + "5 Buku Dengan Rating Tertinggi:" + reset)
	fmt.Println(green + "=============================================" + reset)
	count := 5
	if n < 5 {
		count = n
	}

	for i := 0; i < count; i++ {
		fmt.Printf("%d. "+cyan+"%s (Rating: %d)\n"+reset, i+1, pustaka[i].judul, pustaka[i].rating)
	}
	fmt.Println(green + "=============================================" + reset)
}

// Prosedur untuk mencari buku berdasarkan ratingnya
func CariBuku(pustaka DaftarBuku, n, r int) {
	found := false
	fmt.Printf("\n"+yellow+"Mencari Buku dengan Rating %d:\n"+reset, r)

	for i := 0; i < n; i++ {
		if pustaka[i].rating == r {
			found = true
			fmt.Printf("\n"+purple+"Judul    : %s\n"+reset, pustaka[i].judul)
			fmt.Printf("Penulis  : %s\n", pustaka[i].penulis)
			fmt.Printf("Penerbit : %s\n", pustaka[i].penerbit)
			fmt.Printf("Tahun    : %d\n", pustaka[i].tahun)
			fmt.Printf("Eksemplar: %d\n", pustaka[i].eksemplar)
			fmt.Printf("Rating   : "+green+"%d\n"+reset, pustaka[i].rating)
			fmt.Println("=============================================")
		}
	}

	if !found {
		fmt.Println(red + "\nTidak ada buku dengan rating tersebut." + reset)
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
	fmt.Print("\n" + yellow + "Masukkan rating yang ingin dicari: " + reset)
	fmt.Scan(&ratingCari)
	CariBuku(pustaka, n, ratingCari)
}
