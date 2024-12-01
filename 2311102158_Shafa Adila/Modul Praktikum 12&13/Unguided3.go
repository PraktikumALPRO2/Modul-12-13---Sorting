package main

import "fmt"

// Definisi konstanta dan struct
const nMax = 7919

type Buku struct {
    ID_158       int
    Judul_158    string
    Penbulis_158  string
    Penerbit_158 string
    Eksemplar_158 int
    Tahun_158    int
    Rating_158   int
}

type DaftarBuku [nMax]Buku

// Fungsi untuk memasukkan data buku
func DaftarKanBuku(pustaka *DaftarBuku, n *int) {
    var jumlahBuku int
    fmt.Print("Masukkan jumlah buku: ")
    fmt.Scan(&jumlahBuku)

    for i := 0; i < jumlahBuku; i++ {
        fmt.Printf("Masukkan data buku ke-%d:\n", i+1)
        fmt.Print("ID: ")
        fmt.Scan(&pustaka[i].ID_158)
        fmt.Print("Judul Buku: ")
        fmt.Scan(&pustaka[i].Judul_158)
        fmt.Print("Penbulis Buku: ")
        fmt.Scan(&pustaka[i].Penbulis_158)
        fmt.Print("Penerbit Buku: ")
        fmt.Scan(&pustaka[i].Penerbit_158)
        fmt.Print("Eksemplar: ")
        fmt.Scan(&pustaka[i].Eksemplar_158)
        fmt.Print("Tahun Terbit: ")
        fmt.Scan(&pustaka[i].Tahun_158)
        fmt.Print("Rating: ")
        fmt.Scan(&pustaka[i].Rating_158)
    }
    *n = jumlahBuku
}

// Fungsi untuk mencari buku favorit
func CetakFavorit(pustaka DaftarBuku, n int) {
    if n == 0 {
        fmt.Println("Tidak ada buku dalam pustaka.")
        return
    }

    tertinggi := pustaka[0]
    for i := 1; i < n; i++ {
        if pustaka[i].Rating_158 > tertinggi.Rating_158 {
            tertinggi = pustaka[i]
        }
    }

    fmt.Println("Buku dengan rating tertinggi:")
    fmt.Printf("Judul: %s\n", tertinggi.Judul_158)
    fmt.Printf("Penbulis: %s\n", tertinggi.Penbulis_158)
    fmt.Printf("Penerbit: %s\n", tertinggi.Penerbit_158)
    fmt.Printf("Tahun: %d\n", tertinggi.Tahun_158)
    fmt.Printf("Rating: %d\n", tertinggi.Rating_158)
}

// Fungsi untuk mengurutkan buku menggunakan algoritma selection sort
func UrutkanBuku(pustaka *DaftarBuku, n int) {
    for i := 0; i < n-1; i++ {
        maxIdx := i
        for j := i + 1; j < n; j++ {
            if pustaka[j].Rating_158 > pustaka[maxIdx].Rating_158 {
                maxIdx = j
            }
        }
        // Tukar buku dengan rating tertinggi ke posisi i
        pustaka[i], pustaka[maxIdx] = pustaka[maxIdx], pustaka[i]
    }
}

// Fungsi untuk mencetak buku yang sudah diurutkan
func CetakTerurut(pustaka DaftarBuku, n int) {
    if n == 0 {
        fmt.Println("Tidak ada buku dalam pustaka.")
        return
    }

    fmt.Println("Daftar buku yang sudah diurutkan berdasarkan rating:")
    for i := 0; i < n; i++ {
        fmt.Printf("Judul: %s\n", pustaka[i].Judul_158)
        fmt.Printf("Penbulis: %s\n", pustaka[i].Penbulis_158)
        fmt.Printf("Penerbit: %s\n", pustaka[i].Penerbit_158)
        fmt.Printf("Tahun: %d\n", pustaka[i].Tahun_158)
        fmt.Printf("Rating: %d\n", pustaka[i].Rating_158)
    }
}

// Fungsi untuk mencari buku berdasarkan rating tertentu
func CariBuku(pustaka DaftarBuku, n int, rating int) {
    found := false
    for i := 0; i < n; i++ {
        if pustaka[i].Rating_158 == rating {
            if !found {
                fmt.Println("Buku ditemukan:")
                found = true
            }
            fmt.Printf("Judul: %s\n", pustaka[i].Judul_158)
            fmt.Printf("Penbulis: %s\n", pustaka[i].Penbulis_158)
            fmt.Printf("Penerbit: %s\n", pustaka[i].Penerbit_158)
            fmt.Printf("Tahun: %d\n", pustaka[i].Tahun_158)
            fmt.Printf("Rating: %d\n", pustaka[i].Rating_158)
        }
    }
    if !found {
        fmt.Println("Tidak ada buku dengan rating tersebut.")
    }
}

// Fungsi utama
func main() {
    var pustaka DaftarBuku
    var n int

    DaftarKanBuku(&pustaka, &n)
    CetakFavorit(pustaka, n)
    UrutkanBuku(&pustaka, n)
    CetakTerurut(pustaka, n)

    var cariRating int
    fmt.Print("Masukkan rating yang ingin dicari: ")
    fmt.Scan(&cariRating)
    CariBuku(pustaka, n, cariRating)
}
