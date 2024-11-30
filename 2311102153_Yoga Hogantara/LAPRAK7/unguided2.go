package main
import "fmt"

func selectionsort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func median(arr []int) int {
	if len(arr)%2 == 1 {
		return arr[len(arr)/2] 
	}
	return (arr[len(arr)/2-1] + arr[len(arr)/2]) / 2 
}

func main() {
	var n int
	var data []int

	fmt.Println("INPUT DATA END -5313:")

	for {
		fmt.Scan(&n)
		if n == -5313 {
			break
		}
		if n == 0 {
			if len(data) > 0 {
				selectionsort(data)
				fmt.Println(median(data)) 
				data = []int{}                
			}
		}else {
			data = append(data, n)
		}
	}
}
