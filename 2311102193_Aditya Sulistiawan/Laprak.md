
<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL XII & XIII </strong></h2>
<h2 align="center"><strong> PENGURUTAN DATA </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Aditya Sulistiawan / 2311102193<br>
  S1 IF11-05
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

## I. Dasar Teori
### Definisi Selection Sort
Selection Sort adalah algoritma pengurutan sederhana yang bekerja dengan cara membagi array menjadi dua bagian: bagian yang sudah terurut dan bagian yang belum terurut. Pada setiap iterasi, algoritma mencari elemen terkecil dari bagian yang belum terurut, kemudian menukarnya dengan elemen pertama dari bagian tersebut. Proses ini diulangi hingga seluruh elemen array terurut.

Keunggulan: Mudah diimplementasikan.
Kelemahan: Tidak efisien untuk dataset besar.

### Definisi Insertion Sort
Insertion Sort bekerja dengan cara membangun array yang terurut satu per satu. Algoritma ini memindahkan elemen dari bagian yang belum terurut ke posisi yang sesuai dalam bagian yang sudah terurut, mirip dengan cara orang mengurutkan kartu di tangan.

Keunggulan: Efisien untuk dataset kecil atau dataset yang hampir terurut.
Kelemahan: Tidak cocok untuk dataset besar.

## II. GUIDED
**1. Hercules, preman terkenal seantero Ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort.**

