package main

import (
	"fmt"
	"sort"
)

// Struct untuk menyimpan kelompok nomor rumah
type KelompokRumah struct {
	nomorGanjil []int
	nomorGenap  []int
}

// Fungsi untuk mengelompokkan dan mengurutkan nomor rumah
func kelolaNomorRumah(nomorRumah []int) KelompokRumah {
	var kelompok KelompokRumah

	// Memisahkan nomor rumah menjadi ganjil dan genap
	for _, nomor := range nomorRumah {
		if nomor%2 == 0 {
			kelompok.nomorGenap = append(kelompok.nomorGenap, nomor)
		} else {
			kelompok.nomorGanjil = append(kelompok.nomorGanjil, nomor)
		}
	}

	// Mengurutkan dengan package sort
	sort.Ints(kelompok.nomorGanjil)                             // Urut naik untuk ganjil
	sort.Sort(sort.Reverse(sort.IntSlice(kelompok.nomorGenap))) // Urut turun untuk genap

	return kelompok
}

func main() {
	var totalWilayah int

	// Meminta input jumlah wilayah
	fmt.Print("Masukkan jumlah wilayah: ")
	fmt.Scanln(&totalWilayah)

	// Pemrosesan setiap wilayah
	for wilayah := 1; wilayah <= totalWilayah; wilayah++ {
		var totalRumah int

		fmt.Printf("\nJumlah rumah di wilayah %d: ", wilayah)
		fmt.Scanln(&totalRumah)

		// Membuat slice untuk menyimpan nomor rumah
		daftarRumah := make([]int, totalRumah)
		fmt.Printf("Masukkan %d nomor rumah: ", totalRumah)

		// Input nomor rumah
		for i := 0; i < totalRumah; i++ {
			fmt.Scan(&daftarRumah[i])
		}

		// Memproses pengelompokan dan pengurutan
		kelompokRumah := kelolaNomorRumah(daftarRumah)

		// Menampilkan hasil pengurutan
		fmt.Printf("\nNomor rumah terurut wilayah %d:\n", wilayah)
		fmt.Print("Ganjil: ")
		for _, nomor := range kelompokRumah.nomorGanjil {
			fmt.Printf("%d ", nomor)
		}
		fmt.Print("\nGenap: ")
		for _, nomor := range kelompokRumah.nomorGenap {
			fmt.Printf("%d ", nomor)
		}
		fmt.Println()
	}
}
