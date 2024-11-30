<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL XII-XIII </strong></h2>
<h2 align="center"><strong> PENGURUTAN DATA </strong></h2>

<br>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Wahyu Hidayat / 2311102178<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------

## I. Dasar Teori
### Definisi Pengurutan Data
Pengurutan data adalah proses mengatur elemen-elemen dalam suatu daftar berdasarkan urutan tertentu, seperti urutan menaik atau menurun. Teknik pengurutan ini digunakan secara luas di berbagai bidang untuk meningkatkan efisiensi pencarian, analisis data, dan pemrosesan lainnya.

#### 1. Ide Algoritma Selection Sort
Algoritma selection sort didasarkan pada konsep mencari elemen terkecil (atau terbesar) dari daftar yang belum diurutkan dan menempatkannya pada posisi yang sesuai. Proses ini diulangi untuk setiap elemen dalam daftar hingga seluruh daftar terurut. Prinsip dasarnya adalah membagi daftar menjadi dua bagian: bagian yang sudah diurutkan dan bagian yang belum diurutkan, lalu secara bertahap memperluas bagian yang sudah diurutkan dengan menambahkan elemen yang ditemukan【1】【2】.

#### 2. Algoritma Selection Sort
Langkah-langkah selection sort dapat dijelaskan sebagai berikut:

1. Mulai dari indeks pertama, temukan elemen terkecil di bagian daftar yang belum diurutkan.
2. Tukarkan elemen terkecil tersebut dengan elemen di posisi awal bagian yang belum diurutkan.
3. Geser batas bagian yang belum diurutkan satu elemen ke kanan.
4. Ulangi langkah 1-3 hingga semua elemen terurut.
Kompleksitas waktu selection sort dalam kasus terbaik, rata-rata, dan terburuk adalah O(n^2), di mana n adalah jumlah elemen dalam daftar. Algoritma ini tidak efisien untuk daftar berukuran besar tetapi sederhana untuk diimplementasikan【1】【3】.

#### 3. Ide Algoritma Insertion Sort
Algoritma insertion sort bekerja dengan cara membangun bagian daftar yang terurut secara bertahap. Elemen dari bagian yang belum diurutkan dimasukkan satu per satu ke dalam posisi yang tepat di bagian yang sudah diurutkan. Prinsip utamanya adalah memanfaatkan fakta bahwa bagian kecil dari daftar yang terurut dapat dipertahankan dengan biaya minimal untuk penyisipan elemen baru【4】【5】.

#### 4. Algoritma Insertion Sort
Langkah-langkah insertion sort adalah sebagai berikut:
1. Mulai dari elemen kedua dalam daftar, anggap elemen pertama sudah terurut.
2. Ambil elemen berikutnya dan bandingkan dengan elemen dalam bagian yang sudah diurutkan.
3. Geser elemen yang lebih besar ke kanan untuk membuat ruang bagi elemen yang sedang dimasukkan.
4. Tempatkan elemen tersebut pada posisi yang sesuai dalam bagian yang sudah diurutkan.
5. Ulangi langkah 2-4 untuk semua elemen dalam daftar.
Kompleksitas waktu insertion sort adalah O(n) dalam kasus terbaik (daftar sudah hampir terurut) dan O(n^2) dalam kasus terburuk (daftar terbalik). Algoritma ini efektif untuk daftar kecil dan bekerja dengan baik pada data yang hampir terurut【4】【6】.

