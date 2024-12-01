// 2311102195 - Yesika Widiyani

package main

import (
	"container/heap"
	"fmt"
)

// Struktur untuk heap max
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // MaxHeap menggunakan urutan menurun
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Struktur untuk heap min
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // MinHeap menggunakan urutan menaik
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Fungsi untuk menambah elemen ke dalam heap dan menjaga keseimbangan
func addNumber(num int, low *MaxHeap, high *MinHeap) {
	if low.Len() == 0 || num <= (*low)[0] {
		heap.Push(low, num)
	} else {
		heap.Push(high, num)
	}

	// Jaga keseimbangan heap
	if low.Len() > high.Len()+1 {
		heap.Push(high, heap.Pop(low))
	} else if high.Len() > low.Len() {
		heap.Push(low, heap.Pop(high))
	}
}

// Fungsi untuk menghitung median
func findMedian(low *MaxHeap, high *MinHeap) int {
	if low.Len() > high.Len() {
		return (*low)[0]
	}
	return ((*low)[0] + (*high)[0]) / 2
}

func main() {
	var input int
	low := &MaxHeap{}
	high := &MinHeap{}
	heap.Init(low)
	heap.Init(high)

	for {
		fmt.Scan(&input)
		if input == -5313 {
			break
		}

		if input == 0 {
			fmt.Println(findMedian(low, high))
		} else {
			addNumber(input, low, high)
		}
	}
}