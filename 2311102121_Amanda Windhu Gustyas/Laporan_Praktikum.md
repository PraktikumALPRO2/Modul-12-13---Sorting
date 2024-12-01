# <h2 align="center">LAPORAN PRAKTIKUM</h2>
# <h2 align="center">ALGORITMA DAN PEMROGRAMAN 2</h2>
# <h2 align="center">MODUL 12-13</h2>
# <h2 align="center">PENGURUTAN DATA</h2>
<p align="center">
    <img src="https://github.com/user-attachments/assets/3ccfed0b-72d1-4349-ac08-c4dce379c827" alt="Gambar">
</p>
 <h3  align="center" >Disusun Oleh : </h3>
<p align="center">Amanda Windhu Gustyas - 2311102121</p>
<p align="center">IF-11-05</p>
 <h3 <p align="center" >Dosen Pengampu : </h3> </p>
 <p align="center">Arif Amrulloh, S.Kom., M.Kom.</p>

# <h3 align="center"> PROGRAM STUDI S1 TEKNIK INFORMATIKA </h3>
# <h3 align="center"> FAKULTAS INFORMATIKA </h3>
# <h3 align="center"> TELKOM UNIVERSITY PURWOKERTO </h3>
# <h3 align="center"> 2024 </h3>

## I. DASAR TEORI

### 1. Ide Algoritma Selection Sort
Pengurutan secara seleksi ini idenya adalah mencari nilai ekstrim pada sekumpulan data, kemudian meletakkan pada posisi yang seharusnya. Pada penjelasan berikut ini data akan diurut membesar (ascending), dan data dengan indeks kecil ada di "kiri" dan indeks besar ada di "kanan".<br/>
    1. Cari nilai terkecil di dalam rentang data tersisa.<br/>
    2. Pindahkan/tukar tempat dengan data yang berada pada posisi paling kiri pada rentang data
    tersisa tersebut.<br/>
    3. Ulangi proses ini sampai tersisa hanya satu data saja.<br/>