## II. GUIDED
## 1. Pengurutan Nomor Rumah Kerabat dengan Selection Sort
#### Source Code
```go
package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan selection sort
func selectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat (n) : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("\nMasukkan jumlah nomor rumah kerabat untuk daerah %d : ", i+1)
		fmt.Scan(&m)

		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat : ", m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}
		selectionSort(arr, m)

		// Tampilkan nomor rumah terurut
		fmt.Printf("Nomor rumah terurut untuk daerah %d : ", i+1)
		for _, num := range arr {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
```
#### Screenshoot Source Code
![Screenshot 2024-11-30 064715](https://github.com/user-attachments/assets/c5abbe51-2c94-4887-aa75-ecd6749d67bf)

#### Screenshoot Output
![Screenshot 2024-11-30 064726](https://github.com/user-attachments/assets/6e9a2036-4e04-47c8-be81-6c3714c5b503)

#### Deskripsi Program
Program ini mengurutkan nomor rumah kerabat di beberapa daerah menggunakan algoritma selection sort. Pengguna akan diminta memasukkan jumlah daerah dan nomor rumah kerabat untuk masing-masing daerah. Setelah itu, program akan mengurutkan nomor-nomor rumah di setiap daerah dan menampilkan hasilnya.
#### Algoritma Program
1. Mulai dari indeks pertama hingga indeks terakhir-1:
   - Pilih elemen terkecil dari bagian array yang belum terurut.
   - Tukar elemen terkecil dengan elemen di indeks saat ini.
2. Ulangi langkah ini hingga seluruh array terurut.
Algoritma Program
1. Minta pengguna memasukkan jumlah daerah (n).
2. Untuk setiap daerah:
   - Minta jumlah nomor rumah kerabat (m).
   - Inputkan nomor rumah sebanyak m.
   - Urutkan nomor rumah dengan algoritma selection sort.
   - Tampilkan nomor rumah yang telah diurutkan.
3. Selesai.

#### Cara Kerja
1. Program meminta pengguna memasukkan jumlah daerah kerabat (n).
2. Program memproses setiap daerah:
   - Meminta jumlah nomor rumah (m) di daerah tersebut.
   - Memasukkan nomor-nomor rumah ke dalam array.
   - Mengurutkan array menggunakan algoritma selection sort:
     - Temukan nomor rumah terkecil pada bagian yang belum terurut.
     - Tukar posisi nomor rumah terkecil dengan elemen di indeks saat ini.
     - Ulangi hingga seluruh nomor rumah diurutkan.
3. Setelah array diurutkan, program menampilkan nomor rumah yang terurut untuk daerah tersebut.
4. Proses ini diulang untuk semua daerah.
5. Program selesai.


## 2. Pengurutan dan Analisis Pola Selisih Elemen Array
#### Source Code
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

	// Input data hingga bilangan negatif ditemukan
	fmt.Print("Masukkan data integer (Akhiri dengan bilangan negatif) : ")
	for {
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		arr = append(arr, num)
	}

	n := len(arr)

	// Urutkan array menggunakan Insertion Sort
	insertionSort(arr, n)

	// Periksa apakah selisih elemen tetap
	isConstant, difference := isConstantDifference(arr, n)

	// Tampilkan hasil pengurutan
	fmt.Print("Array setelah diurutkan : ")
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// Tampilkan status jarak
	fmt.Print("Data berjarak")
	if isConstant {
		fmt.Printf(" %d", difference)
	}
	fmt.Println()
}

```
#### Screenshoot Source Code
![Screenshot 2024-11-30 065423](https://github.com/user-attachments/assets/acf79d65-a412-4385-ae32-3373642d4013)

#### Screenshoot Output
![Screenshot 2024-11-30 065431](https://github.com/user-attachments/assets/17d3f723-3e76-45c2-ba9c-e6676afb6fc0)

#### Deskripsi Program
Program ini menerima sejumlah bilangan integer dari pengguna hingga angka negatif dimasukkan sebagai tanda akhir input. Data tersebut diurutkan menggunakan algoritma Insertion Sort, kemudian diperiksa apakah elemen-elemen dalam array memiliki selisih yang konstan. Jika pola selisih konstan ditemukan, program menampilkan nilai selisihnya.

#### Algoritma Program
1. Insertion Sort:
   - Mulai dari elemen kedua hingga elemen terakhir:
     - Simpan elemen saat ini (key).
     - Bandingkan dengan elemen sebelumnya dan geser elemen yang lebih besar ke kanan.
     - Tempatkan elemen saat ini (key) di posisi yang benar.
   - Ulangi hingga seluruh array terurut.
2. Pengecekan Selisih Konstan:
   - Ambil selisih antara elemen pertama dan kedua sebagai referensi.
   - Bandingkan selisih antar elemen berturut-turut dalam array.
   - Jika ada selisih yang tidak sama dengan referensi, maka selisih tidak konstan.
   - Jika semua selisih sama, catat nilai selisih konstan tersebut.
3. Algoritma Program:
   1. Input angka-angka positif dari pengguna hingga ditemukan angka negatif.
   2. Urutkan array menggunakan Insertion Sort.
   3. Periksa pola selisih elemen array menggunakan fungsi isConstantDifference.
   4. Tampilkan array terurut.
   5. Tampilkan hasil apakah elemen memiliki selisih konstan, beserta nilai selisihnya jika ada.

#### Cara Kerja
1. Program meminta pengguna memasukkan angka-angka positif satu per satu.
   - Masukkan angka negatif untuk mengakhiri input.
2. Data yang dimasukkan disimpan dalam array.
3. Array diurutkan menggunakan algoritma Insertion Sort:
   - Setiap elemen ditempatkan pada posisi yang benar dengan membandingkan dan menggeser elemen lainnya.
4. Program memeriksa apakah selisih antar elemen array konstan:
   - Selisih dihitung dari elemen pertama hingga terakhir.
   - Jika semua selisih sama, pola konstan teridentifikasi.
5. Program menampilkan:
   - Array setelah diurutkan.
   - Status apakah selisih elemen konstan, dan jika iya, tampilkan nilai selisih tersebut.


## III. UNGUIDED
## 1. Pengurutan Nomor Rumah Berdasarkan Ganjil dan Genap
#### Source Code
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("\nMasukkan jumlah nomor rumah kerabat untuk daerah %d: ", i+1)
		var m int
		fmt.Scan(&m)

		// Input data rumah
		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		// Pisahkan ganjil dan genap
		var ganjil, genap []int
		for _, num := range arr {
			if num%2 == 0 {
				genap = append(genap, num)
			} else {
				ganjil = append(ganjil, num)
			}
		}

		// Urutkan ganjil descending dan genap ascending
		sort.Sort(sort.Reverse(sort.IntSlice(ganjil))) // Descending untuk ganjil
		sort.Ints(genap)                              // Ascending untuk genap

		// Tampilkan hasil
		fmt.Printf("Nomor rumah terurut untuk daerah %d: ", i+1)
		for _, num := range ganjil {
			fmt.Printf("%d ", num)
		}
		for _, num := range genap {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}

```

#### Screenshoot Source Code
![Screenshot 2024-11-30 070545](https://github.com/user-attachments/assets/ebbbe739-397d-441d-823d-cd8c2fba9c62)

#### Screenshoot Output
![Screenshot 2024-11-30 070551](https://github.com/user-attachments/assets/9ea4248c-6851-403b-946d-d1269425e80b)

#### Deskripsi Program
Program ini mengurutkan nomor rumah kerabat di beberapa daerah berdasarkan sifat ganjil dan genap. Nomor ganjil diurutkan dari yang terbesar ke terkecil, sedangkan nomor genap diurutkan dari yang terkecil ke terbesar. Input berupa sejumlah baris yang masing-masing berisi nomor rumah kerabat. Program akan memisahkan, mengurutkan, dan menampilkan nomor ganjil terlebih dahulu diikuti dengan nomor genap.

#### Algoritma Program
1. Pemisahan Ganjil dan Genap:
   - Iterasi melalui array input.
   - Pisahkan elemen ganjil ke dalam array ganjil.
   - Pisahkan elemen genap ke dalam array genap.
2. Pengurutan:
   - Urutkan array ganjil secara descending (terbesar ke terkecil).
   - Urutkan array genap secara ascending (terkecil ke terbesar).
3. Gabung dan Tampilkan Hasil:
   - Cetak semua elemen array ganjil.
   - Cetak semua elemen array genap.
     
#### Cara Kerja
1. Program meminta pengguna memasukkan jumlah daerah (n).
2. Untuk setiap daerah:
   - Inputkan nomor rumah kerabat.
   - Pisahkan nomor ganjil dan genap ke dalam array berbeda.
   - Urutkan array ganjil secara descending dan array genap secara ascending.
   - Gabungkan hasil pengurutan ke dalam satu output sesuai aturan (ganjil diikuti genap).
3. Tampilkan output untuk setiap daerah.
4. Proses selesai.

## 2. Perhitungan Median dari Data yang Tersusun Secara Dinamis
#### Source Code
```go
package main

import (
	"fmt"
	"sort"
	"math"
)

// Fungsi untuk menghitung median
func calculateMedian(arr []int) int {
	n := len(arr)
	if n%2 == 1 { // Jika jumlah elemen ganjil
		return arr[n/2]
	}
	// Jika jumlah elemen genap
	return int(math.Floor(float64(arr[n/2-1]+arr[n/2]) / 2))
}

func main() {
	var arr []int
	for {
		var num int
		fmt.Scan(&num)

		// Jika input adalah -5313, keluar dari program
		if num == -5313 {
			break
		}

		// Jika input adalah 0, proses data
		if num == 0 {
			// Urutkan array
			sort.Ints(arr)

			// Hitung dan tampilkan median
			if len(arr) > 0 {
				median := calculateMedian(arr)
				fmt.Println(median)
			}

			// Kosongkan array untuk data berikutnya
			arr = nil
		} else {
			// Tambahkan bilangan ke array
			arr = append(arr, num)
		}
	}
}
```

#### Screenshoot Source Code
![Screenshot 2024-11-30 072400](https://github.com/user-attachments/assets/4fac8932-b29b-4405-9a26-0f1607347072)

#### Screenshoot Output
![Screenshot 2024-11-30 072411](https://github.com/user-attachments/assets/25dd7656-22a3-43bb-93cc-42087a17bf0b)

#### Deskripsi Program
Program ini membaca data berupa bilangan bulat positif secara dinamis hingga angka 0 ditemukan. Setelah itu, data yang sudah dibaca diurutkan, dan program menghitung median dari data tersebut. Median adalah nilai tengah dari data yang terurut. Jika jumlah data genap, median adalah rata-rata dari dua nilai tengah. Program terus berjalan hingga menemukan angka terminator -5313 sebagai tanda akhir input.

#### Algoritma Program
1. Input Data Dinamis:
   - Terima input bilangan satu per satu.
   - Masukkan bilangan ke dalam array hingga angka 0 ditemukan.
2. Urutkan Data:
   - Ketika angka 0 ditemukan, urutkan data menggunakan Insertion Sort.
3. Hitung Median:
   - Jika jumlah data ganjil, median adalah elemen tengah.
   - Jika jumlah data genap, median adalah rata-rata dua elemen tengah (dibulatkan ke bawah).
4. Tampilkan Median:
   - Cetak nilai median, dan kosongkan array untuk input berikutnya.
5. Akhiri Program:
   - Jika angka -5313 ditemukan, hentikan program.

#### Cara Kerja
1. Program membaca bilangan satu per satu.
2. Jika bilangan 0 ditemukan:
   - Data yang sudah terbaca diurutkan menggunakan Insertion Sort.
   - Median dari data tersebut dihitung dan ditampilkan.
3. Program mengulangi langkah-langkah di atas hingga angka -5313 ditemukan, yang menandakan akhir program.
4. Semua hasil median dicetak sesuai urutan pengolahan data.

## 3. Pengelolaan Data Buku Perpustakaan
#### Source Code
```go
package main

import (
	"fmt"
	"sort"
)

// Definisi struct Buku
type Buku struct {
	id, tahun, rating int
	judul, penulis, penerbit string
}

// Fungsi untuk memasukkan data buku
func DaftarkanBuku(n int) []Buku {
	var pustaka []Buku
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Println("Masukkan data buku ke-", i+1)
		fmt.Print("Judul: ")
		fmt.Scan(&buku.judul)
		fmt.Print("Penulis: ")
		fmt.Scan(&buku.penulis)
		fmt.Print("Penerbit: ")
		fmt.Scan(&buku.penerbit)
		fmt.Print("Tahun: ")
		fmt.Scan(&buku.tahun)
		fmt.Print("Rating: ")
		fmt.Scan(&buku.rating)
		buku.id = i + 1
		pustaka = append(pustaka, buku)
	}
	return pustaka
}

// Fungsi untuk mencetak buku terfavorit
func CetakTerfavorit(pustaka []Buku) {
	if len(pustaka) > 0 {
		fmt.Println("Buku terfavorit:")
		fmt.Printf("%s oleh %s (%s, %d) dengan rating %d\n",
			pustaka[0].judul, pustaka[0].penulis, pustaka[0].penerbit, pustaka[0].tahun, pustaka[0].rating)
	} else {
		fmt.Println("Tidak ada buku.")
	}
}

// Fungsi untuk mencetak 5 buku terbaru berdasarkan rating
func Cetak5Terbaru(pustaka []Buku) {
	fmt.Println("5 Buku dengan rating tertinggi:")
	for i := 0; i < len(pustaka) && i < 5; i++ {
		fmt.Printf("%d. %s oleh %s (%s, %d) dengan rating %d\n",
			i+1, pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit, pustaka[i].tahun, pustaka[i].rating)
	}
}

// Fungsi untuk mencari buku berdasarkan rating
func CariBuku(pustaka []Buku, rating int) {
	// Lakukan pencarian biner
	idx := sort.Search(len(pustaka), func(i int) bool { return pustaka[i].rating <= rating })
	if idx < len(pustaka) && pustaka[idx].rating == rating {
		fmt.Println("Buku ditemukan:")
		fmt.Printf("%s oleh %s (%s, %d) dengan rating %d\n",
			pustaka[idx].judul, pustaka[idx].penulis, pustaka[idx].penerbit, pustaka[idx].tahun, pustaka[idx].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu.")
	}
}

// Fungsi untuk mengurutkan data buku berdasarkan rating (menurun)
func Urutkan(pustaka []Buku) {
	sort.SliceStable(pustaka, func(i, j int) bool {
		return pustaka[i].rating > pustaka[j].rating
	})
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	// Masukkan data buku
	pustaka := DaftarkanBuku(n)

	// Urutkan data
	Urutkan(pustaka)

	// Cetak buku terfavorit
	CetakTerfavorit(pustaka)

	// Cetak 5 buku terbaru
	Cetak5Terbaru(pustaka)

	// Cari buku berdasarkan rating
	var rating int
	fmt.Print("Masukkan rating yang ingin dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, rating)
}

```

#### Screenshoot Source Code
![Screenshot 2024-11-30 073510](https://github.com/user-attachments/assets/afb063f3-ce3c-4ffe-bc43-fbc1ebd21c0d)

#### Screenshoot Output
![Screenshot 2024-11-30 073527](https://github.com/user-attachments/assets/560be561-90d3-4810-9518-b3fe6460abf7)

#### Deskripsi Program
Program ini mengelola data buku di sebuah perpustakaan menggunakan array dan struct. Program memungkinkan:
1. Menampilkan data buku dengan rating tertinggi.
2. Menampilkan 5 buku dengan rating tertinggi.
3. Mencari buku berdasarkan rating tertentu menggunakan pencarian biner.

#### Algoritma Program
1. Definisi Struktur:
   - Definisikan struct untuk menyimpan data buku.
   - Gunakan array untuk menyimpan kumpulan buku.
2. Masukkan Data Buku:
   - Baca jumlah buku.
   - Input data buku (judul, penulis, penerbit, tahun, rating).
3. Urutkan Data:
   - Gunakan metode Insertion Sort untuk mengurutkan data berdasarkan rating (menurun).
4. Cetak Data Buku Terfavorit:
   - Tampilkan buku dengan rating tertinggi (elemen pertama array).
5. Cetak 5 Buku Terbaru Berdasarkan Rating:
   - Tampilkan hingga 5 buku dengan rating tertinggi.
6. Cari Buku Berdasarkan Rating:
   - Lakukan pencarian biner untuk mencari buku dengan rating tertentu.
   - Jika ditemukan, tampilkan data buku.
   - Jika tidak ditemukan, tampilkan pesan.

#### Cara Kerja
1. Input data buku.
2. Data buku diurutkan berdasarkan rating menggunakan sort.SliceStable.
3. Cetak buku dengan rating tertinggi.
4. Cetak hingga 5 buku dengan rating tertinggi.
5. Cari buku berdasarkan rating tertentu menggunakan sort.Search (pencarian biner).

### Kesimpulan
Selection sort dan insertion sort adalah algoritma pengurutan sederhana. Selection sort bekerja dengan mencari elemen terkecil lalu memindahkannya ke posisi yang tepat, sedangkan insertion sort menyisipkan elemen ke tempat yang sesuai di bagian yang sudah terurut. Keduanya cocok untuk data kecil, tetapi kurang efisien untuk data besar karena memiliki kompleksitas waktu O(n^2) pada kasus rata-rata dan terburuk.

## Referensi 
[1] T. H. Cormen, C. E. Leiserson, R. L. Rivest, and C. Stein, Introduction to Algorithms, 3rd ed. Cambridge, MA: MIT Press, 2009.

[2] R. Sedgewick and K. Wayne, Algorithms, 4th ed. Boston, MA: Addison-Wesley, 2011.

[3] M. A. Weiss, Data Structures and Algorithm Analysis in C++, 4th ed. Upper Saddle River, NJ: Pearson, 2014.

[4] S. Dasgupta, C. H. Papadimitriou, and U. Vazirani, Algorithms, New York, NY: McGraw-Hill, 2008.

[5] A. Drozdek, Data Structures and Algorithms in C++, 4th ed. Boston, MA: Cengage Learning, 2013.

[6] J. Kleinberg and É. Tardos, Algorithm Design, Boston, MA: Pearson, 2006.
