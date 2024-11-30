package main
import "fmt"

func selectionSortmembesar(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
func selectionSortmengecil(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("INPUT JUMLAH BARIS: ")
	fmt.Scanln(&n)

	inputs := make([][]int, n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("INPUT JUMLAH ANGKA PADA BARIS KE-%d: ", i+1)
		fmt.Scanln(&m)

		row := make([]int, m)
		fmt.Printf("INPUT ANGKA BARIS KE-%d : ", i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&row[j])
		}

		inputs[i] = row
	}
	for _, input := range inputs {
		var ganjil []int
		var genap []int
		for _, num := range input {
			if num%2 == 0 {
				genap = append(genap, num)
			} else {
				ganjil = append(ganjil, num)
			}
		}
		selectionSortmembesar(ganjil)
		selectionSortmengecil(genap)
		output := append(ganjil, genap...)
		for _, num := range output {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