Algoritma ini dikenal juga dengan nama Selection Sort, yang mana pada algoritma ini melibatkan dua proses yaitu pencarian indeks nilai ekstrim dan proses pertukaran dua nilai atau swap.<br/>
![image](https://github.com/user-attachments/assets/6cbe2132-e2db-49e2-a640-326880e80a74)<br/>

### 2. Algoritma Selection Sort
Adapun algoritma selection sort untuk mengurutkan array bertipe data bilangan bulat secara membesar atau ascending adalah sebagai berikut!<br/>
![image](https://github.com/user-attachments/assets/c4516483-2c5c-4884-8bb6-72e370cf2c5a)<br/>
![image](https://github.com/user-attachments/assets/ae918a0b-9ed9-4fc0-9d1a-0f2cac8bc139)<br/>
Sama halnya apabila array yang akan diurutkan adalah bertipe data struct, maka tambahkan filed pada saat proses perbandingan nilai ekstrim, kemudian tipe data dar variabel t sama dengan struct dari arraynya.<br/>
![image](https://github.com/user-attachments/assets/7a5b32c0-1b29-409a-a079-b1932021e72c)<br/>

### 3. Ide Algoritma Insertion Sort
Pengurutan secara insertion ini idenya adalah menyisipkan suatu nilai pada posisi yang seharusnya. Berbeda dengan pengurutan seleksi, yang mana pada pengurutan ini tidak dilakukan pencarian nilai ekstrim terlebih dahulu, cukup memilih suatu nilai tertentu kemudian mencari posisinya secara sequential search. Pada penjelasan berikut ini data akan diurut mengecil (descending), dan data dengan indeks kecil ada di "kiri" dan indeks besar ada di "kanan".<br/>
    1. Untuk satu data yang belum terurut dan sejumlah data yang sudah diurutkan:
    Geser data yang sudah terurut tersebut (ke kanan), sehingga ada satu ruang kosong untuk
    memasukkan data yang belum terurut ke dalam data yang sudah terurut dan tetap menjaga
    keterurutan.<br/>
    2. Ulangi proses tersebut untuk setiap data yang belum terurut terhadap rangkaian data yang
    sudah terurut.<br/>
Algoritma ini dikenal juga dengan nama Insertion Sort, yang mana pada algoritma ini melibatkan dua proses yaitu pencarian sekuensial dan penyisipan.<br/>
 ![image](https://github.com/user-attachments/assets/2c9d0508-75e1-4880-8bbc-07018965c451)<br/>

### 4. Algoritma Insertion Sort
Adapun algoritma insertion sort untuk mengurutkan array bertipe data bilangan bulat secara mengecil atau descending adalah sebagai berikut ini!<br/>
![image](https://github.com/user-attachments/assets/65ccfba8-e1d8-4c99-ab21-e2ae51727795)<br/>
Sama halnya apabila array yang akan diurutkan adalah bertipe data struct, maka tambahkan field pada saat proses perbandingan dalam pencarian posisi, kemudian tipe data dari variabel temp sama dengan struct dari arraynya.<br/>
![image](https://github.com/user-attachments/assets/7e0686a3-be59-44cd-ac24-289123575c2b)<br/>
![image](https://github.com/user-attachments/assets/9f06e676-06f5-46ef-989d-c056ed8c564e)<br/>

## II. GUIDED

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. <br/> Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort.<br/> Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat. <br/> Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing-masing daerah.<br/> ![image](https://github.com/user-attachments/assets/b3e2290f-db20-4a73-96a0-a16a733d537e) <br/> Keterangan: Terdapat 3 daerah dalam contoh input, dan di masing-masing daerah mempunyai 5, 6, dan 3 kerabat.

```go
package main

import (
	"fmt"
)

func SelectionSortAscending(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
	
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 0 || n >= 1000 {
		fmt.Println("Nilai n harus antara 1 dan 999.")
		return
	}

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		if m <= 0 || m >= 1000000 {
			fmt.Println("Nilai m harus antara 1 dan 999999.")
			return
		}

		
		arr := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		
		SelectionSortAscending(arr)

		
		for j, num := range arr {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(num)
		}
		fmt.Println()
	}
}
```
## Output:<br/> ![image](https://github.com/user-attachments/assets/dcad3bc1-ba98-42bf-a86b-583599c6bb07)

Program di atas adalah implementasi algoritma Selection Sort dalam bahasa Go, yang digunakan untuk mengurutkan bilangan dalam urutan ascending. Program diawali dengan membaca input jumlah kasus (n) dan memvalidasi agar nilainya berada dalam rentang 1 hingga 999. Setiap kasus kemudian meminta jumlah elemen dalam array (m), yang juga divalidasi agar berada dalam rentang 1 hingga 999999.

Setelah menerima jumlah elemen, program membaca elemen-elemen array, mengurutkannya menggunakan algoritma Selection Sort, lalu mencetak hasil array yang telah terurut. Proses pengurutan dilakukan dengan mencari elemen terkecil pada sisa array yang belum terurut, lalu menukarnya dengan elemen di posisi iterasi saat ini. Hasil akhirnya adalah semua elemen pada array diurutkan dengan benar untuk setiap kasus, dan outputnya ditampilkan dalam satu baris untuk setiap kasus.

### 2. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metode insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.<br/> Masukan: terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non-negatif saja yang disimpan ke dalam array.<br/> Keluaran: terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array, "Data berjarak x" atau "data berjarak tidak tetap".<br/> Contoh masukan dan keluaran<br/>
![image](https://github.com/user-attachments/assets/6f12f4a5-24e3-4053-b700-3eef22375cc9)

```go
package main

import (
	"fmt"
	"math"
)

// Fungsi untuk melakukan insertion sort
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Fungsi untuk memeriksa apakah data memiliki jarak tetap
func checkEqualSpacing(arr []int) string {
	if len(arr) < 2 {
		return "Data berjarak tidak tetap"
	}

	// Hitung jarak awal
	diff := int(math.Abs(float64(arr[1] - arr[0])))
	
	// Periksa jarak setiap elemen
	for i := 2; i < len(arr); i++ {
		if int(math.Abs(float64(arr[i]-arr[i-1]))) != diff {
			return "Data berjarak tidak tetap"
		}
	}
	return fmt.Sprintf("Data berjarak %d", diff)
}

func main() {
	var input []int
	var num int

	fmt.Println("Masukkan bilangan bulat:")

	// Input bilangan hingga ditemukan bilangan negatif
	for {
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		input = append(input, num)
	}

	// Lakukan sorting pada array
	insertionSort(input)

	// Cetak array yang telah diurutkan
	fmt.Println("Array setelah diurutkan:", input)

	// Periksa apakah data memiliki jarak tetap
	result := checkEqualSpacing(input)
	fmt.Println(result)
}
```
## Output:<br/> ![image](https://github.com/user-attachments/assets/f16d6998-8101-440a-88f7-95a65100ea55)

Program di atas membaca bilangan bulat hingga ditemukan bilangan negatif, mengurutkan bilangan tersebut dengan algoritma insertion sort, lalu memeriksa apakah jarak antar elemen di array yang sudah diurutkan konsisten. Hasil akhirnya menampilkan array yang sudah diurutkan dan status jaraknya.

## III. UNGUIDED 

### 1. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya di ujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.<br/> Format Masukan: masih persis sama seperti sebelumnya.<br/> Keluaran: terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.<br/> ![image](https://github.com/user-attachments/assets/950d0183-c431-448e-8ce8-b03a3651ca03) <br/> Keterangan: Terdapat 3 daerah dalam contoh masukan. Baris kedua berisi campuran bilangan ganjil dan genap. Baris berikutnya hanya berisi bilangan ganjil, dan baris terakhir hanya berisi bilangan genap.<br/> 
### Petunjuk:
### - Waktu pembacaan data, bilangan ganjil dan genap dipisahkan ke dalam dua array yang berbeda, untuk kemudian masing-masing diurutkan tersendiri.
### - Atau, tetap disimpan dalam satu array, diurutkan secara keseluruhan. Tetapi pada waktu pencetakan, mulai dengan mencetak semua nilai ganjil lebih dulu, kemudian setelah selesai cetaklah semua nilai genapnya.

```go
package main

import (
	"fmt"
)

// Fungsi untuk melakukan selection sort
func selectionSort(arr []int, ascending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if ascending && arr[j] < arr[minIndex] {
				minIndex = j
			} else if !ascending && arr[j] > arr[minIndex] {
				minIndex = j
			}
		}
		// Tukar elemen
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	var n, num int

	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan bilangan untuk daerah %d (akhiri dengan bilangan negatif):\n", i+1)
		var oddNumbers, evenNumbers []int

		// Input bilangan hingga ditemukan bilangan negatif
		for {
			fmt.Scan(&num)
			if num < 0 {
				break
			}
			if num%2 == 0 {
				evenNumbers = append(evenNumbers, num)
			} else {
				oddNumbers = append(oddNumbers, num)
			}
		}

		// Urutkan ganjil secara ascending dan genap secara descending
		selectionSort(oddNumbers, true)
		selectionSort(evenNumbers, false)

		// Cetak hasil
		fmt.Printf("Hasil untuk daerah %d:\n", i+1)
		for _, val := range oddNumbers {
			fmt.Print(val, " ")
		}
		for _, val := range evenNumbers {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}
```
## Output:<br/> ![image](https://github.com/user-attachments/assets/d8d6113c-87c2-4582-95a4-74dbed59165b)

Program di atas dibuat untuk membaca input bilangan dari beberapa daerah, memisahkan bilangan ganjil dan genap, lalu mengurutkannya sebelum mencetak hasilnya. Program meminta pengguna memasukkan jumlah daerah dan kemudian untuk setiap daerah, pengguna memasukkan bilangan satu per satu hingga bilangan negatif dimasukkan, yang menandakan akhir input. Bilangan ganjil dan genap dipisahkan ke dalam array terpisah. Program kemudian mengurutkan bilangan ganjil secara menaik dan bilangan genap secara menurun menggunakan algoritma selection sort. Setelah proses pengurutan selesai, program mencetak bilangan ganjil diikuti dengan bilangan genap dalam urutan yang telah diatur.

### 2. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemanya?<br/>
### "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rata-rata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rata-rata nilai tengah dibulatkan ke bawah."<br/>
### Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.
### Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313.
### Keluaran adalah median yang diminta, satu data per baris.<br/>
![image](https://github.com/user-attachments/assets/82520e81-fff5-4a78-a675-08b2d79af437)<br/>
### Keterangan:<br/> Sampai bilangan 0 yang pertama, data terbaca adalah 7 23 11, setelah tersusun: 7 11 23, maka median saat itu adalah 11.<br/> Sampai bilangan 0 yang kedua, data adalah 7 23 11 5 19 2 29 3 13 17, setelah tersusun diperoleh: 2 3 5 7 11 13 17 19 23 29. Karena ada 10 data, genap, maka median adalah (11+13)/2 = 12.<br/>
### Petunjuk:<br/> Untuk setiap data bukan 0 (dan bukan marker -5313541) simpan ke dalam array. Dan setiap kali menemukan bilangan 0, urutkanlah data yang sudah tersimpan dengan menggunakan metode insertion sort dan ambil mediannya.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var data []int // Menyimpan data yang sudah terbaca
	
	// Membaca input dari stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid")
			continue
		}

		// Jika angka -5313 ditemukan, hentikan program
		if num == -5313 {
			break
		}

		// Jika angka 0 ditemukan, hitung dan cetak median
		if num == 0 {
			if len(data) == 0 {
				fmt.Println("Tidak ada data untuk menghitung median")
				continue
			}

			selectionSort(data) // Mengurutkan data menggunakan selection sort
			median := calculateMedian(data)
			fmt.Printf("Mediannya adalah %d\n", median)
		} else {
			// Menambahkan data baru ke array
			data = append(data, num)
		}
	}
}

// Fungsi untuk menghitung median
func calculateMedian(arr []int) int {
	length := len(arr)
	if length%2 == 1 {
		// Jika jumlah data ganjil, ambil nilai tengah
		return arr[length/2]
	} else {
		// Jika jumlah data genap, ambil rata-rata dua nilai tengah dan bulatkan ke bawah
		sum := arr[length/2-1] + arr[length/2]
		return sum / 2
	}
}

// Fungsi untuk mengurutkan array menggunakan selection sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Tukar elemen minimum dengan elemen pertama
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
```
## Output:<br/> ![image](https://github.com/user-attachments/assets/a07dc677-be20-4bfc-9264-02586750fc5c)

Program ini menghitung median dari serangkaian bilangan bulat positif yang diberikan sebagai input. Median adalah nilai tengah dari data yang telah diurutkan. Jika jumlah data ganjil, median adalah elemen tengah, sedangkan jika jumlah data genap, median dihitung sebagai rata-rata dari dua elemen tengah yang dibulatkan ke bawah. Program membaca input secara berurutan, menyimpan bilangan yang valid ke dalam array. Setiap kali angka 0 ditemukan, program mengurutkan data yang telah terkumpul menggunakan algoritma selection sort, lalu menghitung dan mencetak median. Ketika angka -5313 ditemukan, program berhenti memproses input.

### 3. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:
![image](https://github.com/user-attachments/assets/2734f8f6-83de-48d1-a546-507f0e97096b)<br/>
### Masukan: terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.<br/>
### Keluaran: terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir.<br/>

```go
package main

import (
	"fmt"
	"sort"
)

type Buku struct {
	ID        string
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

type DaftarBuku []Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Scan(&buku.ID, &buku.Judul, &buku.Penulis, &buku.Penerbit, &buku.Eksemplar, &buku.Tahun, &buku.Rating)
		*pustaka = append(*pustaka, buku)
	}
}

func CetakTerfavorit(pustaka DaftarBuku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}
	terfavorit := pustaka[0]
	for _, buku := range pustaka {
		if buku.Rating > terfavorit.Rating {
			terfavorit = buku
		}
	}
	fmt.Println(terfavorit.Judul, terfavorit.Penulis, terfavorit.Penerbit, terfavorit.Tahun)
}

func UrutBuku(pustaka *DaftarBuku) {
	sort.Slice(*pustaka, func(i, j int) bool {
		return (*pustaka)[i].Rating > (*pustaka)[j].Rating
	})
}

func Cetak5Terbaru(pustaka DaftarBuku) {
	limit := 5
	if len(pustaka) < limit {
		limit = len(pustaka)
	}
	for i := 0; i < limit; i++ {
		fmt.Println(pustaka[i].Judul, pustaka[i].Rating)
	}
}

func CariBuku(pustaka DaftarBuku, rating int) {
	low, high := 0, len(pustaka)-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].Rating == rating {
			fmt.Println(pustaka[mid].Judul, pustaka[mid].Penulis, pustaka[mid].Penerbit, pustaka[mid].Tahun, pustaka[mid].Rating)
			return
		} else if pustaka[mid].Rating < rating {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var n int
	fmt.Scan(&n)

	var pustaka DaftarBuku
	DaftarkanBuku(&pustaka, n)

	CetakTerfavorit(pustaka)

	UrutBuku(&pustaka)

	Cetak5Terbaru(pustaka)

	var cariRating int
	fmt.Scan(&cariRating)
	CariBuku(pustaka, cariRating)
}
```
## Output:<br/> ![image](https://github.com/user-attachments/assets/2d31ed02-9a46-462a-8716-430e3f10497f)

Program di atas adalah aplikasi sederhana untuk mengelola data buku perpustakaan. Fitur utama meliputi DaftarkanBuku untuk memasukkan data buku, CetakTerfavorit untuk menampilkan buku dengan rating tertinggi, dan UrutBuku untuk mengurutkan buku berdasarkan rating secara menurun. Selain itu, program dapat menampilkan 5 buku dengan rating tertinggi melalui Cetak5Terbaru, serta mencari buku berdasarkan rating tertentu menggunakan pencarian biner dengan fitur CariBuku. 

