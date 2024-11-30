
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
  Andreas Besar Wibowo / 2311102198<br>
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
Selection Sort merupakan salah satu algoritma pengurutan yang pengunaannya sangat efisien dan sederhana, dimana melakukan pencarian index dengan isi nilai terkecil setelah itu ditukar dengan isi index yang lebih besar dari index dengan nilai terkecil itu, selanjutnya index terkecil kedua akan disimpan pada variable tmp (sementara) lalu ditukara dengan index kedua dan berulang terus sampai semua selesai terurut[1]. 

### Definisi Insertion Sort
Insertion Sort merupakan salah satu algoritma pengurutan yang dimana melakukan perbandingkan elemen array dengan elemen array dengan memiliki nilai tertinggi dalam array. Jika elemen array pembanding memiliki nilai lebih tinggi, maka elemen array tersebut akan tukar posisi dengan elemen array yang memiliki nilai lebih rendah dan elemen array yang memiliki nilai lebih rendah akan menjadi elemen pembanding sampai posisi yang tepat ditemukan[2].

## II. GUIDED
**1. Hercules, preman terkenal seantero Ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort.**

_Masukan dimulai dengan sebuah integer (01000), banyaknya daerah kerabat Hercules tinggal. Isi a baris berikutnya selalu dimulai dengan sebuah integer m (0<m< 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkalan bilangan bulat positif, nomor rumah para kerabat._

_Keluaran terdiri dari a baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing- masing daerah._

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
![Guided1 carbon](https://github.com/user-attachments/assets/13f25bee-c47b-4f17-8b94-3ce430e80dd4)
#### Screenshoot Output
![Guided1 go](https://github.com/user-attachments/assets/4e09b4c4-89e2-49e5-9f4b-751a251feae9)
#### Deskripsi Program
Program ini merupakan program untuk mengurutkan nomor rumah untuk beberapa daerah, yang diambil berdasarkan inputan user. Setiap daerah memiliki jumlah nomor rumah, program akan mengurutkan nomor rumah. Program ini menggunakan Selection Sort. 

Algoritma Program :
1. Input jumlah daerah kerabat (n).
2. Lakukan pengulangan sebanyak n untuk setiap daerah : 
   - Input jumlah kerabat di daerah ke - i (m).
   - Input m nomor rumah.
   - Gunakan Selection Sort untuk mengurutkan nomor rumahnya.
   - Tampilkan nomor rumah yang sudah terurut.
3. Ulangi langkah ke - 2 untuk semua daerah yang ada.

Cara Kerja dari program ini yaitu user menginputkan jumlah daerah kerabat, untuk setiap daerah minta inputan jumlah rumah (m) dan nomor rumah. Setelah itu, nomor rumah diurutkan dengan Selection Sort. Jika nomor sudah terurut, program menampilkan untuk setiap daerah.

**2. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda Insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.**

_Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array._

_Keluaran terdiri dari dua haris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap"._

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
![Guided2 carbon](https://github.com/user-attachments/assets/9776143b-9db5-4304-ba32-aab14dcd9b37)
#### Screenshoot Output
![Guided2(1) go](https://github.com/user-attachments/assets/e69dd582-ca20-407f-8975-5986ea3091a0)
![Guided2(2) go](https://github.com/user-attachments/assets/9b347c43-5ec9-48be-ab63-2d400e36769f)
#### Deskripsi Program
Program ini merupakan program untuk mengurutkan bilangan yang diinputkan oleh user dengan menggunakan Selection Sort dan nengecek apakah selisih antar bilangan itu tetap atau tidak. Jika selisihnya tetap, program akan menampilkan nilai selisihnya. Jika tidak, program akan memberi tahu bahwa selisihnya tidak tetap.

Algortima program :
1. Input deretan bilangan dari user, diakhir dengan bilangan negatif.
2. Tentukan panjang array n.
3. Ururtkan dengan menggunakan Insertion Sort.
4. Cek apakah selisih antar bilangan itu konstan.
5. Tampilkan array yang sudah diurutkan.
6. Tampilkan hasil pemeriksaan jarak antar elemen array.

Cara kerja dari program ini yaitu user memasukkan deretan bilangan, dan diakhiri dengan bilangan negatif. Urutkan bilangan-bilangan tersebut dengan Insertion Sort. Periksa apakah selisih antar elemennya konstan atau tidak. Program akan menampilkan array yang sudah terurut dan hasil pemeriksaan selisih antar elemnnya.

## III. UNGUIDED
**1. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari normor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.**

_Masukan masih persis sama seperti sebelumnya._

_Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing masing daerah,_

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / IF-11-05

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

func main() {
	var n int
	// Input jumlah daerah kerabat
	fmt.Print("Masukkan jumlah daerah kerabat (n) : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		// Input jumlah rumah kerabat untuk setiap daerah
		fmt.Printf("\nMasukkan jumlah rumah kerabat di daerah ke-%d (m): ", i+1)
		fmt.Scan(&m)

		arr := make([]int, m)
		// Input nomor rumah kerabat
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		// untuk memisahkan nomor ganjil dan genap
		var odd, even []int
		for _, num := range arr {
			if num%2 == 0 {
				even = append(even, num)
			} else {
				odd = append(odd, num)
			}
		}

		// untuk mengurutkan ganjil dengan menaik
		selectionSort(odd, true)

		// untuk mengurutkan genap dengan menurun
		selectionSort(even, false)

		// Output nomor rumah yang terurut
		fmt.Printf("\nNomor rumah terurut untuk daerah ke-%d:\n", i+1)

		// Tampilkan nomor ganjil
		for _, num := range odd {
			fmt.Printf("%d ", num)
		}

		// Tampilkan nomor genap
		for _, num := range even {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
```
#### Screenshoot Source Code
![Unguided1 carbon](https://github.com/user-attachments/assets/3943daa6-6e30-480c-8068-3a35e4ec283d)

#### Screenshoot Output
![Unguided1 go](https://github.com/user-attachments/assets/eee9c1fe-79fd-4b4a-8325-c038ffd1b4c2)

#### Deskripsi Program
Program ini merupakan program untuk mengurutkan nomor rumah di beberapa daerah yang diinputkan oleh user. Nomor rumah dibagi 2 yaitu ganjil dan genap. Nomor rumah ganjil diururtkan menaik, sedangkan nomor rumah genap diurutkan menurun. Program ini menggunakan Selection Sort. Program menampilkan nomor rumah untuk setiap daerah.

Algortima Program :
1. Input jumlah daerah kerabat (n).
2. Lakukan pengulangan n untuk setiap daerah :
   - Input jumlah rumah kerabat di daerah ke - i (m).
   - Input m nomor rumah kerabat.
   - Memisahkan nomor ganjil dan genap.
   - Mengurutkan nomor ganjil dan genap menggunakan Selection Sort.
   - Tampilkan nomor rumah terurut.
3. Ulangi langkah ke - 2 untuk semua daerah.

Cara kerja dari program ini yaitu user memasukkan jumlah daerah yang ingin diproses. Pada setiap daerah meminta input jumlah dan nomor rumah yang ada. Nomor rumah dipisahkan ganjil dan genap. Program ini menggunakan Selection Sort untuk mengurutkan secara menaik dan menurun. Program menampilkan nomor rumah untuk setiap daerah.

**2. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan, Dari 13 problem yang diberikan, ada satu problem yang menarik, Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?**

_"Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap,  maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0._

_Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data O merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313._

_Keluaran adalah median yang diminta, satu data per baris._

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / IF-11-05

package main

import "fmt"

// Fungsi untuk mengurutkan array dengan selection sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Cari elemen terkecil di subarray yang belum terurut
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Tukar elemen terkecil dengan elemen pertama subarray
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// Fungsi untuk menghitung nilai median dari array yang sudah diurutkan
func cariMedian(arr []int) int {
	n := len(arr)
	if n%2 == 1 {
		// Jika jumlah elemen bernilai ganjil, maka ambil elemen tengah
		return arr[n/2]
	}
	// Jika jumlah elemen bernilai genap, maka ambil rata-rata dari dua elemen tengah dan dibulatkan
	return (arr[(n/2)-1] + arr[n/2]) / 2
}

func main() {
	var data []int
	fmt.Print("Masukkan angka-angka: ")

	for {
		var num int
		_, err := fmt.Scan(&num)

		if num == -5313 {
			break
		}

		if err != nil {
			break
		}

		if num == 0 {
			// untuk mengurutkan data dan mencetak median
			selectionSort(data)
			fmt.Println(cariMedian(data))
		} else {
			// menambahkan bilangan ke array
			data = append(data, num)
		}
	}
}
```
#### Screenshoot Source Code
![Unguided2 carbon](https://github.com/user-attachments/assets/34429bfd-0e58-4d7d-ac63-8d5144e8c037)

#### Screenshoot Output
![Unguided2 go](https://github.com/user-attachments/assets/62024cb9-6e88-423c-a964-4050210fb368)

#### Deskripsi Program
Program ini merupakan program untuk menghitung dan menampilkan nilai median dari data yang diinputkan oleh user. Data dimasukkan secara berlanjut sampai pengguna menginputkan nilai 0, yang menunjukan bahwa median dari data yang sudah dihitung dan ditampilkan. Program ini menggunakan Selection Sort. Jika datanya ganjil, ambil nilai yang ditengah. Jika datanya genap, ambil rata rata dari dua nilai yang ditengah.

Algoritma Program :
1. Lakukan input data :
   - Jika input = 0, urutkan data secara Selection Sort dan hitung median.
   - Jika input tidak sama dengan 0 atau -5313, tambahkan nilai itu ke dalam array data.
   - Jika inputnya -5313, hentikan untuk input.
2. Urutkan data yang dimasukkan.
3. Hitung median dari data ini :
   - Jika jumlah data ganjil, maka ambil elemen yang ditengah.
   - Jika jumlah data genap, maka ambil rata rata dua elemen yang ditengah.
4. Tampilkan hasil median.

Cara kerja dari program ini yaitu meminta input dari user. Jika pengguna memasukkan angka selain 0 atau -5313, program akan menambahkan data ke array data. Program ini menggunakan Selection Sort. Setelah data diurutkan, program menghitung median. Jika ganjil, ambil nilai tengah. Jika genap, ambil rata rata dari dua nilai tengah. Program menampilkan hasil median yang telah dihitung.

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
// Andreas Besar Wibowo
// 2311102198 / IF-11-05

package main

import "fmt"

const nMax = 7919

type Buku struct {
	id       string
	judul    string
	penulis  string
	penerbit string
	eksemplar, tahun, rating int
}

type DaftarBuku [nMax]Buku

// Prosedur untuk mendaftarkan buku ke perpustakaan
func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Printf("\nMasukkan data buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit,
			&pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

// Prosedur untuk menampilkan buku dengan rating tertinggi
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}

	terfavorit := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > terfavorit.rating {
			terfavorit = pustaka[i]
		}
	}

	fmt.Println("\nBuku Terfavorit:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
		terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun, terfavorit.rating)
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
	fmt.Println("\n5 Judul Buku Dengan Rating Tertinggi:")
	count := 5
	if n < 5 {
		count = n
	}

	for i := 0; i < count; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
	}
}

// Prosedur untuk mencari buku berdasarkan ratingnya
func CariBuku(pustaka DaftarBuku, n, r int) {
	found := false
	fmt.Printf("\nBuku dengan Rating %d:\n", r)

	for i := 0; i < n; i++ {
		if pustaka[i].rating == r {
			found = true
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
				pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit,
				pustaka[i].tahun, pustaka[i].eksemplar, pustaka[i].rating)
		}
	}

	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu.")
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
	fmt.Print("\nMasukkan rating yang ingin dicari: ")
	fmt.Scan(&ratingCari)
	CariBuku(pustaka, n, ratingCari)
}
```
#### Screenshoot Source Code
![Unguided3 carbon](https://github.com/user-attachments/assets/236b1355-b607-43c5-821c-db39aa4aba08)

#### Screenshoot Output
![Unguided3 go](https://github.com/user-attachments/assets/3d43a2c0-8e6f-4dfe-8e30-41964f1a46f1)

#### Deskripsi Program
Program ini merupakan program untuk mengelola buku di perpustakaan. User diminta untuk mendaftarkan buku, menampilkan buku dengan rating tertinggi, mengurutkan buku berdasarkan rating, mencetak 5 buku dengan rating tertinggi, dan mencari buku dalam rating tertentu. Data buku meliputi ID, judul, penulis, penerbit, eksemplar, tahun, dan rating. Program ini menggunakan algoritma Insertion Sort.

Algoritma Program :
1. Minta input jumlah buku yang didaftarkan.
2. Masukkan data buku yang meliputi ID, judul, penulis, penerbit, eksemplar, tahun, dan rating.
3. Tampilkan buku dengan rating tertinggi.
4. Urutkan buku dengan descending.
5. Tampilkan 5 buku dengan rating tertinggi.
6. Minta rating buku yang ingin dicari.
7. Cari buku dengan rating dan tampilkan buku tersebut.

Cara kerja dari program ini yaitu meminta user untuk menginputkan jumlah buku yang ingin didaftarkan. User memasukkan data buku yang meliputi ID, judul, penulis, penerbit, eksemplar, tahun, dan rating. Program menampilkan buku dengan rating tertinggi. Mengurutkan buku dengan menggunakan Insertion Sort dan menampilkan 5 buku dengan rating tertinggi. User menginputkan rating buku yang ingin dicari. Program akan menampilkan informasi buku sesuai dengan rating buku yang ingin dicari oleh user.

## Referensi
[1] Kopicoding. Algoritma Selection Sort di Golang. Diakses dari https://www.kopicoding.com/algoritma-selection-sort-di-golang/.                                        

[2] Kopicoding. Insertion Sort di Golang. Diakses dari https://www.kopicoding.com/insertion-sort-di-golang/.
