<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL XII & XIII</strong></h2>
<h2 align="center"><strong> PENGURUTAN DATA </strong></h2> 

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Caroline Carren A.R / 2311102174<br>
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
1. [Tujuan Praktikum](#tujuan-praktikum)
2. [Dasar Teori](#dasar-teori)
3. [Guided](#guided)
4. [Unguided](#unguided)
5. [Kesimpulan](#kesimpulan)
6. [Daftar Pustaka](#daftar-pustaka)

## Tujuan Praktikum
1. Mahasiswa dapat memahami konsep utama yang diajarkan dalam modul, termasuk implementasi algoritma dalam program komputer.
2. Mahasiswa dapat meningkatkan kemampuan mahasiswa dalam menganalisis dan memecahkan masalah menggunakan pendekatan algoritmik.
3. Mahasiswa dapat melatih keterampilan pemrograman dengan menerapkan teori ke dalam bentuk kode yang efisien dan efektif.
   
## Dasar Teori

**Selection Sort**

*Selection Sort* adalah algoritma pengurutan sederhana yang bekerja dengan membagi array menjadi dua bagian: bagian yang terurut dan yang belum terurut. Pada setiap iterasi, algoritma ini memilih elemen terkecil dari bagian yang belum terurut dan menukarnya dengan elemen pertama dari bagian tersebut.

**Karakteristik**

1. Kompleksitas waktu:
- Kasus terbaik, rata-rata, dan terburuk: O (n^2).

2. Kompleksitas ruang:
- Menggunakan ruang tambahan minimal: O(1).

3. Stabilitas:
- Tidak stabil, karena elemen-elemen dengan nilai yang sama dapat berpindah posisi relatif.

**Insertion Sort**

*Insertion Sort* adalah algoritma pengurutan sederhana yang bekerja dengan mengambil elemen dari bagian yang belum terurut, lalu menyisipkannya ke dalam posisi yang sesuai di bagian yang sudah terurut.

**Karakteristik**
1. Kompleksitas waktu:
- Kasus terbaik: O(n) (ketika array sudah terurut).
- Rata-rata dan terburuk: O(n^2).

2. Kompleksitas ruang:
- Menggunakan ruang tambahan minimal O(1).

3. Stabilitas:
- Stabil, karena tidak mengubah urutan relatif elemen-elemen dengan nilai yang sama.

## Guided 

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumah kerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort.  

**Masukan** dimulai dengan sebuah integer 𝒏 (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi 𝒏 baris berikutnya selalu dimulai dengan sebuah integer 𝒎 (0 < m < 
1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat. 

**Keluaran** terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing masing daerah. 

### Source Code :
```go
// Caroline Carren
// 2311102174
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
![image](https://github.com/user-attachments/assets/0a3d87c0-af5a-416a-b077-d4a7b45e1ccc)

### Full code Screenshot :
![carbon (5)](https://github.com/user-attachments/assets/685e1ac2-31af-479f-9cc4-a98bb1a75e79)

### Deskripsi Program : 
Program ini digunakan untuk mengurutkan nomor rumah kerabat di berbagai daerah menggunakan algoritma Selection Sort. Pengguna diminta untuk memasukkan jumlah daerah kerabat, lalu memasukkan jumlah rumah di setiap daerah beserta nomor-nomor rumah tersebut. Setelah data dimasukkan, program akan mengurutkan nomor-nomor rumah di setiap daerah secara menaik (ascending). Program memastikan bahwa input berupa jumlah daerah dan rumah valid (lebih besar dari nol). Hasil pengurutan nomor rumah akan ditampilkan untuk setiap daerah secara terpisah.

### Algoritma Program :
1. Input Jumlah Daerah
- Minta pengguna memasukkan jumlah daerah kerabat (n).
- Validasi input agar n > 0. Jika tidak valid, program akan berhenti.

2. Loop untuk Setiap Daerah
- Untuk setiap daerah, minta pengguna memasukkan jumlah rumah kerabat (m).
- Validasi input agar m > 0. Jika tidak valid, abaikan input untuk daerah tersebut dan lanjutkan ke daerah berikutnya.

3. Input Nomor Rumah
- Minta pengguna memasukkan nomor rumah sebanyak m untuk daerah tersebut.
- Simpan nomor rumah ke dalam slice.

4. Pengurutan dengan Selection Sort
- Gunakan algoritma Selection Sort untuk mengurutkan nomor rumah secara menaik.

5. Output Hasil Pengurutan
- Tampilkan nomor rumah yang sudah diurutkan untuk setiap daerah.

6. Selesai
- Setelah semua daerah selesai diproses, program akan menampilkan pesan bahwa pengurutan telah selesai.
  
### Cara Kerja Program :
1. Program dimulai dengan meminta pengguna memasukkan jumlah daerah kerabat (n). Jika input n kurang dari atau sama dengan 0, program akan memberikan pesan kesalahan dan berhenti.
2. Untuk setiap daerah, pengguna diminta memasukkan jumlah rumah kerabat (m). Jika m tidak valid (≤ 0), program akan melewati daerah tersebut.
3. Pengguna diminta memasukkan nomor rumah kerabat sebanyak m untuk daerah tersebut. Nomor-nomor rumah disimpan dalam sebuah slice.
4. Program mengurutkan nomor-nomor rumah dalam slice menggunakan algoritma Selection Sort. Algoritma ini mencari elemen terkecil dari bagian yang belum terurut, kemudian menukarnya dengan elemen pertama dari bagian tersebut.
5. Setelah pengurutan selesai, program menampilkan hasil nomor rumah yang telah diurutkan untuk setiap daerah.
6. Program mengulang langkah 2-5 hingga semua daerah selesai diproses. Terakhir, program menampilkan pesan bahwa pengurutan selesai.
   
### 2. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya. 

**Masukan** terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array. 

**Keluaran** terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap". 

### Source Code :
```go
// Caroline Carren
// 2311102174
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
![image](https://github.com/user-attachments/assets/f9cba8c9-32c0-4952-9730-9c0e3ba65fe7)

### Full code Screenshot :
![carbon (6)](https://github.com/user-attachments/assets/0e01c999-b879-4f74-9e21-c4b9d8922a11)

### Deskripsi Program : 
Program ini mengimplementasikan Insertion Sort untuk mengurutkan elemen-elemen dalam sebuah array yang dimasukkan oleh pengguna. Pengguna dapat memasukkan elemen-elemen array satu per satu, dan input diakhiri dengan bilangan negatif. Setelah data diurutkan, program memeriksa apakah selisih antar elemen pada array yang telah diurutkan bersifat tetap (konstan). Jika selisihnya tetap, program akan mencetak nilai selisih tersebut. Sebaliknya, jika tidak, program akan menyatakan bahwa data tidak memiliki selisih konstan.

### Algoritma Program :
1. Input Data
- Pengguna diminta memasukkan elemen-elemen array (tipe integer). Input dihentikan jika pengguna memasukkan bilangan negatif.
- Jika tidak ada data yang dimasukkan (array kosong), program akan menampilkan pesan bahwa tidak ada data yang dimasukkan, kemudian berhenti.

2. Urutkan Data dengan Insertion Sort
- Algoritma Insertion Sort digunakan untuk mengurutkan elemen array secara menaik.

3. Pemeriksaan Selisih Konstan
- Program menghitung selisih antara dua elemen berturutan pada array yang sudah diurutkan.
- Jika selisih antar elemen tetap untuk seluruh pasangan elemen, program menyatakan bahwa array memiliki selisih konstan dan menampilkan nilai selisih tersebut.
- Jika selisihnya tidak tetap, program menyatakan bahwa array memiliki selisih tidak tetap.

4. Tampilkan Hasil
- Cetak array yang telah diurutkan.
- Tampilkan status selisih array (tetap atau tidak tetap).

### Cara Kerja Program :
1. Input Data
- Pengguna memasukkan data elemen array satu per satu. Jika input bukan angka atau negatif, program memproses input sesuai dengan aturan (negatif untuk mengakhiri input).
- Array yang valid harus memiliki minimal satu elemen.

2. Pengurutan dengan Insertion Sort
- Mulai dari elemen kedua, setiap elemen disisipkan ke posisi yang tepat di bagian array yang sudah terurut sebelumnya dengan membandingkannya secara iteratif.

3. Pemeriksaan Selisih Konstan
- Jika array memiliki kurang dari dua elemen, program langsung menyatakan bahwa array memiliki selisih konstan dengan nilai 0.
- Untuk array dengan dua elemen atau lebih, program memeriksa selisih antar elemen berturutan. Jika selisih konstan untuk seluruh pasangan elemen, program mencetak nilai selisih tersebut.

4. Output Hasil
- Array yang sudah diurutkan ditampilkan.
- Status dan nilai selisih konstan (jika ada) ditampilkan.

## Unguided 

### 1. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.   

**Masukan** masih persis sama seperti sebelumnya (guided 1). 

**Keluaran** terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk 
nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah. 

### Source Code :
```go
// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Fungsi untuk mengurutkan array dengan selection sort
func selectionSort(arr []int, asc bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if (asc && arr[j] < arr[idx]) || (!asc && arr[j] > arr[idx]) {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

// Kode warna untuk terminal
const (
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	reset   = "\033[0m"
)

// Fungsi untuk mencetak garis batas panjang
func printSeparator() {
	fmt.Println("\n" + cyan + "=================================================" + reset)
}

func printDashedSeparator() {
	fmt.Println(yellow + "-------------------------------------------------" + reset)
}

func printDoubleDashedSeparator() {
	fmt.Println(magenta + "*************************************************" + reset)
}

// Fungsi untuk menampilkan nomor rumah berdasarkan daerah
func processDaerah(i, m int) ([]int, []int) {
	var arr = make([]int, m)
	// Input nomor rumah kerabat untuk setiap daerah
	fmt.Printf(cyan+"Masukkan %d nomor rumah kerabat untuk daerah ke-%d: "+reset, m, i+1)
	for j := 0; j < m; j++ {
		fmt.Scan(&arr[j])
	}

	// Memisahkan nomor ganjil dan genap
	var odd, even []int
	for _, num := range arr {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	return odd, even
}

// Fungsi untuk menampilkan hasil terurut
func printSortedResults(i int, odd, even []int) {
	// Mengurutkan nomor rumah
	selectionSort(odd, true)   // urutkan ganjil secara menaik
	selectionSort(even, false) // urutkan genap secara menurun

	// Output nomor rumah yang terurut
	printDoubleDashedSeparator()
	fmt.Printf(magenta+"\nNomor rumah terurut untuk daerah ke-%d:\n"+reset, i+1)

	// Tampilkan nomor ganjil
	fmt.Print(green)
	for _, num := range odd {
		fmt.Printf("%d ", num)
	}

	// Tampilkan nomor genap
	fmt.Print(blue)
	for _, num := range even {
		fmt.Printf("%d ", num)
	}
	fmt.Println(reset)
}

// Fungsi utama program
func main() {
	var n int
	// Input jumlah daerah kerabat
	printSeparator()
	fmt.Print(green + "Masukkan jumlah daerah kerabat (n): " + reset)
	fmt.Scan(&n)

	// Proses untuk setiap daerah
	for i := 0; i < n; i++ {
		var m int
		// Input jumlah rumah kerabat untuk setiap daerah
		printDashedSeparator()
		fmt.Printf(yellow+"Masukkan jumlah rumah kerabat di daerah ke-%d (m): "+reset, i+1)
		fmt.Scan(&m)

		// Ambil nomor rumah untuk daerah tersebut
		odd, even := processDaerah(i, m)

		// Menampilkan hasil terurut untuk daerah tersebut
		printSortedResults(i, odd, even)

		// Cetak garis pemisah untuk setiap daerah
		printDashedSeparator()
	}
}
```

### Output:
![image](https://github.com/user-attachments/assets/1a529f6e-a90a-4d4d-ba91-728966863d8a)

### Full code Screenshot :
![carbon (11)](https://github.com/user-attachments/assets/83e5c85b-4bde-45dc-bad6-6c52dbb66868)


### Deskripsi Program : 
Program ini mengelola dan mengurutkan nomor rumah di beberapa daerah berdasarkan kategori ganjil dan genap. Pengguna diminta memasukkan jumlah daerah dan jumlah rumah di setiap daerah, kemudian menginputkan nomor rumah untuk masing-masing daerah. Nomor rumah yang ganjil diurutkan secara membesar (ascending), sedangkan nomor rumah genap diurutkan secara mengecil (descending) menggunakan algoritma Selection Sort. Program memisahkan hasil pengurutan ke dalam dua kelompok (ganjil dan genap) dan menampilkan hasil pengurutan untuk setiap daerah. Program memastikan validitas input, seperti jumlah daerah dan rumah harus lebih besar dari nol.

### Algoritma Program :
1. Input Jumlah Daerah
- Minta pengguna memasukkan jumlah daerah (n).
- Validasi input agar n > 0. Jika tidak valid, program akan berhenti.

2. Input Data Setiap Daerah
- Untuk setiap daerah, minta pengguna memasukkan jumlah rumah (m).
- Jika m <= 0, program melewatkan daerah tersebut dan memberikan peringatan.

3. Input Nomor Rumah
- Pengguna memasukkan nomor rumah sebanyak m untuk daerah tersebut, disimpan dalam array.

4. Pisahkan Ganjil dan Genap
- Nomor rumah dibagi menjadi dua kelompok: ganjil dan genap.
- Nomor ganjil diurutkan menggunakan Selection Sort secara ascending.
- Nomor genap diurutkan menggunakan Selection Sort secara descending.

4. Tampilkan Hasil
- Cetak nomor rumah ganjil yang telah diurutkan.
- Cetak nomor rumah genap yang telah diurutkan.

### Cara Kerja Program :
1. Input Jumlah Daerah
- Pengguna memasukkan jumlah daerah (n). Jika n ≤ 0, program berhenti dengan pesan kesalahan.

2. Input Nomor Rumah
- Untuk setiap daerah:
  a. Masukkan jumlah rumah (m). Jika m ≤ 0, program melewati daerah tersebut.
  b. Masukkan nomor rumah sebanyak m.

3. Pisah dan Urutkan
- Nomor rumah dipisahkan menjadi dua kelompok: ganjil dan genap.
- Kelompok ganjil diurutkan membesar, dan kelompok genap diurutkan mengecil menggunakan algoritma Selection Sort.

4. Output Hasil Pengurutan
- Nomor rumah yang sudah terurut ditampilkan dalam format:
a. Ganjil (ascending)
b. Genap (descending)

5. Program Berakhir
- Setelah seluruh daerah selesai diproses, program menampilkan pesan bahwa pengurutan selesai.

### 2.Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya? "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0. 

**Masukan** berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 
data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak 
termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313. 

**Keluaran** adalah median yang diminta, satu data per baris.

### Source Code :
```go
// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi untuk melakukan selection sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Cari elemen terkecil di bagian yang belum terurut
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Tukar elemen terkecil yang ditemukan dengan elemen pertama
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// Fungsi untuk menghitung median dari array yang sudah terurut
func calculateMedian(arr []int) float64 {
	n := len(arr)
	if n%2 == 0 {
		// Jika jumlah elemen genap, median adalah rata-rata dari dua nilai tengah
		return float64(arr[n/2-1]+arr[n/2]) / 2.0
	}
	// Jika jumlah elemen ganjil, median adalah elemen tengah
	return float64(arr[n/2])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Membaca seluruh input sebagai satu baris
	fmt.Println("Masukkan angka yang dipisahkan dengan spasi (akhiri dengan -5313):")
	scanner.Scan()
	line := scanner.Text()

	// Memisahkan input ke dalam bagian-bagian
	parts := strings.Split(line, " ")
	var data []int

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Input tidak valid, harap masukkan angka bulat saja.")
			return
		}

		if num == -5313 {
			// Akhir dari input
			break
		} else if num == 0 {
			// Urutkan data dan hitung median
			selectionSort(data)
			median := calculateMedian(data)
			fmt.Printf("%.0f\n", median)
		} else {
			// Tambahkan angka ke array data
			data = append(data, num)
		}
	}
}
```

### Output:
![image](https://github.com/user-attachments/assets/e4f70350-ccf7-4b8d-acac-f950f26c7203)

### Full code Screenshot :
![carbon (8)](https://github.com/user-attachments/assets/b3761ffa-fd62-4361-b7e6-ba6d092b4697)

### Deskripsi Program : 
Program ini menerima input berupa angka-angka bulat yang dipisahkan dengan spasi. Input dapat dihentikan dengan angka khusus -5313. Program mengolah data dengan memanfaatkan algoritma Selection Sort untuk mengurutkan elemen-elemen yang dimasukkan. Saat pengguna memasukkan angka 0, program akan menghitung median dari data yang telah dimasukkan sejauh ini dan menampilkannya. Median dihitung berdasarkan data yang telah diurutkan: jika jumlah elemen ganjil, median adalah elemen tengah, sedangkan jika genap, median adalah rata-rata dari dua elemen tengah.

### Algoritma Program :
1. Input Data Dinamis
- Program membaca input sebagai satu baris yang dipisahkan dengan spasi.
- Input angka -5313 menandakan akhir dari masukan.

2. Pengelolaan Data
- Jika angka 0 dimasukkan:
  - Data yang telah dimasukkan sebelumnya diurutkan menggunakan Selection Sort.
  - Median dari data yang terurut dihitung dan ditampilkan.
- Angka selain 0 dan -5313 akan ditambahkan ke array data.

3. Pengurutan dengan Selection Sort
- Algoritma Selection Sort digunakan untuk mengurutkan data secara menaik.

4. Perhitungan Median
- Jika jumlah elemen ganjil, median adalah elemen tengah.
- Jika jumlah elemen genap, median adalah rata-rata dari dua elemen tengah.

5. Output Median
- Setiap kali angka 0 dimasukkan, median dari data yang terurut ditampilkan.

6. Akhir Program
- Program berhenti saat angka -5313 dimasukkan

### Cara Kerja Program :
1. Program meminta pengguna untuk memasukkan angka-angka bulat yang dipisahkan dengan spasi.
2. Setiap angka yang dimasukkan diproses sebagai berikut:
  - Jika angka adalah 0, program:
    - Mengurutkan array data menggunakan Selection Sort.
    - Menghitung dan menampilkan median dari data yang sudah diurutkan.
- Jika angka adalah -5313, program berhenti.
- Jika angka lainnya, program menambahkannya ke array data.
3. Program terus berjalan hingga pengguna memasukkan -5313.

### 3.Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:

![image](https://github.com/user-attachments/assets/64ac0f9e-108b-4131-a37a-f80ddc1212c0)

**Masukan** terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan 
banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya 
adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan 
bulat yang menyatakan rating buku yang akan dicari. 

**Keluaran** terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir.

![image](https://github.com/user-attachments/assets/e2ecab74-aab1-40e1-887f-1d8874995c00)

### Source Code :
```go
// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

const nMax = 10000

type Buku struct {
	id                       string
	judul                    string
	penulis                  string
	penerbit                 string
	eksemplar, tahun, rating int
}

type DaftarBuku [nMax]Buku

// Warna ANSI
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	white  = "\033[37m"
)

// Prosedur untuk mendaftarkan buku ke perpustakaan
func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Print(yellow + "Masukkan jumlah buku yang ingin didaftarkan: " + reset)
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Printf("\n"+blue+"Masukkan data buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n"+reset, i+1)
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit,
			&pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

// Prosedur untuk menampilkan buku dengan rating tertinggi
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println(red + "\nTidak ada data buku." + reset)
		return
	}

	terfavorit := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > terfavorit.rating {
			terfavorit = pustaka[i]
		}
	}

	fmt.Println("\n" + green + "=============================================" + reset)
	fmt.Println(green + "Buku Terfavorit:" + reset)
	fmt.Println(green + "=============================================" + reset)
	fmt.Printf("Judul    : "+cyan+"%s\n"+reset, terfavorit.judul)
	fmt.Printf("Penulis  : "+cyan+"%s\n"+reset, terfavorit.penulis)
	fmt.Printf("Penerbit : "+cyan+"%s\n"+reset, terfavorit.penerbit)
	fmt.Printf("Tahun    : "+cyan+"%d\n"+reset, terfavorit.tahun)
	fmt.Printf("Rating   : "+green+"%d\n"+reset, terfavorit.rating)
	fmt.Println(green + "=============================================" + reset)
}

// Prosedur untuk mengurutkan buku berdasarkan rating secara descending
func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1

		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

// Prosedur untuk mencetak 5 buku dengan rating tertinggi
func CetakTerbaru(pustaka DaftarBuku, n int) {
	fmt.Println("\n" + green + "=============================================" + reset)
	fmt.Println(green + "5 Buku Dengan Rating Tertinggi:" + reset)
	fmt.Println(green + "=============================================" + reset)
	count := 5
	if n < 5 {
		count = n
	}

	for i := 0; i < count; i++ {
		fmt.Printf("%d. "+cyan+"%s (Rating: %d)\n"+reset, i+1, pustaka[i].judul, pustaka[i].rating)
	}
	fmt.Println(green + "=============================================" + reset)
}

// Prosedur untuk mencari buku berdasarkan ratingnya
func CariBuku(pustaka DaftarBuku, n, r int) {
	found := false
	fmt.Printf("\n"+yellow+"Mencari Buku dengan Rating %d:\n"+reset, r)

	for i := 0; i < n; i++ {
		if pustaka[i].rating == r {
			found = true
			fmt.Printf("\n"+purple+"Judul    : %s\n"+reset, pustaka[i].judul)
			fmt.Printf("Penulis  : %s\n", pustaka[i].penulis)
			fmt.Printf("Penerbit : %s\n", pustaka[i].penerbit)
			fmt.Printf("Tahun    : %d\n", pustaka[i].tahun)
			fmt.Printf("Eksemplar: %d\n", pustaka[i].eksemplar)
			fmt.Printf("Rating   : "+green+"%d\n"+reset, pustaka[i].rating)
			fmt.Println("=============================================")
		}
	}

	if !found {
		fmt.Println(red + "\nTidak ada buku dengan rating tersebut." + reset)
	}
}

func main() {
	var pustaka DaftarBuku
	var n int
	var ratingCari int

	// Untuk menjalankan Prosedur Daftarkan Buku
	DaftarkanBuku(&pustaka, &n)

	// Untuk menjalankan Prosedur Cetak Buku Terfavorit
	CetakTerfavorit(pustaka, n)

	// Untuk menjalankan Prosedur Urutkan Buku berdasarkan Rating
	UrutBuku(&pustaka, n)

	// Untuk menjalankan Prosedur Cetak 5 buku dengan rating tertinggi
	CetakTerbaru(pustaka, n)

	// Untuk mencari buku berdasarkan rating
	fmt.Print("\n" + yellow + "Masukkan rating yang ingin dicari: " + reset)
	fmt.Scan(&ratingCari)
	CariBuku(pustaka, n, ratingCari)
}
```
### Output:
Tampilan Inputan Buku:

![image](https://github.com/user-attachments/assets/617cf974-d0f7-4914-b56d-379d6cfae8ed)

Tampilan Buku Terfavorite:

![image](https://github.com/user-attachments/assets/62bfe7c1-47aa-4252-bf7d-8d1e01a1a371)

Tampilan 5 Buku Dengan Rating Tertinggi:

![image](https://github.com/user-attachments/assets/b1aa770b-7270-4282-bda7-3e85b3f1fccd)

Tampilan Mencari Buku Dengan Rating:

![image](https://github.com/user-attachments/assets/c1423ba9-690c-43d5-a568-73122dd110ca)

### Full code Screenshot :
![carbon (10)](https://github.com/user-attachments/assets/043e0688-5e84-43f0-a6c1-db6838c3d03f)

### Deskripsi Program : 
Program ini berfungsi untuk mengelola data buku dalam sebuah perpustakaan, termasuk pendaftaran buku, pengurutan buku berdasarkan rating, dan pencarian buku berdasarkan rating tertentu. Pengguna dapat memasukkan informasi buku seperti ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating. Program akan menampilkan buku dengan rating tertinggi, mengurutkan buku-buku berdasarkan rating secara menurun, serta mencari buku yang memiliki rating tertentu. Selain itu, program juga dapat menampilkan lima buku dengan rating tertinggi jika tersedia. Semua data buku disimpan dalam array dan diolah sesuai dengan input yang diberikan oleh pengguna.

### Algoritma Program :
1. Input Jumlah Buku
- Minta pengguna memasukkan jumlah buku yang ingin didaftarkan (n).
- Validasi input agar n > 0. Jika tidak valid, program akan berhenti.

2. Input Data Buku
- Untuk setiap buku, minta pengguna memasukkan data buku: ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating.
- Data buku disimpan dalam array DaftarBuku.

3. Cetak Buku Terfavorit
- Program mencari buku dengan rating tertinggi di dalam array DaftarBuku dan mencetak informasi buku tersebut.

4. Urutkan Buku Berdasarkan Rating
- Buku diurutkan berdasarkan rating secara menurun menggunakan algoritma Insertion Sort.

5. Tampilkan 5 Buku dengan Rating Tertinggi
- Program mencetak lima buku dengan rating tertinggi, atau sebanyak n buku jika jumlah buku kurang dari 5.

6. Cari Buku Berdasarkan Rating
- Program meminta pengguna untuk memasukkan rating yang ingin dicari.
- Program mencari buku dengan rating tersebut dan menampilkan informasi buku yang ditemukan.

### Cara Kerja Program :
1. Input Jumlah Buku
- Pengguna memasukkan jumlah buku yang akan didaftarkan (n). Jika n ≤ 0, program berhenti dengan pesan kesalahan.

2. Input Data Buku
- Untuk setiap buku: a. Pengguna memasukkan ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating. b. Data buku disimpan dalam array DaftarBuku.

3. Cetak Buku Terfavorit
- Program memeriksa rating setiap buku dan menemukan buku dengan rating tertinggi. Buku ini kemudian dicetak.

4. Urutkan Buku Berdasarkan Rating
- Program menggunakan algoritma Insertion Sort untuk mengurutkan buku berdasarkan rating secara menurun (rating lebih tinggi lebih dulu).

5. Tampilkan 5 Buku dengan Rating Tertinggi
- Setelah pengurutan, program menampilkan lima buku dengan rating tertinggi. Jika jumlah buku kurang dari 5, program hanya menampilkan sebanyak jumlah buku yang ada.

6. Cari Buku Berdasarkan Rating
- Program meminta input rating dan mencari buku yang memiliki rating tersebut. Buku yang ditemukan ditampilkan, jika tidak ada, program memberi peringatan bahwa buku dengan rating tersebut tidak ditemukan.

7. Program Berakhir
- Setelah seluruh proses selesai, program selesai dan keluar.

## Kesimpulan 
Dari praktikum ini, dapat disimpulkan sebagai berikut:

1. Memahami penggunaan struktur data seperti array dan cara menyimpan data dengan atribut yang berbeda untuk pengelolaan informasi buku secara efisien.
2. Mengimplementasikan algoritma pengurutan, seperti Selection Sort, untuk mengurutkan data dengan benar dan memastikan data disajikan dalam urutan yang sesuai dengan kebutuhan.
3. Memahami penerapan pencarian data berdasarkan kriteria tertentu, seperti rating buku, yang memungkinkan program untuk menampilkan informasi yang relevan bagi pengguna.

## Daftar Pustaka

[1] A. A. A. Donovan and B. W. Kernighan, *The Go Programming Language*. Boston, MA: Addison-Wesley, 2015.

[2] W. Kennedy, B. Ketelsen, and E. St. Martin, *Go in Action*. Greenwich, CT: Manning Publications, 2016.
