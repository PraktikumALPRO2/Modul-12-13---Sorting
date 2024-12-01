
<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br> 

<h2 align="center"><strong>MODUL 12&13</strong></h2>
<h2 align="center"><strong> SORTING </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Fatik Nurimamah / 2311102190<br>
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
2. [Guided](#guided)
3. [Unguided](#unguided)
4. [Daftar Pustaka](#daftar-pustaka)


## Dasar Teori
**Algoritma Insertion Sort**

Algoritma insertion sort, adalah metode pengurutan dengan cara menyisipkan elemen data pada posisi yang tepat. Pencarian posisi yang tepat dilakukan dengan melakukan pencarian berurutan didalam barisan elemen, selama pencarian posisi yang tepat dilakukan pergeseran elemen. Pengurutan insertion sort sangat mirip dengan konsep permainan kartu, bahwa setiap kartu disisipkan secara berurutan dari kiri ke kanan sesuai dengan besar nilai kartu tersebut, dengan syarat apabila sebuah kartu disisipkan pada posisi tertentu kartu yang lain akan bergeser maju atau mundur sesuai dengan besaran nilai yang dimiliki.

Cara Kerja:
1. Mulai dari elemen kedua, bandingkan elemen saat ini dengan elemen sebelumnya.
2. Jika elemen saat ini lebih kecil, geser elemen yang lebih besar ke kanan.
3. Sisipkan elemen saat ini ke posisi yang benar.
4. Ulangi proses hingga seluruh array terurut.

```go
package main

import "fmt"

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

func main() {
	arr := []int{12, 11, 13, 5, 6}
	insertionSort(arr)
	fmt.Println("Sorted array:", arr)
}
```

**Algoritma Selection Sort**

Algoritma selection sort sering juga disebut dengan metode maksimum atau minimum. Metode maksimum karena didasarkan pada pemilihan data atau elemen maksimum sebagai dasar pengurutan. Konsepnya dengan memilih elemen maksimum kemudian mempertukarkan elemen maksimum tersebut dengan elemen paling akhir untuk urutan ascending dan elemen pertama untuk descending.

Algoritma selection sort disebut juga dengan metode minimum karena didasarkan pada pemilihan elemen minimum sebagai dasar pengurutan. Konsepnya dengan memilih elemen minimum kemudian mempertukarkan elemen minimum dengan elemen paling akhir untuk urutan ascending dan elemen pertama untuk urutan descending. Proses yang dilakukan oleh algoritma selection sort adalah mengambil nilai terbesar dari susunan data dan
menggantikannya dengan data yang paling kanan.

Cara Kerja:
1. Temukan elemen terkecil dalam array.
2. Tukar elemen terkecil dengan elemen pertama.
3. Ulangi proses untuk bagian array yang tersisa hingga seluruh array terurut.

```go
package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	selectionSort(arr)
	fmt.Println("Sorted array:", arr)
}
```
  

## Guided 

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu.

Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort.

Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah Integer m (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat.

Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing- masing daerah.
### Source Code :
```go
package main
import (
	"fmt"
)

//fungsi untuk mengurutkan array menggunakan selection sort
func SelectionSort(arr[]int, n int){
	for i := 0; i < n-1; i++{
		idxMin := i
		for j := i + 1; j < n; j++{
			//cari elemen terkecil
			if arr[j] < arr[idxMin]{
				idxMin = j
			}
		}
		//tuker elemen terkecil dengan elemen di posisi i
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func main(){
	var n int
	fmt.Printf("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	//proses tiap daerah
	for daerah := 1; daerah <= n; daerah++{
		var m int
		fmt.Printf("\nMasukkan jumlah nomor kerabat untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		//membaca nomor rumah untuk daerah ini
		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++{
			fmt.Scan(&arr[i])
		}

		//urutkan array dari terkecil ke terbesar
		SelectionSort(arr, m)

		// Tampilkan hasil
		fmt.Printf("Nomor rumah terurut untuk daerah %d: ", daerah)
		for i, num := range arr {
    	    if i > 0 {
            	 fmt.Printf(" ") 
   	        }
   		    fmt.Printf("%d", num)
        }
 		fmt.Println()

	}
}
```

### Output:
![image](https://github.com/user-attachments/assets/dcb3ba99-4bf7-408d-a860-e1de5e22ca92)

### Full code Screenshot:
![code](https://github.com/user-attachments/assets/6e79ab2f-7c08-4005-b420-312598c1254f)

### Deskripsi Program : 
Program ini adalah program yang dibuat untuk mengurutkan nomor rumah kerabat berdasarkan wilayah. Pengguna diminta untuk memasukkan jumlah wilayah kerabat (n) yang ingin diproses. Selanjutnya, untuk setiap wilayah, pengguna akan menginput jumlah nomor rumah kerabat (m) dan daftar nomor rumah tersebut. Program ini menggunakan algoritma Selection Sort untuk mengurutkan nomor-nomor rumah dari yang terkecil hingga yang terbesar. Setelah pengurutan selesai, program menampilkan hasil nomor rumah yang telah diurutkan untuk setiap wilayah secara terpisah di layar.

### 2. Buatlah sebuah program yang digunakan untuk membaca data Integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.

Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array.

Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".

### Source Code :
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
		fmt.Println("Data berjarak tidak tetap")
	}
}

```

### Output:
![image](https://github.com/user-attachments/assets/4eab52de-99aa-4f1b-a859-bd3950c25ebd)

### Full code Screenshot:
![code](https://github.com/user-attachments/assets/02f22c16-ed8f-429e-bc09-0a53801004ff)

### Deskripsi Program : 
Program ini adalah program yang dibuat untuk mengolah sekumpulan bilangan integer, mengurutkannya, dan memeriksa apakah selisih antar elemen dalam array memiliki nilai yang sama secara konsisten. Pengguna diminta untuk memasukkan bilangan, dan proses input akan berhenti ketika bilangan negatif dimasukkan. Data yang diterima akan diurutkan menggunakan algoritma Insertion Sort, yang menyisipkan elemen ke posisi yang sesuai berdasarkan perbandingan nilai. Setelah data diurutkan, program akan mengecek apakah selisih antara elemen-elemen yang berdekatan dalam array selalu sama. Jika pola tersebut ditemukan, program akan menampilkan nilai selisih tersebut; jika tidak, program akan menginformasikan bahwa selisihnya tidak konsisten. Aplikasi ini sangat berguna untuk menganalisis pola data numerik sekaligus menyusunnya dalam urutan teratur.


## Unguided 

### 1. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.

Format Masukan masih persis sama seperti sebelumnya.

Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.

Petunjuk:

- Waktu pembacaan data, bilangan ganjil dan genap dipisahkan ke dalam dua array yang berbeda, untuk kemudian masing-masing diurutkan tersendiri.

- Atau, tetap disimpan dalam satu array, diurutkan secara keseluruhan. Tetapi pada waktu pencetakan, mulai dengan mencetak semua nilai ganjil lebih dulu, kemudian setelah selesai cetaklah semua nilai genapnya.
### Source Code :
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi selection sort untuk mengurutkan array
func selectionSort(angka []int, ascending bool) {
	panjang := len(angka)
	for i := 0; i < panjang-1; i++ {
		indeksEkstrem := i
		for j := i + 1; j < panjang; j++ {
			if (ascending && angka[j] < angka[indeksEkstrem]) || (!ascending && angka[j] > angka[indeksEkstrem]) {
				indeksEkstrem = j
			}
		}
		angka[i], angka[indeksEkstrem] = angka[indeksEkstrem], angka[i]
	}
}

func pisahDanUrutkan(angka []int) ([]int, []int) {
	var ganjil, genap []int

	for _, nilai := range angka {
		if nilai%2 == 0 {
			genap = append(genap, nilai)
		} else {
			ganjil = append(ganjil, nilai)
		}
	}

	selectionSort(ganjil, false) // Ganjil descending
	selectionSort(genap, true)  // Genap ascending

	return ganjil, genap
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Memasukkan jumlah daerah
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	jumlahDaerah, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error: Input jumlah daerah tidak valid!")
		return
	}

	for i := 1; i <= jumlahDaerah; i++ {
		// Memasukkan jumlah nomor rumah untuk daerah ini
		fmt.Printf("\nMasukkan jumlah nomor kerabat untuk daerah %d: ", i)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		jumlahNomor, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: Input jumlah nomor tidak valid!")
			return
		}

		// Memasukkan nomor rumah
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", jumlahNomor)
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		stringAngka := strings.Split(input, " ")

		var angkaDaerah []int
		for _, stringNilai := range stringAngka {
			angka, err := strconv.Atoi(stringNilai)
			if err == nil {
				angkaDaerah = append(angkaDaerah, angka)
			} else {
				fmt.Printf("Error: '%s' bukan angka yang valid\n", stringNilai)
			}
		}

		// Memastikan jumlah input sesuai dengan jumlah nomor yang dimasukkan
		if len(angkaDaerah) != jumlahNomor {
			fmt.Printf("Error: jumlah nomor yang dimasukkan (%d) tidak sesuai dengan jumlah yang diharapkan (%d)\n", len(angkaDaerah), jumlahNomor)
			continue
		}

	
		// Mengurutkan dan memisahkan angka ganjil dan genap
		ganjil, genap := pisahDanUrutkan(angkaDaerah)

		// Menampilkan hasil urutan langsung di bawah input
		fmt.Printf("Ganjil (descending): %v\n", ganjil)
		fmt.Printf("Genap (ascending): %v\n", genap)
	}
}

```
### Output:
![image](https://github.com/user-attachments/assets/17665ec6-85c1-48ee-b30c-8b35277c7a55)

### Full code Screenshot:
![code](https://github.com/user-attachments/assets/b2174d59-e1c2-43b4-9bfd-113027020c3f)

### Deskripsi Program : 
Program ini adalah program yang digunakan untuk memisahkan dan mengurutkan bilangan ganjil dan genap dari sejumlah data yang diberikan. Pengguna diminta untuk menentukan jumlah daerah, kemudian memasukkan deretan angka untuk masing-masing daerah dengan format angka yang dipisahkan oleh spasi. Program akan memisahkan angka-angka tersebut menjadi dua kelompok: bilangan ganjil dan bilangan genap. Bilangan ganjil diurutkan secara menurun (descending), sedangkan bilangan genap diurutkan secara menaik (ascending) menggunakan algoritma Selection Sort. Hasil pengurutan ditampilkan untuk setiap daerah, dengan mekanisme validasi input untuk memastikan hanya bilangan yang valid yang akan diproses.


### 2. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setlap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?

"Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilal tengah dibulatkan ke bawah."

elkom iversiti Buatlah program median yang mencetak nilal median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.

Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan O. Data O merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat-5313.

Keluaran adalah median yang diminta, satu data per baris.

Petunjuk:

Untuk setiap data bukan 0 (dan bukan marker -5313541) simpan ke dalam array, Dan setiap kali menemukan bilangan 0, urutkanlah data yang sudah tersimpan dengan menggunakan metode insertion sort dan ambil mediannya.
### Source Code :
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan data (akhiri dengan -5313): ")
	scanner.Scan()
	inputData := scanner.Text()

	data := strings.Fields(inputData)
	var processedNumbers []int
	for _, value := range data {
		number, _ := strconv.Atoi(value)
		if number == -5313 {
			break
		}
		if number == 0 {
			// Ketika menemukan angka 0, hitung median dari angka yang telah terkumpul
			if len(processedNumbers) > 0 {
				sort.Ints(processedNumbers)
				medianValue := CariMedian(processedNumbers)
				fmt.Printf("Median dari data saat ini (%v) adalah: %d\n", processedNumbers, medianValue)
			}
		} else {
			processedNumbers = append(processedNumbers, number)
		}
	}
}

func CariMedian(numbers []int) int {
	length := len(numbers)
	if length%2 == 1 {
		return numbers[length/2]
	}
	// Untuk jumlah data genap, ambil rata-rata dua nilai tengah
	return (numbers[length/2-1] + numbers[length/2]) / 2
}

```
### Output:
![image](https://github.com/user-attachments/assets/50139e46-7157-46de-990c-afafb71c6335)

### Full code Screenshot:
![code](https://github.com/user-attachments/assets/45c8ff16-bb41-4aca-8da9-b9ccc8febf2a)

### Deskripsi Program : 
Program ini adalah program yang dibuat untuk membaca dan memproses input angka secara bertahap dari pengguna, kemudian menghitung median pada kondisi tertentu. Pengguna memasukkan sejumlah angka, dengan -5313 sebagai tanda untuk menghentikan proses input. Setiap kali angka 0 dimasukkan, program akan menghitung median dari angka-angka yang telah dimasukkan sebelumnya, kecuali angka 0 dan -5313. Perhitungan median dilakukan dengan mengurutkan angka-angka tersebut, lalu menentukan nilai tengah (atau rata-rata dua nilai tengah jika jumlah data genap). Program ini berguna untuk memahami cara pengolahan data secara sekuensial, penerapan konsep statistik sederhana, dan pengelolaan input pengguna dalam aplikasi berbasis teks.

   
### 3. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:

![image](https://github.com/user-attachments/assets/37c53aaa-b298-4c03-81ed-02fdf683c1c3)


Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.

Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima Judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesual rating yang diberikan pada masukan baris terakhir.

Lengkapi subprogram-subprogram dibawah ini, sesuai dengan I.S. dan F.S yang diberikan:

![image](https://github.com/user-attachments/assets/88e2bef0-7e59-47c9-a417-27040216a3e4)


### Source Code :
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

// Struct untuk menyimpan data buku
type Buku struct {
	KodeBuku      string
	NamaBuku      string
	Pengarang     string
	Penerbit      string
	JumlahEksemplar int
	TahunTerbit   int
	NilaiRating   int
}

// Struct untuk menyimpan daftar buku
type KoleksiBuku struct {
	KumpulanBuku []Buku
	TotalBuku    int
}

// Procedure DaftarkanBuku
func DaftarkanBuku(pustaka *KoleksiBuku, n int) {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("\nMasukkan data untuk buku ke-%d:\n", i+1)

		// Input kode buku
		fmt.Print("Kode Buku: ")
		buku.KodeBuku, _ = reader.ReadString('\n')
		buku.KodeBuku = strings.TrimSpace(buku.KodeBuku)

		// Input nama buku
		fmt.Print("Nama Buku: ")
		buku.NamaBuku, _ = reader.ReadString('\n')
		buku.NamaBuku = strings.TrimSpace(buku.NamaBuku)

		// Input pengarang
		fmt.Print("Pengarang: ")
		buku.Pengarang, _ = reader.ReadString('\n')
		buku.Pengarang = strings.TrimSpace(buku.Pengarang)

		// Input penerbit
		fmt.Print("Penerbit: ")
		buku.Penerbit, _ = reader.ReadString('\n')
		buku.Penerbit = strings.TrimSpace(buku.Penerbit)

		// Input jumlah eksemplar
		fmt.Print("Jumlah Eksemplar: ")
		eksemplar, _ := reader.ReadString('\n')
		buku.JumlahEksemplar, _ = strconv.Atoi(strings.TrimSpace(eksemplar))

		// Input tahun terbit
		fmt.Print("Tahun Terbit: ")
		tahun, _ := reader.ReadString('\n')
		buku.TahunTerbit, _ = strconv.Atoi(strings.TrimSpace(tahun))

		// Input nilai rating
		fmt.Print("Nilai Rating: ")
		rating, _ := reader.ReadString('\n')
		buku.NilaiRating, _ = strconv.Atoi(strings.TrimSpace(rating))

		// Tambahkan buku ke dalam koleksi
		pustaka.KumpulanBuku = append(pustaka.KumpulanBuku, buku)
	}
	pustaka.TotalBuku = n
}

// Procedure CetakTerfavorit
func CetakTerfavorit(pustaka KoleksiBuku) {
	if len(pustaka.KumpulanBuku) == 0 {
		fmt.Println("Tidak ada buku dalam koleksi.")
		return
	}

	maxRating := -1
	var favorit Buku
	for _, buku := range pustaka.KumpulanBuku {
		if buku.NilaiRating > maxRating {
			maxRating = buku.NilaiRating
			favorit = buku
		}
	}

	fmt.Printf("Buku terfavorit:\nNama: %s, Pengarang: %s, Penerbit: %s, Tahun: %d\n",
		favorit.NamaBuku, favorit.Pengarang, favorit.Penerbit, favorit.TahunTerbit)
}

// Procedure UrutBuku (menggunakan Insertion Sort)
func UrutBuku(pustaka *KoleksiBuku) {
	sort.Slice(pustaka.KumpulanBuku, func(i, j int) bool {
		return pustaka.KumpulanBuku[i].NilaiRating > pustaka.KumpulanBuku[j].NilaiRating
	})
}

// Procedure Cetak5Terbaru
func Cetak5Terbaru(pustaka KoleksiBuku) {
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	limit := 5
	if len(pustaka.KumpulanBuku) < 5 {
		limit = len(pustaka.KumpulanBuku)
	}
	for i := 0; i < limit; i++ {
		buku := pustaka.KumpulanBuku[i]
		fmt.Printf("%d. Nama: %s, Rating: %d\n", i+1, buku.NamaBuku, buku.NilaiRating)
	}
}

// Procedure CariBuku
func CariBuku(pustaka KoleksiBuku, r int) {
	for _, buku := range pustaka.KumpulanBuku {
		if buku.NilaiRating == r {
			fmt.Printf("Buku ditemukan:\nNama: %s, Pengarang: %s, Penerbit: %s, Tahun: %d\n",
				buku.NamaBuku, buku.Pengarang, buku.Penerbit, buku.TahunTerbit)
			return
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka KoleksiBuku
	var jumlahBuku int

	// Masukkan data buku
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scanln(&jumlahBuku)
	DaftarkanBuku(&pustaka, jumlahBuku)

	// Cetak buku terfavorit
	CetakTerfavorit(pustaka)

	// Urutkan buku berdasarkan rating
	UrutBuku(&pustaka)

	// Cetak 5 buku dengan rating tertinggi
	Cetak5Terbaru(pustaka)

	// Cari buku berdasarkan rating
	var nilaiRating int
	fmt.Print("Masukkan rating untuk mencari buku: ")
	fmt.Scanln(&nilaiRating)
	CariBuku(pustaka, nilaiRating)
}

```
### Output:
![image](https://github.com/user-attachments/assets/7a710bcd-c0b9-4ec6-ac1e-b37ca502dd4e)

![image](https://github.com/user-attachments/assets/b8c888ba-51b1-45e8-a65f-35743a3fc8b6)


### Full code Screenshot:
![code](https://github.com/user-attachments/assets/382eb1df-dfd8-46d2-93b8-4040e81042e0)

### Deskripsi Program : 
Program ini adalah program yang dibuat untuk mengelola koleksi buku, dengan berbagai fitur seperti pendaftaran, pencarian, pengurutan, dan pencetakan data buku berdasarkan kriteria tertentu. Pengguna dapat memasukkan informasi buku, seperti kode, nama, pengarang, penerbit, jumlah eksemplar, tahun terbit, dan rating. Setelah data dimasukkan, program akan menampilkan buku dengan rating tertinggi (terfavorit) dan mencetak lima buku teratas berdasarkan rating. Pengguna juga bisa mencari buku dengan rating tertentu, dan program akan menampilkan informasi buku yang sesuai jika ditemukan.



## Daftar Pustaka
[1] Retnoningsih, E. (2018). Algoritma pengurutan data (sorting) dengan metode insertion sort dan selection sort. Information Management For Educators And Professionals: Journal of Information Management, 3(1), 95-106.
