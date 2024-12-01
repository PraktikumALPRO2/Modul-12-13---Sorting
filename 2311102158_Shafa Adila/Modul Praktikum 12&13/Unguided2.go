package main

import (
    "fmt"
)

// Fungsi untuk melakukan selection sort menaik
func SelectionSort(arr []int, n_158 int) {
    for i := 0; i < n_158-1; i++ {
        idxMin := i
        for j := i + 1; j < n_158; j++ {
            if arr[j] < arr[idxMin] {
                idxMin = j
            }
        }
        arr[i], arr[idxMin] = arr[idxMin], arr[i]
    }
}

// Fungsi untuk mencari median
func findMedian(data []int) int {
    n_158 := len(data)
    if n_158%2 == 1 {
        return data[n_158/2]
    }
    return (data[n_158/2-1] + data[n_158/2]) / 2
}

func main() {
    var data []int
    var input int

    for {
        fmt.Scan(&input)
        if input == -5313 {
            break
        }
        if input == 0 {
            if len(data) == 0 {
                continue
            }
            SelectionSort(data, len(data))
            median := findMedian(data)
            fmt.Println(median)
        } else {
            data = append(data, input)
        }
    }
}
