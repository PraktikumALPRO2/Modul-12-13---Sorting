package main

import "fmt"

// Fungsi untuk mengurutkan array menggunakan selection sort dengan parameter urutan naik atau turun
func urutkanDenganSelection(arr []int, urutNaik bool) {
    panjang := len(arr)
    for i := 0; i < panjang-1; i++ {
        indeksTerpilih := i
        for j := i + 1; j < panjang; j++ {
            if (urutNaik && arr[j] < arr[indeksTerpilih]) || (!urutNaik && arr[j] > arr[indeksTerpilih]) {
                indeksTerpilih = j
            }
        }
        arr[i], arr[indeksTerpilih] = arr[indeksTerpilih], arr[i]
    }
}

func main() {
    var totalDaerah int
    // Meminta input jumlah daerah kerabat
    fmt.Print("Masukkan jumlah daerah kerabat: ")
    fmt.Scan(&totalDaerah)

    // Proses input dan pengurutan untuk setiap daerah
    for i := 0; i < totalDaerah; i++ {
        var totalRumah int
        // Input jumlah rumah untuk daerah tertentu
        fmt.Printf("\nMasukkan jumlah rumah di daerah %d: ", i+1)
        fmt.Scan(&totalRumah)

        nomorRumah := make([]int, totalRumah)
        // Input nomor rumah untuk daerah tersebut
        fmt.Printf("Masukkan %d nomor rumah: ", totalRumah)
        for j := 0; j < totalRumah; j++ {
            fmt.Scan(&nomorRumah[j])
        }

        // Pisahkan nomor rumah ganjil dan genap
        var rumahGanjil, rumahGenap []int
        for _, nomor := range nomorRumah {
            if nomor%2 == 0 {
                rumahGenap = append(rumahGenap, nomor)
            } else {
                rumahGanjil = append(rumahGanjil, nomor)
            }
        }

        // Urutkan rumah ganjil secara menaik
        urutkanDenganSelection(rumahGanjil, true)

        // Urutkan rumah genap secara menurun
        urutkanDenganSelection(rumahGenap, false)

        // Tampilkan hasil nomor rumah yang terurut
        fmt.Printf("\nNomor rumah terurut untuk daerah %d:\n", i+1)

        // Tampilkan nomor rumah ganjil yang sudah terurut
        for _, nomor := range rumahGanjil {
            fmt.Printf("%d ", nomor)
        }

        // Tampilkan nomor rumah genap yang sudah terurut
        for _, nomor := range rumahGenap {
            fmt.Printf("%d ", nomor)
        }

        fmt.Println()
    }
}
