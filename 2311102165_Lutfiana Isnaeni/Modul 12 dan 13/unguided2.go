// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
	"sort"
)

func main() {
	var input int
	var data []int

	for {
		fmt.Scan(&input)

		// Jika input adalah -5313, program berhenti
		if input == -5313 {
			break
		}

		// Jika input adalah 0, hitung median
		if input == 0 {
			if len(data) == 0 {
				fmt.Println(0)
			} else {
				// Urutkan data
				sort.Ints(data)

				// Hitung median
				n := len(data)
				if n%2 == 1 {
					// Jika jumlah data ganjil
					fmt.Println(data[n/2])
				} else {
					// Jika jumlah data genap
					median := (data[n/2-1] + data[n/2]) / 2
					fmt.Println(median)
				}
			}
		} else {
			// Tambahkan input ke dalam data
			data = append(data, input)
		}
	}
}
