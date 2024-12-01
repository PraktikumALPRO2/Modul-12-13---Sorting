<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL XII & XIII </strong></h2>
<h2 align="center"><strong> PENGURUTAN DATA </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Lutfiana Isnaeni L /2311102165<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh S.Kom., M.Kom. 
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
2. [guided](#guided)
3. [Unguided](#unguided)

## Dasar Teori
**Pengurutan Data (Sorting)**

Pengurutan data atau sorting adalah salah satu konsep dasar dalam ilmu komputer yang bertujuan untuk mengatur elemen-elemen dalam suatu kumpulan data (seperti array atau list) ke dalam urutan tertentu. Pengurutan ini umumnya dilakukan dalam urutan naik (ascending) atau turun (descending). Proses pengurutan data sangat penting karena dapat meningkatkan efisiensi berbagai algoritma, seperti pencarian data, dan mempermudah analisis data.

**Tujuan Pengurutan Data**

  Tujuan utama pengurutan data adalah untuk menyusun elemen dalam suatu urutan yang sistematis, sehingga proses pencarian, pencocokan, dan analisis menjadi lebih mudah dan cepat. Beberapa manfaat pengurutan data   antara lain:

1. Pencarian Efisien: Pengurutan data memungkinkan penggunaan algoritma pencarian yang lebih efisien, seperti binary search, yang hanya dapat digunakan pada data yang sudah terurut.
   
2. Analisis Data: Data yang terurut memudahkan dalam menganalisis pola atau distribusi data.
   
3. Pengelompokan: Dengan data yang terurut, proses pengelompokan atau klasifikasi data menjadi lebih sederhana.
   
4. Penyajian Data yang Jelas: Data yang terurut memberikan tampilan yang lebih rapi dan mudah dipahami.

**Jenis-jenis Pengurutan Data**

  Ada berbagai jenis algoritma pengurutan, masing-masing dengan karakteristik dan efisiensi yang berbeda-beda. Beberapa algoritma pengurutan yang populer adalah:

1. Bubble Sort:

    Bubble sort adalah algoritma pengurutan yang paling sederhana. Algoritma ini bekerja dengan membandingkan setiap elemen dengan elemen berikutnya dan menukarnya jika urutannya salah. Proses ini diulang hingga     seluruh data terurut.

    Kelemahan: Memiliki waktu eksekusi yang relatif lambat dengan kompleksitas waktu O(n²).
  
2. Selection Sort:

    Selection sort bekerja dengan cara mencari elemen terkecil (atau terbesar, tergantung urutan) dalam daftar yang belum terurut dan menukarnya dengan elemen pertama yang belum terurut. Proses ini diulang untuk     setiap elemen dalam daftar.

    Kelemahan: Kompleksitas waktu O(n²), meskipun lebih efisien daripada bubble sort dalam beberapa kasus.

3. Insertion Sort:

    Insertion sort mengurutkan elemen dengan cara menyisipkan elemen berikutnya ke posisi yang benar di dalam sublist yang sudah terurut. Proses ini berulang hingga seluruh data terurut.

    Kelemahan: Memiliki kompleksitas waktu O(n²) pada kasus terburuk.
  
4. Merge Sort:

    Merge sort adalah algoritma pengurutan yang menggunakan prinsip divide and conquer (bagi dan taklukkan). Data dibagi menjadi dua bagian, diurutkan secara rekursif, dan kemudian digabungkan kembali dengan         cara yang terurut.

    Kelebihan: Memiliki kompleksitas waktu O(n log n) pada semua kasus, sehingga lebih efisien dibandingkan algoritma pengurutan lainnya.

5. Quick Sort:

    Quick sort juga menggunakan prinsip divide and conquer. Data dibagi menjadi dua bagian berdasarkan elemen pivot, dan setiap bagian diurutkan secara rekursif.

    Kelebihan: Kompleksitas waktu rata-rata O(n log n), namun dapat memiliki kompleksitas O(n²) pada kasus terburuk.

6. Heap Sort:

    Heap sort menggunakan struktur data heap untuk mengurutkan elemen. Data diubah menjadi heap terlebih dahulu, kemudian elemen terbesar diambil dan dipindahkan ke posisi akhir, lalu heap dikurangi ukurannya.

    Kelebihan: Memiliki kompleksitas waktu O(n log n), dan tidak memerlukan ruang tambahan seperti merge sort.

**Algoritma Pengurutan Berdasarkan Kebutuhan**
    Pilihan algoritma pengurutan sangat tergantung pada ukuran data dan kebutuhan aplikasi. Misalnya:

  - Jika data relatif kecil dan implementasi sederhana diinginkan, bubble sort atau insertion sort bisa digunakan meskipun kurang efisien.
  
  - Untuk pengurutan data besar dengan efisiensi tinggi, merge sort, quick sort, atau heap sort lebih direkomendasikan.

**Kompleksitas Waktu**
    Kompleksitas waktu algoritma pengurutan adalah salah satu faktor penting dalam memilih algoritma. Kompleksitas waktu yang sering digunakan adalah:

  - O(n²): Terjadi pada bubble sort, selection sort, dan insertion sort di kasus terburuk.
  
  - O(n log n): Terjadi pada merge sort, quick sort, dan heap sort di kasus rata-rata.
  
  - O(n): Dapat tercapai oleh algoritma pengurutan seperti counting sort atau radix sort, jika kondisi tertentu terpenuhi (misalnya, data terbatasi dalam rentang tertentu).

## Guided
### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumah kerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort.
Masukan dimulai dengan sebuah integer 𝒏 (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi 𝒏 baris berikutnya selalu dimulai dengan sebuah integer 𝒎 (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat.

Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing masing daerah.
### Source Code: 
```go
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
#### Output Program:
![image](https://github.com/user-attachments/assets/0bfe5d55-f307-4935-ae25-358e95d86d5e)

### Deskripsi Program:
Program di atas digunakan untuk mengurutkan nomor rumah kerabat di beberapa daerah menggunakan algoritma Selection Sort. Pada fungsi utama `(main)`, program pertama-tama meminta pengguna untuk memasukkan jumlah daerah kerabat yang akan diolah. Setelah itu, untuk setiap daerah, program meminta input mengenai jumlah rumah kerabat dan nomor rumah yang ingin diurutkan. Input nomor rumah kemudian diolah dengan menggunakan fungsi `selectionSort`, yang mengurutkan nomor rumah dengan cara mencari elemen terkecil dalam setiap iterasi dan menukarnya dengan elemen di posisi yang sesuai. Program juga dilengkapi dengan validasi untuk memastikan bahwa jumlah daerah dan jumlah rumah lebih besar dari 0. Jika input tidak valid, program akan menampilkan pesan kesalahan dan meminta input ulang. Setelah proses pengurutan selesai, program menampilkan hasil pengurutan nomor rumah untuk setiap daerah yang dimasukkan oleh pengguna.

### 2. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.
Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array.

Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".
### Source Code: 
```go
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
#### Output Program:
![image](https://github.com/user-attachments/assets/327a10a4-5772-4556-b85c-37ca68cb6ce1)

### Deskripsi Program:
Program di atas digunakan untuk mengurutkan serangkaian angka yang dimasukkan oleh pengguna menggunakan algoritma Insertion Sort, serta memeriksa apakah selisih antara elemen-elemen dalam array yang telah diurutkan tetap konstan. Pengguna diminta untuk memasukkan angka satu per satu, dengan input diakhiri oleh angka negatif yang menandakan penghentian input. Setelah data dimasukkan, program akan mengurutkan angka-angka tersebut menggunakan metode Insertion Sort, yang bekerja dengan memindahkan setiap elemen ke posisi yang sesuai di bagian array yang sudah terurut. Setelah pengurutan, program memeriksa apakah selisih antar elemen tetap konstan. Jika ya, program akan menampilkan nilai selisih tersebut, dan jika tidak, program akan menginformasikan bahwa selisih antar elemen tidak tetap. Program kemudian menampilkan array yang sudah terurut dan memberikan status tentang jarak antar elemen, apakah tetap atau tidak. Misalnya, jika pengguna memasukkan angka 5, 3, 9, 7, 1, dan angka negatif sebagai tanda akhir input, program akan mengurutkan data menjadi [1, 3, 5, 7, 9] dan memeriksa selisih antar elemen, yang dalam hal ini adalah 2. Jika selisihnya tetap, program akan mencetak "Data berjarak 2". Jika tidak, program akan menunjukkan bahwa selisih antar elemen tidak tetap. 


## Unguided
### 1. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.
Masukan masih persis sama seperti sebelumnya (guided 1).

Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.


### Source Code: 
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int

	// Meminta jumlah daerah kerabat
	fmt.Println("Masukkan jumlah daerah kerabat (n):")
	fmt.Scan(&n)

	// Validasi input jumlah daerah kerabat harus lebih besar dari 0
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
		houses := make([]int, m)

		// Meminta input nomor rumah kerabat
		fmt.Printf("Masukkan %d nomor rumah kerabat untuk daerah ke-%d: ", m, i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}

		// Mengurutkan nomor rumah menggunakan sort.Ints
		sort.Ints(houses)

		// Menampilkan hasil pengurutan
		fmt.Printf("Nomor rumah terurut untuk daerah ke-%d: ", i+1)
		for _, house := range houses {
			fmt.Printf("%d ", house)
		}
		fmt.Println()
	}

	fmt.Println("Pengurutan selesai.")
}
```

#### Output Program:
![image](https://github.com/user-attachments/assets/e97f814a-bcc4-4dd7-84fc-57518c8ababe)

### Deskripsi Program:
Program di atas digunakan ntuk mengelola dan mengurutkan nomor rumah kerabat di berbagai daerah. program meminta input jumlah daerah kerabat (n). Jika jumlah daerah tersebut tidak valid (misalnya kurang dari atau sama dengan 0), program akan menampilkan pesan kesalahan dan berhenti. Selanjutnya, program melakukan iterasi untuk setiap daerah yang dimasukkan, dan pada setiap daerah, program meminta jumlah rumah kerabat yang ada (m) di daerah tersebut. Jika jumlah rumah kerabat juga tidak valid, program akan menampilkan pesan kesalahan untuk daerah tersebut dan melanjutkan ke daerah berikutnya.

Untuk setiap daerah, program akan meminta pengguna untuk memasukkan nomor rumah kerabat yang ada di daerah tersebut, kemudian menyimpannya dalam sebuah slice. Setelah itu, program akan mengurutkan nomor rumah tersebut menggunakan fungsi `sort.Ints()` yang disediakan oleh paket `sort` dari Go. Hasil pengurutan ini akan ditampilkan untuk setiap daerah secara terpisah. Setelah semua daerah diproses dan nomor rumah terurut, program akan menampilkan pesan "Pengurutan selesai".

### 2. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya? "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.
Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313.

Keluaran adalah median yang diminta, satu data per baris.
### Source Code: 
``` go
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
```

#### Output Program:
![image](https://github.com/user-attachments/assets/90c2dd1b-fd2c-421c-84c8-e0dc9dd39d95)

### Deskripsi Program:
Program di atas digunakan untuk menghitung median dari sekumpulan angka yang diberikan. Median dihitung setiap kali angka 0 ditemukan, sedangkan angka -5313 menandai akhir input. Angka-angka selain 0 dan -5313 disimpan dalam array. Ketika angka 0 muncul, array diurutkan, dan median dihitung: untuk jumlah data ganjil, median adalah nilai tengah; untuk data genap, median adalah rata-rata dua nilai tengah yang dibulatkan ke bawah.

Setiap angka diproses satu per satu. Angka 0 memicu penghitungan median, sedangkan angka lainnya hanya ditambahkan ke array. Misalnya, jika input adalah `7 23 11 0`, data `[7, 23, 11]` diurutkan menjadi `[7, 11, 23]`, sehingga median yang dihasilkan adalah 11. Program terus mencetak median setiap kali 0 muncul hingga input berakhir dengan -5313.

### 3. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:
![image](https://github.com/user-attachments/assets/702a0296-11c1-460c-89f7-f535b9e0bd72)
Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.

Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir.
![image](https://github.com/user-attachments/assets/653fda98-c588-42ec-8d35-23c922f684bc)

### Source Code: 
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
	"sort"
)

// Definisi Struct Buku
type Buku struct {
	id        int
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

// Daftar Buku
type DaftarBuku []Buku

// Fungsi untuk mencatat data buku
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("Masukkan data buku ke-%d:\n", i+1)
		fmt.Print("ID: ")
		fmt.Scan(&buku.id)
		fmt.Print("Judul: ")
		fmt.Scan(&buku.judul)
		fmt.Print("Penulis: ")
		fmt.Scan(&buku.penulis)
		fmt.Print("Penerbit: ")
		fmt.Scan(&buku.penerbit)
		fmt.Print("Eksemplar: ")
		fmt.Scan(&buku.eksemplar)
		fmt.Print("Tahun: ")
		fmt.Scan(&buku.tahun)
		fmt.Print("Rating: ")
		fmt.Scan(&buku.rating)
		*pustaka = append(*pustaka, buku)
	}
}

// Fungsi untuk menampilkan buku dengan rating tertinggi
func CetakTerfavorit(pustaka DaftarBuku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada buku.")
		return
	}

	// Mencari buku dengan rating tertinggi
	terfavorit := pustaka[0]
	for _, buku := range pustaka {
		if buku.rating > terfavorit.rating {
			terfavorit = buku
		}
	}

	// Menampilkan data buku
	fmt.Println("Buku dengan rating tertinggi:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
		terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun, terfavorit.rating)
}

// Fungsi untuk mengurutkan buku berdasarkan rating (Descending)
func UrutkanBuku(pustaka *DaftarBuku) {
	sort.Slice(*pustaka, func(i, j int) bool {
		return (*pustaka)[i].rating > (*pustaka)[j].rating
	})
}

// Fungsi untuk menampilkan 5 buku dengan rating tertinggi
func Cetak5Terbaru(pustaka DaftarBuku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada buku.")
		return
	}

	fmt.Println("5 Buku dengan rating tertinggi:")
	for i := 0; i < len(pustaka) && i < 5; i++ {
		buku := pustaka[i]
		fmt.Printf("%d. Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
			i+1, buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.rating)
	}
}

// Fungsi untuk mencari buku berdasarkan rating
func CariBuku(pustaka DaftarBuku, rating int) {
	left, right := 0, len(pustaka)-1
	for left <= right {
		mid := (left + right) / 2
		if pustaka[mid].rating == rating {
			buku := pustaka[mid]
			fmt.Println("Buku yang ditemukan:")
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
				buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.rating)
			return
		} else if pustaka[mid].rating > rating {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	var n int

	// Input jumlah buku
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	// Input data buku
	DaftarkanBuku(&pustaka, n)

	// Urutkan buku berdasarkan rating
	UrutkanBuku(&pustaka)

	// Cetak buku dengan rating tertinggi
	CetakTerfavorit(pustaka)

	// Cetak 5 buku dengan rating tertinggi
	Cetak5Terbaru(pustaka)

	// Cari buku berdasarkan rating
	var rating int
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, rating)
}
```
#### Output Program:
![image](https://github.com/user-attachments/assets/a21b9398-981d-46a3-8fee-4c303548b16e)

### Deskripsi Program:
Program di atas digunakan untuk mengelola data buku di sebuah perpustakaan. Dengan program ini, pengguna dapat memasukkan data buku secara interaktif, mencakup atribut seperti ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating. Data buku yang dimasukkan kemudian dapat diurutkan berdasarkan rating dalam urutan menurun untuk memudahkan pencarian buku dengan kualitas terbaik. Program ini juga menyediakan fitur untuk menampilkan buku dengan rating tertinggi serta lima buku terbaik berdasarkan rating. Selain itu, pengguna dapat mencari buku tertentu berdasarkan rating menggunakan metode pencarian biner. Hasil dari setiap operasi ditampilkan dengan format yang informatif dan mudah dipahami. Program ini dirancang untuk mempermudah pengelolaan buku di perpustakaan, memberikan rekomendasi buku berkualitas tinggi, dan memastikan efisiensi dalam pencarian serta pengurutan data buku.


