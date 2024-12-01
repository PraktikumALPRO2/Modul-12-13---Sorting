package main

import "fmt"

func urutkanMeningkat(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		temp := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[temp] {
				temp = j
			}
		}
		arr[i], arr[temp] = arr[temp], arr[i]
	}
}

func urutkanMenurun(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		temp := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[temp] {
				temp = j
			}
		}
		arr[i], arr[temp] = arr[temp], arr[i]
	}
}

func main() {
	var jumlahDaerah int
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&jumlahDaerah)

	for daerah := 1; daerah <= jumlahDaerah; daerah++ {
		var jumlahRumah int
		var rumah []int
		
		fmt.Printf("Masukkan jumlah nomor rumah kerabat untuk daerah %d: ", daerah)
		fmt.Scan(&jumlahRumah)

		rumah = make([]int, jumlahRumah)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", jumlahRumah)
		for i := 0; i < jumlahRumah; i++ {
			fmt.Scan(&rumah[i])
		}

		rumahGanjil := []int{}
		rumahGenap := []int{}
		for _, nomorRumah := range rumah {
			if nomorRumah%2 != 0 {
				rumahGanjil = append(rumahGanjil, nomorRumah)
			} else {
				rumahGenap = append(rumahGenap, nomorRumah)
			}
		}

		urutkanMeningkat(rumahGanjil, len(rumahGanjil))

		urutkanMenurun(rumahGenap, len(rumahGenap))

		rumah = append(rumahGanjil, rumahGenap...)

		fmt.Printf("Nomor rumah terurut untuk daerah %d: ", daerah)
		for _, nomor := range rumah {
			fmt.Printf("%d ", nomor)
		}
		fmt.Println()
	}
}
