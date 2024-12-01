package main
import (
	"fmt"
)

//fungsi untuk mengurutkan array menggunakan selection sort
func SelectionSort(arr[]int, n int){
	for i := 0; i < n-1; i++{
		idxMin := i
		for j := i + 1; j < n; j++{
			//cari elemen terkecil
			if arr[j] < arr[idxMin]{
				idxMin = j
			}
		}
		//tuker elemen terkecil dengan elemen di posisi i
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func main(){
	var n int
	fmt.Printf("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	//proses tiap daerah
	for daerah := 1; daerah <= n; daerah++{
		var m int
		fmt.Printf("\nMasukkan jumlah nomor kerabat untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		//membaca nomor rumah untuk daerah ini
		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++{
			fmt.Scan(&arr[i])
		}

		//urutkan array dari terkecil ke terbesar
		SelectionSort(arr, m)

		// Tampilkan hasil
		fmt.Printf("Nomor rumah terurut untuk daerah %d: ", daerah)
		for i, num := range arr {
    	    if i > 0 {
            	 fmt.Printf(" ") 
   	        }
   		    fmt.Printf("%d", num)
        }
 		fmt.Println()

	}
}