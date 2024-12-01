package main

import "fmt"

func urutkanAngka(angka []int, urutkanTurun bool) {
	jumlahAngka := len(angka)
	for i := 0; i < jumlahAngka-1; i++ {
		indeksTerkecilAtauTerbesar := i
		for j := i + 1; j < jumlahAngka; j++ {
			if urutkanTurun {
				if angka[j] > angka[indeksTerkecilAtauTerbesar] {
					indeksTerkecilAtauTerbesar = j
				}
			} else {
				if angka[j] < angka[indeksTerkecilAtauTerbesar] {
					indeksTerkecilAtauTerbesar = j
				}
			}
		}
		angka[i], angka[indeksTerkecilAtauTerbesar] = angka[indeksTerkecilAtauTerbesar], angka[i]
	}
}

func pisahkanGanjilGenap(bilangan []int) ([]int, []int) {
	var bilanganGanjil, bilanganGenap []int
	for _, bil := range bilangan {
		if bil%2 == 0 {
			bilanganGenap = append(bilanganGenap, bil)
		} else {
			bilanganGanjil = append(bilanganGanjil, bil)
		}
	}
	return bilanganGanjil, bilanganGenap
}

func main() {
	var jumlahDaerah int
	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scan(&jumlahDaerah)

	for i := 1; i <= jumlahDaerah; i++ {
		var jumlahRumah int
		fmt.Printf("\nMasukkan jumlah rumah untuk daerah %d: ", i)
		fmt.Scan(&jumlahRumah)

		nomorRumah := make([]int, jumlahRumah)
		fmt.Printf("Masukkan %d nomor rumah: ", jumlahRumah)
		for j := 0; j < jumlahRumah; j++ {
			fmt.Scan(&nomorRumah[j])
		}

		rumahGanjil, rumahGenap := pisahkanGanjilGenap(nomorRumah)

		urutkanAngka(rumahGanjil, true)
		urutkanAngka(rumahGenap, false)

		fmt.Printf("\nHasil untuk daerah %d:\n", i)
		for _, num := range rumahGanjil {
			fmt.Printf("%d ", num)
		}
		for _, num := range rumahGenap {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
