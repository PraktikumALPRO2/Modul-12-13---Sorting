package main
import "fmt"

func insertionsort(arr []int, n int){
	for i := 1; i < n; i++ {
		key := arr[i]
		j:= i-1

		//GESER ELEMENT YANG LEBIH BESAR DARI KEY KE KANAN
		for j >= 0 && arr[j] > key {
			arr[j+1]=arr[j]
			j--
		}
		arr[j+1]=key
	}
}
func isconstantdifference(arr  []int,n int)(bool, int){
	if n <2{
		return true , 0
	}

	difference := arr[1] - arr [0]
	for i := 1; i < n; i++ {
		if arr[i+1]-arr[i] !=difference{
			return false, 0
		}
	}
	return true, difference
}

func main(){
	var arr []int
	var num int

	//INPUT DATA HINGGA ILANGAN NEGATIF DITEMUKAN
	fmt.Println("MASUKKAN DATA INTEGER (AKHIRI DENGAN BILANGAN NEGATIF):")
	for{
		fmt.Scan(&num)
		if num < 0{
			break
		}
		arr = append(arr,num)
	}
	n:= len(arr)

	//URUTKAN ARRAY MENGGUNAKAN INSERTION SORT
	insertionsort(arr, n)
	//PERIKSA APAKAH SELISIH SETIAP ELEMENT TETAP
	isConstan,difference := isconstantdifference(arr, n)

	//TEMPILKAN HASIL PENGURUTAN
	fmt.Println("ARRAY SETELAH DIURUTKAN")
	for _,val := range arr{
		fmt.Printf("%d",val)
	}
	fmt.Println()
	//TAMPILKAN STATUS JARAK

	if isConstan{
		fmt.Printf("DATA BERJARAK %d\n",difference)
	}else {
		fmt.Println("DATA BERJARAK TIDAK TETAP")
	}
}