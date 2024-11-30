package main
import "fmt"

const Max = 7919
type Buku struct {
	id        int
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}
type DaftarBuku struct {
	Pustaka []Buku
	nPustaka int
}

func DaftarkanBuku(pustaka *DaftarBuku, buku Buku) {
	if pustaka.nPustaka < Max {
		pustaka.Pustaka = append(pustaka.Pustaka, buku)
		pustaka.nPustaka++
	}
}

func CetakFavorit(pustaka DaftarBuku) {
	fmt.Println("Daftar Buku:")
	for _, buku := range pustaka.Pustaka {
		fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Rating: %d\n",
			buku.id, buku.judul, buku.penulis, buku.penerbit, buku.rating)
	}
}

func UrutBuku(pustaka *DaftarBuku) {
	for i := 1; i < pustaka.nPustaka; i++ {
		key := pustaka.Pustaka[i]
		j := i - 1
		for j >= 0 && pustaka.Pustaka[j].rating < key.rating {
			pustaka.Pustaka[j+1] = pustaka.Pustaka[j]
			j--
		}
		pustaka.Pustaka[j+1] = key
	}
}

func Cetak5Tertinggi(pustaka DaftarBuku) {
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 0; i < 5 && i < pustaka.nPustaka; i++ {
		buku := pustaka.Pustaka[i]
		fmt.Printf("Judul: %s, Penulis: %s, Rating: %d\n", buku.judul, buku.penulis, buku.rating)
	}
}

func CariBuku(pustaka DaftarBuku, rating int) {
	fmt.Printf("Buku dengan Rating %d:\n", rating)
	found := false
	for _, buku := range pustaka.Pustaka {
		if buku.rating == rating {
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n",
				buku.judul, buku.penulis, buku.penerbit, buku.tahun)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var pustaka DaftarBuku
	var n int
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("Masukkan data buku ke-%d (ID Judul Penulis Penerbit Eksemplar Tahun Rating):\n", i+1)
		fmt.Scan(&buku.id, &buku.judul, &buku.penulis, &buku.penerbit, &buku.eksemplar, &buku.tahun, &buku.rating)
		DaftarkanBuku(&pustaka, buku)
	}

	CetakFavorit(pustaka)
	UrutBuku(&pustaka)
	Cetak5Tertinggi(pustaka)
	var rating int
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scanln(&rating)
	CariBuku(pustaka, rating)
}
