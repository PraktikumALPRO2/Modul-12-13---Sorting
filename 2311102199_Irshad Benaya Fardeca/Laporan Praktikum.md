<h2 align="center">LAPORAN PAKTIKUM ALGORITMA DAN PEMROGRAMAN 2</h2>
<h2 align="center">MODUL 12 & 13</h2>
<h2 align="center">PENGURUTAN DATA</h2>

<p align="center"><img src=https://github.com/user-attachments/assets/37e9c953-078b-4ef4-97e7-a5d25344e50b alt="Logo" width="200"/><p>

<p align="center">Disusun Oleh : </p>
<p align="center">Irshad Benaya Fardeca / 2311102199</p>
<p align="center">IF-11-05</p>
<br></br>
<p align="center">Dosen Pengampu : </p>
<p align="center">Arif Amrulloh, S.Kom., M.Kom.</p>
<br></br>
<h3 align="center">PROGRAM STUDI S1 TEKNIK INFORMATIKA</h3>
<h3 align="center">FAKULTAS INFORMATIKA</h3>
<h3 align="center">TELKOM UNIVERSITY PURWOKERTO</h3>
<h3 align="center">2024</p>

<br></br>

#### I. DASAR TEORI
#### 1. Selection Sort
Selection Sort adalah suatu metode pengurutan yang membandingkan elemen yang sekarang dengan elemen berikutnya sampai ke elemen yang terakhir. Jika ditemukan elemen lain yang lebih kecil dari elemen sekarang maka dicatat posisinya dan langsung ditukar. Metode selection sort adalah melakukan pemilihan dari suatu nilai yang terkecil dan kemudian menukarnya dengan elemen paling awal, lalu membandingkan dengan elemen yang sekarang dengan elemen berikutnya sampai dengan elemen terakhir, perbandingan dilakukan terus sampai tidak ada lagi pertukaran data.

#### 2. Insertion Sort
Insertion sort adalah suatu metode pengurutan data dengan cara menyimpan data ke suatu variabel sementara, kemudian dibandingkan dengan data-data lainnya yang ada disebelah kiri posisi data tersebut. Demikian seterusnya hingga data terakhir.


<br></br>


#### II. GUIDED

##### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort. 3 Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing- masing daerah. Keterangan: Terdapat 3 daerah dalam contoh input, dan di masing-masing daerah mempunyai 5, 6, dan 3 kerabat.
##### Source code
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
		// Tukar elemen Lerkecil dengan elemen indeks i
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

		fmt.Printf("Nomor rumah terurut untuk daerah %d : ", i+1)
		for _, num := range arr {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}

```
##### Screenshoot Output
![Screenshot 2024-12-01 195624](https://github.com/user-attachments/assets/356d68c6-f74a-4484-9b2e-62f2b1b74f3c)

##### Deskripsi Program
Program ini dapat mengurutkan nomor rumah kerabat berdasarkan daerahnya. User diminta untuk memasukkan jumlah daerah dan jumlah nomor rumah untuk setiap daerah. Kemudian, program akan mengurutkan nomor rumah masing-masing daerah menggunakan algoritma selection sort. Hasil pengurutan berupa daftar nomor rumah yang sudah terurut untuk setiap daerah akan ditampilkan.

##### 2. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya. Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array. 4 Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".
##### Source code
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
	fmt.Println("Masukkan data integer (akhiri dengan bilangan negatif):")
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
##### Screenshoot Output
![Screenshot 2024-12-01 200251](https://github.com/user-attachments/assets/a1198b66-4da0-4be6-a73d-ee475ce1fedc)

##### Deskripsi Program
Program ini dibuat untuk mengurutkan sekumpulan bilangan bulat yang diinput user, kemudian memeriksa apakah selisih antara bilangan-bilangan yang berdekatan dalam urutan tersebut selalu konstan.


<br></br>


#### III. UNGUIDED

##### 1. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.

Format Masukan masih persis sama seperti sebelumnya.

Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah. 

Keterangan: Terdapat 3 daerah dalam contoh masukan. Baris kedua berisi campuran bilangan ganjil dan genap. Baris berikutnya hanya berisi bilangan ganjil, dan baris terakhir hanya berisi bilangan genap.

Petunjuk: Waktu pembacaan data, bilangan ganjil dan genap dipisahkan ke dalam dua array yang berbeda, untuk kemudian masing-masing diurutkan tersendiri. Atau, tetap disimpan dalam satu array, diurutkan secara keseluruhan. Tetapi pada waktu pencetakan, mulai dengan mencetak semua nilai ganjil lebih dulu, kemudian setelah selesal cetaklah semua nilai genapnya.
#### Source Code
```go
package main

import "fmt"

func selectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {

			if (arr[j]%2 != 0 && arr[idxMin]%2 == 0) || (arr[j]%2 == arr[idxMin]%2 && arr[j] < arr[idxMin]) {
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

		fmt.Printf("Nomor rumah terurut untuk daerah %d : ", i+1)
		for _, num := range arr {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}

```
##### Screenshoot Output
![Screenshot 2024-12-01 230313](https://github.com/user-attachments/assets/bad4a128-c68a-4d36-80a9-efbaf6906667)

##### Deskripsi Program
Program ini dibuat untuk mengurutkan nomor rumah kerabat dari berbagai daerah. User akan diminta memasukkan jumlah daerah dan jumlah nomor rumah untuk setiap daerah. Kemudian, program akan mengurutkan nomor rumah tersebut berdasarkan dua kriteria: nomor rumah ganjil akan diprioritaskan lebih dulu, dan jika ada nomor rumah dengan paritas yang sama, maka nomor rumah yang lebih kecil akan diletakkan di depan.

##### 2. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?

"Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah."

Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.

Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan O. Data O merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313.

Keluaran adalah median yang diminta, satu data per baris.

Keterangan: Sampai bilangan O yang pertama, data terbaca adalah 7 23 11, setelah tersusun: 7 11 23, maka median saat itu adalah 11.
Sampai bilangan O yang kedua, data adalah 7 23 11 5 19 2 29 3 13 17, setelah tersusun diperoleh: 2 3 5 7 11 13 17 19 23 29. Karena ada 10 data, genap, maka median adalah (11+13)/2=12.

Petunjuk: Untuk setiap data bukan 0 (dan bukan marker -5313541) simpan ke dalam array, Dan setiap kali menemukan bilangan O, urutkanlah data yang sudah tersimpan dengan menggunakan metode insertion sort dan ambil mediannya.
#### Source Code
```go
package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func median(arr []int) int {
	n := len(arr)
	if n%2 == 0 {
		return (arr[n/2-1] + arr[n/2]) / 2
	} else {
		return arr[n/2]
	}
}

func main() {
	var data []int
	var bil int

	for {
		fmt.Scan(&bil)
		if bil == -5313 {
			break
		}
		if bil == 0 {
			insertionSort(data)
			fmt.Println(median(data))
		} else {
			data = append(data, bil)
		}
	}
}
```
##### Screenshoot Output
![Screenshot 2024-12-01 231020](https://github.com/user-attachments/assets/2a7ab9db-e75b-48d1-a716-389298c3187b)

##### Deskripsi Program
Program ini dapat menghitung median dari serangkaian bilangan bulat yang dimasukkan oleh user. User dapat terus memasukkan bilangan hingga memasukkan nilai -5313 sebagai tanda berhenti. Setiap kali user memasukkan angka 0, program akan mengurutkan semua bilangan yang telah dimasukkan menggunakan algoritma Insertion Sort, kemudian menghitung dan menampilkan nilai mediannya.

##### 3. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:
![Screenshot 2024-12-01 202609](https://github.com/user-attachments/assets/536f36d1-d375-40e9-b7de-3d51f2869f29)

Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari. Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir.
Lengkapi subprogram-subprogram dibawah ini, sesuai dengan I.S. dan F.S yang diberikan.
<br><br>
![Screenshot 2024-12-01 202637](https://github.com/user-attachments/assets/6763eb43-3193-4863-b618-4f47eb7c9d1c)
#### Source Code
```go
package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan ID: ")
		fmt.Scan(&pustaka[i].id)
		fmt.Printf("Masukkan Judul: ")
		fmt.Scan(&pustaka[i].judul)
		fmt.Printf("Masukkan Penulis: ")
		fmt.Scan(&pustaka[i].penulis)
		fmt.Printf("Masukkan Penerbit: ")
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Printf("Masukkan Tahun: ")
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Printf("Masukkan Eksemplar: ")
		fmt.Scan(&pustaka[i].tahun)
		fmt.Printf("Masukkan Rating: ")
		fmt.Scan(&pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka *DaftarBuku, n int) {
	maxRating := pustaka[0].rating
	fav := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > maxRating {
			maxRating = pustaka[i].rating
			fav = pustaka[i]
		}
	}
	fmt.Printf("Buku Tervavorit : %s, %s, %s, %d\n", fav.judul, fav.penulis, fav.penerbit, fav.tahun)
}

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

func Cetak5Terbaru(pustaka *DaftarBuku, n int) {
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 0; i < 5 && i < n; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka *DaftarBuku, n int, r int) {
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if pustaka[mid].rating == r {
			fmt.Printf("Buku Ditemukan: %s, %s, %s, %d, %d, %d\n",
				pustaka[mid].judul, pustaka[mid].penulis, pustaka[mid].penerbit,
				pustaka[mid].tahun, pustaka[mid].eksemplar, pustaka[mid].rating)
			return
		}
		if pustaka[mid].rating > r {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating ", r)
}

func main() {
	var pustaka DaftarBuku
	var n, rating int

	fmt.Print("Jumlah Buku : ")

	DaftarkanBuku(&pustaka, &n)
	fmt.Println("")

	CetakTerfavorit(&pustaka, n)
	fmt.Println("")

	Cetak5Terbaru(&pustaka, n)
	fmt.Println("")

	UrutBuku(&pustaka, n)

	fmt.Print("Rating buku yang ingin dicari : ")
	fmt.Scan(&rating)
	CariBuku(&pustaka, n, rating)
}

```
##### Screenshoot Output
![Screenshot 2024-12-01 232843](https://github.com/user-attachments/assets/5ee12109-f56f-436c-a2c0-fb6e584b02b4)

##### Deskripsi Program
Program ini dirancang untuk mengelola data buku dalam sebuah perpustakaan. User dapat memasukkan data buku secara manual, lalu program akan mengurutkan buku berdasarkan rating, mencari buku berdasarkan rating, serta menampilkan buku dengan rating tertinggi. Fitur-fitur utamanya meliputi penambahan data buku, pencarian buku berdasarkan rating, pengurutan buku berdasarkan rating, dan tampilan buku dengan rating tertinggi.

### Referensi
[1] Sandria, Y. A., Nurhayoto, M. R. A., Ramadhani, L., Harefa, R. S., & Syahputra, A. (2023). Penerapan Algoritma Selection Sort untuk Melakukan Pengurutan Data dalam Bahasa Pemrograman PHP. Hello World Jurnal Ilmu Komputer, 1(4), 190-194.
<br>
[2] Sunandar, E. (2019). Perbandingan Metode Selection Sort dan Insertion Sort Dalam Pengurutan Data Menggunakan Bahasa Program Java. PETIR.