_Masukan dimulai dengan sebuah integer (01000), banyaknya daerah kerabat Hercules tinggal. Isi a baris berikutnya selalu dimulai dengan sebuah integer m (0<m< 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkalan bilangan bulat positif, nomor rumah para kerabat._

_Keluaran terdiri dari a baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing- masing daerah._

#### Source Code
```go
/// Aditya Sulistiawan
/// 2311102193

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
![carbon](https://github.com/user-attachments/assets/cb62425a-6516-4676-90ab-a62459224c2a)

#### Screenshoot Output
![Screenshot 2024-11-30 123426](https://github.com/user-attachments/assets/162810f6-7dc9-416b-92fd-1464d032d02a)

#### Deskripsi Program
Program ini adalah program yang dirancang untuk menyusun nomor rumah di beberapa lokasi, berdasarkan masukan dari pengguna. Tiap lokasi memiliki sejumlah nomor rumah, program ini akan menyortir nomor rumah tersebut. Program ini menggunakan Selection Sort.

Algoritma Program :
1. Input Jumlah Daerah: Program meminta pengguna untuk memasukkan jumlah daerah kerabat yang akan diproses.
2. Loop untuk Setiap Daerah: Untuk setiap daerah, program meminta pengguna untuk memasukkan jumlah nomor rumah kerabat.
3. Input Nomor Rumah: Pengguna diminta untuk memasukkan nomor rumah kerabat sesuai dengan jumlah yang telah ditentukan.
4. Sorting: Nomor rumah yang dimasukkan akan diurutkan menggunakan algoritma selection sort.
5. Output: Setelah diurutkan, program menampilkan nomor rumah kerabat yang telah terurut untuk setiap daerah.

Cara Kerja dari program ini yaituProgram ini dimulai dengan meminta pengguna untuk memasukkan total daerah kerabat yang ingin diproses. Setelah itu, untuk tiap daerah, pengguna diminta untuk memasukkan jumlah nomor rumah yang akan diinput. Program kemudian mengumpulkan nomor-nomor tersebut dan mengurutkannya dengan menggunakan metode selection sort. Setelah proses pengurutan selesai, program menampilkan nomor rumah yang telah diurutkan untuk setiap daerah. Dengan cara ini, pengguna dapat dengan mudah melihat daftar nomor rumah kerabat yang teratur.

**2. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda Insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.**

_Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array._

_Keluaran terdiri dari dua haris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap"._

#### Source Code
```go
/// Aditya Sulistiawan
/// 2311102193

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
![carbon](https://github.com/user-attachments/assets/cfb308f4-8dbb-4eee-8031-8572963c5e3d)

#### Screenshoot Output
![Screenshot 2024-11-30 123859](https://github.com/user-attachments/assets/3b03eb67-0d14-4954-a634-392ff652d7b6)

#### Deskripsi Program
Program ini adalah program yang bertujuan untuk mengatur angka yang dimasukkan oleh pengguna dengan metode Selection Sort dan memeriksa apakah jarak antara angka-angka tersebut konsisten atau tidak. Apabila jaraknya konsisten, program akan menunjukkan nilai jaraknya. Jika tidak, program akan memberi informasi bahwa jaraknya tidak konsisten.

Algortima program :
1. Input deretan bilangan dari user, diakhir dengan bilangan negatif.
2. Tentukan panjang array n.
3. Ururtkan dengan menggunakan Insertion Sort.
4. Cek apakah selisih antar bilangan itu konstan.
5. Tampilkan array yang sudah diurutkan.
6. Tampilkan hasil pemeriksaan jarak antar elemen array.

Cara kerja dari program ini adalah pengguna memasukkan sekumpulan angka, yang diakhiri dengan angka negatif. Susunlah angka-angka tersebut menggunakan Insertion Sort. Cek apakah selisih antara elemen-elemen tersebut tetap atau tidak. Program akan menunjukkan array yang telah diurutkan dan hasil pemeriksaan selisih antar elemen-elemen tersebut.

## III. UNGUIDED
**1. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari normor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.**

_Masukan masih persis sama seperti sebelumnya._

_Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing masing daerah,_

#### Source Code
```go
/// Aditya  Sulistiawan
/// 2311102193

package main

import "fmt"

func selectionSortGanjil(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i] // Tukar elemen
	}
}

func selectionSortGenap(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i] // Tukar elemen
	}
}

func main() {
	var n, input int
	fmt.Print("Masukkan jumlah daerah (n): ")
	fmt.Scan(&n)

	if n <= 0 || n >= 1000 {
		fmt.Println("n harus lebih besar dari 0 dan kurang dari 1000.")
		return
	}

	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("Masukkan jumlah rumah kerabat untuk daerah ke-%d: ", i+1)
		fmt.Scan(&m)

		if m <= 0 || m >= 1000000 {
			fmt.Println("m harus lebih besar dari 0 dan kurang dari 1000000.")
			return
		}

		// Masukkan nomor rumah
		housesGanjil := make([]int, 0)
		housesGenap := make([]int, 0)
		fmt.Printf("Masukkan nomor rumah kerabat untuk daerah ke-%d: ", i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&input)
			if input%2 == 0 {
				housesGenap = append(housesGenap, input)
			} else {
				housesGanjil = append(housesGanjil, input)
			}
		}

		// Urutkan dengan selection sort
		selectionSortGanjil(housesGanjil)
		selectionSortGenap(housesGenap)

		// Cetak hasil
		fmt.Printf("Hasil urutan rumah untuk daerah ke-%d: ", i+1)
		for _, house := range housesGanjil {
			fmt.Printf("%d ", house)
		}
		for _, house := range housesGenap {
			fmt.Printf("%d ", house)
		}
		fmt.Println()
	}
}
```
#### Screenshoot Source Code
![carbon](https://github.com/user-attachments/assets/821dbd80-0d24-4350-8e46-38810afcd742)

#### Screenshoot Output
![Screenshot 2024-11-30 145540](https://github.com/user-attachments/assets/d192f7cd-3438-4527-88fb-a3628d60efd6)


#### Deskripsi Program
Program ini berfungsi untuk mengatur nomor rumah di berbagai tempat yang ditentukan oleh pengguna. Nomor rumah terbagi menjadi dua kelompok, yaitu ganjil dan genap. Angka rumah ganjil diatur secara naik, sedangkan angka rumah genap diatur secara turun. Program ini menggunakan Selection Sort. Program akan menampilkan nomor rumah untuk setiap lokasi.

Algortima Program :
1. Input jumlah daerah kerabat (n).
2. Lakukan pengulangan n untuk setiap daerah :
   - Input jumlah rumah kerabat di daerah ke - i (m).
   - Input m nomor rumah kerabat.
   - Memisahkan nomor ganjil dan genap.
   - Mengurutkan nomor ganjil dan genap menggunakan Selection Sort.
   - Tampilkan nomor rumah terurut.
3. Ulangi langkah ke - 2 untuk semua daerah.

Cara kerja program ini yaitu pengguna memasukkan jumlah daerah yang ingin diolah. Untuk setiap daerah, diminta masukan jumlah dan nomor rumah yang ada. Nomor rumah dibedakan menjadi odd dan even. Program ini menggunakan Selection Sort untuk mengurutkan secara ascending dan descending. Program menampilkan nomor rumah untuk setiap daerah.

**2. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan, Dari 13 problem yang diberikan, ada satu problem yang menarik, Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?**

_"Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap,  maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0._

_Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data O merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313._

_Keluaran adalah median yang diminta, satu data per baris._

#### Source Code
```go
/// Aditya Sulistiawan
/// 2311102193


package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
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

func median(arr []int) int {
	n := len(arr)
	if n%2 == 0 {
		return (arr[(n/2)-1] + arr[(n/2)]) / 2

	} else {

		return arr[(n / 2)]
	}
}

func main() {
	var input int
	var kumpulanAngka = make([]int, 0)
	var angkaSliced = make([]int, 0)
	for input != -5313 {
		fmt.Scan(&input)
		if input != -5313 {
			kumpulanAngka = append(kumpulanAngka, input)
		}
	}
	for _, angka := range kumpulanAngka {
		if angka == 0 {
			insertionSort(angkaSliced)
			fmt.Println(median(angkaSliced))

		} else {
			angkaSliced = append(angkaSliced, angka)
		}
	}
}
```
#### Screenshoot Source Code
![carbon](https://github.com/user-attachments/assets/e69451f4-6978-4fb9-b316-0808951319c4)

#### Screenshoot Output
![Screenshot 2024-11-30 150430](https://github.com/user-attachments/assets/0ce66303-add3-47a6-8328-76b7df9a8134)


#### Deskripsi Program
Program ini bertujuan untuk menghitung dan memperlihatkan nilai median dari data yang dimasukkan oleh pengguna. Data akan dimasukkan secara terus-menerus hingga pengguna memasukkan angka 0, yang menandakan bahwa median dari data yang telah dihitung dan ditampilkan. Program ini menerapkan metode Selection Sort. Jika jumlah datanya ganjil, ambil nilai yang berada di tengah. Jika jumlah datanya genap, ambil rata-rata dari dua nilai yang ada di tengah.

Algoritma Program :
1. Lakukan input data :
   - Jika input = 0, urutkan data secara Selection Sort dan hitung median.
   - Jika input tidak sama dengan 0 atau -5313, tambahkan nilai itu ke dalam array data.
   - Jika inputnya -5313, hentikan untuk input.
2. Urutkan data yang dinput.
3. Hitung median dari data ini :
   - Jika jumlah data ganjil, maka ambil elemen yang ditengah.
   - Jika jumlah data genap, maka ambil rata rata dua elemen yang ditengah.
4. Tampilkan hasil median.

Cara kerja program ini adalah meminta masukan dari pengguna. Jika pengguna menginput angka yang bukan 0 atau -5313, program akan menambahkan informasi ke dalam array data. Program ini menggunakan Selection Sort. Setelah data diurutkan, program akan menghitung median. Jika jumlah data ganjil, ambil nilai tengahnya. Jika jumlah data genap, ambil rata-rata dari dua nilai tengah. Program akan menunjukkan hasil median yang telah dihitung.

**3. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:**
```go
const nMax : integer = 7919
type Buku = <
  id, judul, penulis, penerbit : string
  eksemplar, tahun, rating : integer

type DaftarBuku = array [1..nMax) of Buku
Pustaka : Daftar Buku
nPustaka : integer
```

_Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari._

_Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku. yang dicari sesual rating yang diberikan pada masukan baris terakhir._

#### Source Code
```go
/// Aditya sulistiawan
/// 2311102193

package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	fmt.Println("Format masukan (id, judul, penulis, penerbit, eksamplar, tahun, rating)")
	for i := 0; i < n; i++ {
		fmt.Print("Buku ke ", i+1, " : ")
		fmt.Scanln(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var indexFav = 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[indexFav].rating {
			indexFav = i
		}
	}

	fmt.Printf("Buku terfavorit sekarang: : %s %s %s %s %d %d %d\n",
		pustaka[indexFav].id, pustaka[indexFav].judul, pustaka[indexFav].penerbit, pustaka[indexFav].penulis,
		pustaka[indexFav].eksemplar, pustaka[indexFav].rating, pustaka[indexFav].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		var key Buku = pustaka[i]
		var j int = i - 1

		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	if n > 5 {
		n = 5
	}
	fmt.Print(n, " buku rating tertinggi :")
	for i := 0; i < n; i++ {
		fmt.Print(" ", pustaka[i].judul)
	}
	fmt.Println()
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	kr := 0
	kn := n - 1
	var med int
	var found bool = false

	for kr <= kn && !found {
		med = (kr + kn) / 2

		if pustaka[med].rating > r {
			kr = med + 1
		} else if pustaka[med].rating < r {
			kn = med - 1
		} else {
			found = true
		}
	}
	if found {
		fmt.Printf("Buku dengan rating %d: %s %s %s %s %d %d %d\n",
			r, pustaka[med].id, pustaka[med].judul, pustaka[med].penerbit, pustaka[med].penulis,
			pustaka[med].eksemplar, pustaka[med].rating, pustaka[med].tahun)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka, chooseRating int
	fmt.Print("Masukkan banyaknya buku : ")
	fmt.Scan(&nPustaka)
	DaftarkanBuku(&pustaka, nPustaka)
	fmt.Print("Masukkan rating buku yang dicari : ")
	fmt.Scan(&chooseRating)
	CetakTerfavorit(pustaka, nPustaka)
	UrutBuku(&pustaka, nPustaka)
	Cetak5Terbaru(pustaka, nPustaka)
	CariBuku(pustaka, nPustaka, chooseRating)
}
```
#### Screenshoot Source Code
![carbon](https://github.com/user-attachments/assets/24cb78f9-fad8-42b1-8485-6cd95bccf90b)


#### Screenshoot Output
![Screenshot 2024-11-30 151818](https://github.com/user-attachments/assets/f7c5d969-30ce-4ad2-940e-c99b5b8475f5)


#### Deskripsi Program
Program ini adalah sistem untuk mengatur buku di perpustakaan. Pengguna diminta untuk mendaftar buku, menunjukkan buku dengan peringkat terbaik, mengurutkan buku menurut peringkat, mencetak 5 buku dengan peringkat tertinggi, dan mencari buku berdasarkan peringkat tertentu. Informasi buku mencakup ID, judul, pengarang, penerbit, eksemplar, tahun, dan peringkat. Sistem ini menggunakan algoritma Insertion Sort.

Algoritma Program :
1. Minta input jumlah buku yang didaftarkan.
2. Masukkan data buku yang meliputi ID, judul, penulis, penerbit, eksemplar, tahun, dan rating.
3. Tampilkan buku dengan rating tertinggi.
4. Urutkan buku dengan descending.
5. Tampilkan 2 buku dengan rating tertinggi.
6. Minta rating buku yang ingin dicari.
7. Cari buku dengan rating dan tampilkan buku tersebut.

Cara kerja program ini ialah meminta pengguna untuk memasukkan total buku yang ingin didaftarkan. Pengguna mengisi data buku yang terdiri dari ID, judul, penulis, penerbit, eksemplar, tahun, dan rating. Program kemudian menunjukkan buku yang memiliki rating tertinggi. Buku-buku diurutkan dengan metode Insertion Sort dan menampilkan 2 buku dengan rating terbaik. Pengguna juga memasukkan rating buku yang ingin dicari. Program akan menunjukan informasi buku yang sesuai dengan rating yang dicari oleh pengguna.

## Referensi
[1] Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2009). Introduction to Algorithms. MIT Press.                                       

[2] GeeksforGeeks. "Selection Sort and Insertion Sort." https://www.geeksforgeeks.org.
