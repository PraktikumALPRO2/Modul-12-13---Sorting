<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL 12 & 13</strong></h2>
<h2 align="center"><strong> PENGURUTAN DATA </strong></h2> 

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Haifa Zahra Azzimmi / 2311102163<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh,S.Kom.,M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------

## Daftar Isi
1. [Dasar Teori](#dasar-teori)
3. [Guided](#guided)
4. [Unguided](#unguided)


## Dasar Teori

**Selection Sort**


**Karakteristik**


**Karakteristik**


## Guided 

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumah kerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort.  

**Masukan** dimulai dengan sebuah integer 𝒏 (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi 𝒏 baris berikutnya selalu dimulai dengan sebuah integer 𝒎 (0 < m < 
1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat. 

**Keluaran** terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing masing daerah. 

### Source Code :
```go
// Haifa Zahra Azzimmi
// 2311102163
// S1 IF 11 5

package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan Selection Sort
func selectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		// Tukar elemen terkecil dengan elemen pada indeks i
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func main() {
	var n int

	// Meminta jumlah daerah kerabat
	fmt.Println("Masukkan jumlah daerah kerabat (n):")
	fmt.Scan(&n)

	// Validasi input jumlah daerah harus lebih besar dari 0
	if n <= 0 {
		fmt.Println("Jumlah daerah kerabat harus lebih besar dari 0.")
		return
	}

	// Loop untuk setiap daerah
	for i := 0; i < n; i++ {
		var m int

		// Meminta jumlah rumah kerabat di daerah ke-i
		fmt.Printf("Masukkan jumlah rumah kerabat di daerah ke-%d (m): ", i+1)
		fmt.Scan(&m)

		// Validasi input jumlah rumah kerabat harus lebih besar dari 0
		if m <= 0 {
			fmt.Printf("Jumlah rumah kerabat di daerah ke-%d harus lebih besar dari 0.\n", i+1)
			continue
		}

		// Membuat slice untuk menyimpan nomor rumah kerabat
		arr := make([]int, m)

		// Meminta input nomor rumah kerabat
		fmt.Printf("Masukkan %d nomor rumah kerabat untuk daerah ke-%d: ", m, i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		// Mengurutkan nomor rumah menggunakan Selection Sort
		selectionSort(arr, m)

		// Menampilkan hasil pengurutan
		fmt.Printf("Nomor rumah terurut untuk daerah ke-%d: ", i+1)
		for _, num := range arr {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}

	fmt.Println("Pengurutan selesai.")
}
```

### Output:

### Full code Screenshot :

### Deskripsi Program : 


   
### 2. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya. 

**Masukan** terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array. 

**Keluaran** terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap". 

### Source Code :
```go
// Haifa Zahra Azzimmi
// 2311102163
// S1 IF 11 5

package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan Insertion Sort
func insertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		// Geser elemen yang lebih besar dari key ke kanan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Fungsi untuk memeriksa apakah selisih elemen array tetap
func isConstantDifference(arr []int, n int) (bool, int) {
	if n < 2 {
		return true, 0
	}

	difference := arr[1] - arr[0]
	for i := 1; i < n-1; i++ {
		if arr[i+1]-arr[i] != difference {
			return false, 0
		}
	}
	return true, difference
}

func main() {
	var arr []int
	var num int

	fmt.Println("Masukkan data integer (akhiri dengan bilangan negatif):")

	// Input data hingga bilangan negatif ditemukan
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Input tidak valid. Masukkan bilangan integer.")
			continue
		}
		if num < 0 {
			break
		}
		arr = append(arr, num)
	}

	// Periksa apakah array kosong
	if len(arr) == 0 {
		fmt.Println("Tidak ada data yang dimasukkan.")
		return
	}

	n := len(arr)

	// Urutkan array menggunakan Insertion Sort
	insertionSort(arr, n)

	// Periksa apakah selisih elemen tetap
	isConstant, difference := isConstantDifference(arr, n)

	// Tampilkan hasil pengurutan
	fmt.Println("Array setelah diurutkan:")
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// Tampilkan status jarak
	if isConstant {
		fmt.Printf("Data berjarak %d\n", difference)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
```
### Output:

### Full code Screenshot :

### Deskripsi Program : 




## Unguided 

### 1. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.   

**Masukan** masih persis sama seperti sebelumnya (guided 1). 

**Keluaran** terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk 
nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah. 

### Source Code :
```go

```

### Output:

### Full code Screenshot :


### Deskripsi Program : 



### 2.Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya? "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0. 

**Masukan** berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 
data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak 
termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313. 

**Keluaran** adalah median yang diminta, satu data per baris.

### Source Code :
```go

```

### Output:

### Full code Screenshot :

### Deskripsi Program : 



### 3.Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:


**Masukan** terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan 
banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya 
adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan 
bulat yang menyatakan rating buku yang akan dicari. 

**Keluaran** terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir.


### Source Code :
```go

```
### Output:
Tampilan Inputan Buku:


Tampilan Buku Terfavorite:


Tampilan 5 Buku Dengan Rating Tertinggi:


Tampilan Mencari Buku Dengan Rating:


### Full code Screenshot :

### Deskripsi Program : 
