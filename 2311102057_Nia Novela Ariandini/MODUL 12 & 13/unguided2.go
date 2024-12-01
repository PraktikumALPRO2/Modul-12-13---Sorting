package main

import "fmt"

// Fungsi untuk mengurutkan array menggunakan metode seleksi
func urutkanArray(data []int) {
  panjang := len(data)
  for i := 0; i < panjang-1; i++ {
    indeksTerkecil := i
    for j := i + 1; j < panjang; j++ {
      if data[j] < data[indeksTerkecil] {
        indeksTerkecil = j
      }
    }
    // Menukar elemen agar array terurut
    data[i], data[indeksTerkecil] = data[indeksTerkecil], data[i]
  }
}

// Fungsi untuk menghitung nilai median dari array
func hitungMedian(data []int) int {
  panjang := len(data)
  if panjang%2 == 0 {
    return (data[panjang/2-1] + data[panjang/2]) / 2
  }
  return data[panjang/2]
}

func main() {
  var angka int
  var dataSementara []int
  var kumpulanGrup [][]int

  fmt.Println("Masukkan data (akhiri dengan -5313)")
  for {
    fmt.Scan(&angka)
    if angka == -5313 { // Penghentian input
      break
    }
    if angka == 0 { // Menyimpan salinan data ke dalam grup baru
      kumpulanGrup = append(kumpulanGrup, append([]int{}, dataSementara...))
    } else {
      dataSementara = append(dataSementara, angka) // Tambahkan angka ke data sementara
    }
  }

  // Memproses setiap grup data untuk menghitung mediannya
  for indeks, grup := range kumpulanGrup {
    urutkanArray(grup) // Urutkan grup data
    nilaiMedian := hitungMedian(grup) // Hitung median grup
    fmt.Printf("Median %d : %d\n", indeks+1, nilaiMedian)
  }
}
