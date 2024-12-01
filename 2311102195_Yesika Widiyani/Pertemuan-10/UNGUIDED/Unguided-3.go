// 2311102195 - Yesika Widiyani

package main

import (
	"fmt"
)

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

func DaftarkanBuku(pustaka *[]Buku, n int) {
	fmt.Println("Format masukan (id, judul, penulis, penerbit, eksamplar, tahun, rating):")
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("Buku ke-%d: ", i+1)
		fmt.Scan(&buku.id, &buku.judul, &buku.penulis, &buku.penerbit, &buku.eksemplar, &buku.tahun, &buku.rating)
		*pustaka = append(*pustaka, buku)
	}
}

func CetakTerfavorit(pustaka []Buku) {
	indexFav := 0
	for i := 1; i < len(pustaka); i++ {
		if pustaka[i].rating > pustaka[indexFav].rating {
			indexFav = i
		}
	}
	buku := pustaka[indexFav]
	fmt.Println("Buku terfavorit:")
	fmt.Printf("%s | %s | %s | %s | %d eksemplar | Tahun: %d | Rating: %d\n", 
		buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
}

func UrutBuku(pustaka *[]Buku) {
	n := len(*pustaka)
	for i := 1; i < n; i++ {
		key := (*pustaka)[i]
		j := i - 1
		for j >= 0 && (*pustaka)[j].rating < key.rating {
			(*pustaka)[j+1] = (*pustaka)[j]
			j--
		}
		(*pustaka)[j+1] = key
	}
}

func Cetak5Terbaru(pustaka []Buku) {
	fmt.Println("5 Buku dengan rating tertinggi:")
	n := len(pustaka)
	if n > 5 {
		n = 5
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
	}
}

func CariBuku(pustaka []Buku, r int) {
	low, high := 0, len(pustaka)-1
	found := false
	for low <= high && !found {
		mid := (low + high) / 2
		if pustaka[mid].rating > r {
			low = mid + 1
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			found = true
			fmt.Println("Buku dengan rating yang dicari:")
			buku := pustaka[mid]
			fmt.Printf("%s | %s | %s | %s | %d eksemplar | Tahun: %d | Rating: %d\n", 
				buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating yang dicari.")
	}
}

func main() {
	var pustaka []Buku
	var nPustaka, chooseRating int

	fmt.Print("Masukkan banyaknya buku: ")
	fmt.Scan(&nPustaka)

	DaftarkanBuku(&pustaka, nPustaka)

	fmt.Print("Masukkan rating buku yang dicari: ")
	fmt.Scan(&chooseRating)

	CetakTerfavorit(pustaka)

	UrutBuku(&pustaka)

	Cetak5Terbaru(pustaka)

	CariBuku(pustaka, chooseRating)
}