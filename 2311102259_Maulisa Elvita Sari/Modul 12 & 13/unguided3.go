package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Buku struct {
	ID        int
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

func DaftarkanBuku(pustaka *[]Buku, n int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan data buku:")
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("Buku ke-%d:\n", i+1)

		fmt.Print("ID: ")
		scanner.Scan()
		buku.ID, _ = strconv.Atoi(scanner.Text())

		fmt.Print("Judul: ")
		scanner.Scan()
		buku.Judul = scanner.Text()

		fmt.Print("Penulis: ")
		scanner.Scan()
		buku.Penulis = scanner.Text()

		fmt.Print("Penerbit: ")
		scanner.Scan()
		buku.Penerbit = scanner.Text()

		fmt.Print("Eksemplar: ")
		scanner.Scan()
		buku.Eksemplar, _ = strconv.Atoi(scanner.Text())

		fmt.Print("Tahun: ")
		scanner.Scan()
		buku.Tahun, _ = strconv.Atoi(scanner.Text())

		fmt.Print("Rating: ")
		scanner.Scan()
		buku.Rating, _ = strconv.Atoi(scanner.Text())

		*pustaka = append(*pustaka, buku)
	}
}

func UrutkanBuku(pustaka *[]Buku) {
	data := *pustaka
	n := len(data)
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1

		for j >= 0 && data[j].Rating < key.Rating {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
	*pustaka = data
}

func CetakTerfavorit(pustaka []Buku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada buku dalam daftar.")
		return
	}

	fmt.Println("Buku dengan rating tertinggi:")
	fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Rating: %d\n",
		pustaka[0].ID, pustaka[0].Judul, pustaka[0].Penulis, pustaka[0].Penerbit, pustaka[0].Rating)
}

func CetakTopLima(pustaka []Buku) {
	fmt.Println("Lima buku dengan rating tertinggi:")
	top := 5
	if len(pustaka) < 5 {
		top = len(pustaka)
	}
	for i := 0; i < top; i++ {
		fmt.Printf("%d. Judul: %s, Rating: %d\n", i+1, pustaka[i].Judul, pustaka[i].Rating)
	}
}

func CariBuku(pustaka []Buku, targetRating int) {
	var hasil []Buku
	for _, buku := range pustaka {
		if buku.Rating == targetRating {
			hasil = append(hasil, buku)
		}
	}

	fmt.Printf("\nBuku dengan rating %d:\n", targetRating)
	if len(hasil) == 0 {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	} else {
		for _, buku := range hasil {
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Rating: %d\n", buku.Judul, buku.Penulis, buku.Penerbit, buku.Rating)
		}
	}
}

func main() {
	var pustaka []Buku
	var n, targetRating int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Jumlah buku harus lebih dari 0.")
		return
	}

	DaftarkanBuku(&pustaka, n)

	UrutkanBuku(&pustaka)

	CetakTerfavorit(pustaka)

	CetakTopLima(pustaka)

	fmt.Print("\nMasukkan rating buku yang ingin dicari: ")
	fmt.Scan(&targetRating)
	CariBuku(pustaka, targetRating)
}
