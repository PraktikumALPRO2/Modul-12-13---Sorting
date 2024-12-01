// Haifa Zahra Azzimmi
// 2311102163
// S1 IF 11 5

package main

import "fmt"

// Fungsi untuk mengurutkan elemen dalam slice menggunakan algoritma selection sort
func sortSlice(slice []int) {
    panjangSlice := len(slice)
    for i := 0; i < panjangSlice-1; i++ {
        // Cari indeks elemen terkecil pada bagian slice yang belum diurutkan
        indeksMinimum := i
        for j := i + 1; j < panjangSlice; j++ {
            if slice[j] < slice[indeksMinimum] {
                indeksMinimum = j
            }
        }
        // Tukar elemen terkecil dengan elemen pada posisi pertama yang belum diurutkan
        slice[i], slice[indeksMinimum] = slice[indeksMinimum], slice[i]
    }
}

// Fungsi untuk menghitung nilai median dari slice yang sudah terurut
func cariMedian(slice []int) int {
    jumlahElemen := len(slice)
    if jumlahElemen%2 == 1 {
        // Jika jumlah elemen ganjil, ambil elemen yang ada di posisi tengah
        return slice[jumlahElemen/2]
    }
    // Jika jumlah elemen genap, hitung rata-rata dua elemen di tengah (dibulatkan ke bawah)
    return (slice[(jumlahElemen/2)-1] + slice[jumlahElemen/2]) / 2
}

func main() {
    var daftarAngka []int
    fmt.Print("Masukkan sejumlah angka (akhiri dengan -5313): ")

    // Loop untuk membaca input angka
    for {
        var angka int
        _, err := fmt.Scan(&angka)

        // Periksa apakah input adalah -5313 untuk menghentikan program
        if angka == -5313 {
            break
        }

        // Jika terjadi kesalahan saat membaca input, hentikan program
        if err != nil {
            fmt.Println("Input tidak valid, coba lagi.")
            break
        }

        // Jika angka 0 dimasukkan, urutkan slice dan cetak median
        if angka == 0 {
            sortSlice(daftarAngka)
            fmt.Printf("Nilai median: %d\n", cariMedian(daftarAngka))
        } else {
            // Tambahkan angka ke slice
            daftarAngka = append(daftarAngka, angka)
        }
    }
}
