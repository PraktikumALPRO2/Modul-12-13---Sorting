<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br> 

<h2 align="center"><strong>MODUL XII & XIII</strong></h2>
<h2 align="center"><strong> SORTING</strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Ria Wulandari / 2311102173<br>
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

## Dasar Teori
### Insertion sort
Insertion Sort adalah algoritma pengurutan yang bekerja dengan menyusun elemen-elemen dari array satu per satu. Setiap elemen dibandingkan dengan elemen-elemen yang sudah terurut sebelumnya, kemudian ditempatkan pada posisi yang sesuai. Proses ini dilakukan hingga seluruh elemen array terurut.

Pada dasarnya, algoritma ini melibatkan perulangan untuk membandingkan elemen yang sedang diproses dengan elemen-elemen sebelumnya. Jika elemen sebelumnya lebih besar, elemen-elemen tersebut akan digeser untuk memberikan ruang hingga posisi yang tepat ditemukan.

Kompleksitas waktu dalam kasus terbaik adalah O(n) jika array hampir terurut, sementara dalam kasus terburuk kompleksitasnya adalah O(n²).
#### Contoh implementasi insertion sort
```go
func insertionSort(arr []int) {
    for i := 1; i < len(arr); i++ {
        key := arr[i]
        j := i - 1
        // Geser elemen yang lebih besar ke kanan
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = key
    }
}
```
### Selection sort
Selection Sort adalah algoritma pengurutan yang bekerja dengan memilih elemen terkecil dari array yang belum terurut, lalu menukarnya dengan elemen pada posisi awal dari array tersebut. Proses ini dilakukan secara berulang untuk elemen-elemen berikutnya hingga seluruh array terurut.

Dalam penerapannya, algoritma ini menggunakan dua perulangan: perulangan luar untuk menentukan posisi elemen yang akan diisi, dan perulangan dalam untuk menemukan elemen terkecil dari sisa array yang belum terurut. Setelah elemen terkecil ditemukan, dilakukan penukaran dengan elemen pada posisi awal dari bagian yang belum terurut.

Kompleksitas waktu dari algoritma ini adalah O(n²), baik dalam kasus terbaik maupun terburuk, karena jumlah perbandingan yang harus dilakukan selalu sama.
#### Contoh implementasi selection sort
```go
func selectionSort(arr []int) {
    for i := 0; i < len(arr)-1; i++ {
        minIndex := i
        for j := i + 1; j < len(arr); j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }
        // Tukar elemen terkecil dengan elemen pada indeks i
        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }
}
```

## Guided
### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu.
Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort.

Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah Integer m (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat.

Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing- masing daerah.
#### Sorce code :
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
	fmt.Printf(" masukan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	//proses tiap daerah
	for daerah := 1; daerah <= n; daerah++{
		var m int
		fmt.Printf("\nMasukan jumlah nomor kerabat untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		//membaca nomor rumah untuk daerah ini
		arr := make([]int, m)
		fmt.Printf("masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++{
			fmt.Scan(&arr[i])
		}

		//urutkan array dari terkecil ke terbesar
		SelectionSort(arr, m)

		//tampilkan hasil
		fmt.Printf("nomor rumah terurut untuk daerah %d: ", daerah)
		for _, num := range arr {
			fmt.Printf("%d", num)
		}
		fmt.Println()
	}
}
```
#### Output
![image](https://github.com/user-attachments/assets/9f4d15a2-d780-4e22-b8e4-1482d7a66511)
#### Deskripsi program

### 2. Buatlah sebuah program yang digunakan untuk membaca data Integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.
Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array.

Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".
#### Source code :
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
#### Output
![image](https://github.com/user-attachments/assets/af7b9dd5-ab71-420e-b8a2-163dcc147e8a)

#### Deskripsi program

## Unguided
### 1. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.
Format Masukan masih persis sama seperti sebelumnya.

Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.

Petunjuk:

- Waktu pembacaan data, bilangan ganjil dan genap dipisahkan ke dalam dua array yang berbeda, untuk kemudian masing-masing diurutkan tersendiri.

- Atau, tetap disimpan dalam satu array, diurutkan secara keseluruhan. Tetapi pada waktu pencetakan, mulai dengan mencetak semua nilai ganjil lebih dulu, kemudian setelah selesai cetaklah semua nilai genapnya.
#### Source Code :
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi untuk insertion sort (ascending)
func insertionSortAsc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

// Fungsi untuk insertion sort (descending)
func insertionSortDesc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] < key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Masukkan jumlah baris input
	fmt.Print("Masukkan jumlah baris masukan: ")
	scanner.Scan()
	t, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Input jumlah baris tidak valid.")
		return
	}

	// Proses setiap baris masukan
	for i := 1; i <= t; i++ {
		fmt.Printf("Masukkan baris ke-%d (pisahkan dengan spasi): ", i)
		scanner.Scan()
		input := scanner.Text()

		// Pisahkan input ke dalam array bilangan
		numbers := strings.Fields(input)
		var ganjil, genap []int
		for _, num := range numbers {
			val, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Bilangan '%s' tidak valid, diabaikan.\n", num)
				continue
			}
			if val%2 == 0 {
				genap = append(genap, val)
			} else {
				ganjil = append(ganjil, val)
			}
		}

		// Urutkan bilangan ganjil (ascending) dan genap (descending)
		ganjil = insertionSortAsc(ganjil)
		genap = insertionSortDesc(genap)

		// Tampilkan hasil
		fmt.Printf("Hasil untuk baris ke-%d:\n", i)
		fmt.Println("Ganjil terurut membesar:", ganjil)
		fmt.Println("Genap terurut mengecil:", genap)
		fmt.Println()
	}
}

```
#### Output
![image](https://github.com/user-attachments/assets/66c9736c-3664-4afc-b142-a63a995529b1)

#### Deskripsi program
Program ini bertujuan untuk membantu Hercules mengatur nomor rumah di jalan sesuai aturannya. Setiap masukan terdiri dari beberapa baris nomor rumah. Program akan memisahkan bilangan ganjil dan genap, lalu mengurutkannya. Bilangan ganjil akan diurutkan membesar, sedangkan bilangan genap akan diurutkan mengecil. Hasil akhirnya akan ditampilkan dalam format sesuai kebutuhan Hercules.

### 2. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setlap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?
"Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilal tengah dibulatkan ke bawah."

elkom iversiti Buatlah program median yang mencetak nilal median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.

Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan O. Data O merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat-5313.

Keluaran adalah median yang diminta, satu data per baris.

Petunjuk:

Untuk setiap data bukan 0 (dan bukan marker -5313541) simpan ke dalam array, Dan setiap kali menemukan bilangan 0, urutkanlah data yang sudah tersimpan dengan menggunakan metode insertion sort dan ambil mediannya.

#### Source Code :
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var data []int

	fmt.Println("Masukkan angka satu per satu. Input '0' untuk mencetak median. Input '-5313' untuk keluar.")

	for {
		// Membaca input
		scanner.Scan()
		input := scanner.Text()

		// Konversi input menjadi integer
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid, masukkan angka!")
			continue
		}

		// Periksa apakah input adalah -5313 untuk keluar
		if num == -5313 {
			fmt.Println("Program selesai.")
			break
		}

		// Jika input adalah 0, hitung median
		if num == 0 {
			if len(data) == 0 {
				fmt.Println("Tidak ada data untuk menghitung median.")
				continue
			}

			// Salin data untuk pengurutan
			tempData := make([]int, len(data))
			copy(tempData, data)

			// Urutkan data
			sort.Ints(tempData)

			// Hitung median
			median := calculateMedian(tempData)

			// Tampilkan hasil median
			fmt.Printf("Data terurut: %v\n", tempData)
			fmt.Printf("Median: %d\n\n", median)
			continue
		}

		// Tambahkan input ke data
		data = append(data, num)
	}
}

// Fungsi untuk menghitung median dari array integer terurut
func calculateMedian(sortedData []int) int {
	n := len(sortedData)
	if n%2 == 1 {
		// Jika jumlah elemen ganjil, ambil elemen tengah
		return sortedData[n/2]
	} else {
		// Jika jumlah elemen genap, ambil rata-rata dari dua elemen tengah
		return (sortedData[(n/2)-1] + sortedData[n/2]) / 2
	}
}

```
#### Output
![image](https://github.com/user-attachments/assets/39b8ab41-e73c-410c-9ed8-351af3f53438)

#### Deskripsi program
Program ini dirancang untuk menghitung median dari sekumpulan angka yang dimasukkan oleh pengguna. Angka-angka tersebut akan disimpan dalam sebuah array, kemudian diurutkan setiap kali pengguna memasukkan angka 0. Setelah diurutkan, median dihitung berdasarkan posisi tengah data: jika jumlah elemen ganjil, median adalah elemen tengah, sedangkan jika jumlah elemen genap, median adalah rata-rata dari dua elemen tengah. Program akan terus menerima input angka hingga pengguna memasukkan angka -5313, yang menandakan program berhenti. Selain itu, jika data kosong saat median diminta, program akan memberikan notifikasi bahwa median tidak dapat dihitung. Output berupa median dan data terurut akan dicetak setiap kali angka 0 dimasukkan.

### 3. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:
![image](https://github.com/user-attachments/assets/1dff1608-33fc-4862-b3f1-caf3e14ec6f3)
Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.

Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima Judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesual rating yang diberikan pada masukan baris terakhir.

Lengkapi subprogram-subprogram dibawah ini, sesuai dengan I.S. dan F.S yang diberikan:
![image](https://github.com/user-attachments/assets/abf1e622-15a5-4e7b-99c2-7ca8c721a901)


#### Source Code :
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Buku struct {
	ID        int
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

type DaftarBuku []Buku

// Fungsi untuk membaca input angka (integer)
func bacaInt(reader *bufio.Reader, prompt string) int {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, _ := strconv.Atoi(input)
	return num
}

// Fungsi untuk membaca input string
func bacaString(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func DaftarkanBuku(pustaka *DaftarBuku, n int, reader *bufio.Reader) {
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("\nMasukkan data buku ke-%d:\n", i+1)

		buku.ID = bacaInt(reader, "ID: ")
		buku.Judul = bacaString(reader, "Judul: ")
		buku.Penulis = bacaString(reader, "Penulis: ")
		buku.Penerbit = bacaString(reader, "Penerbit: ")
		buku.Eksemplar = bacaInt(reader, "Eksemplar: ")
		buku.Tahun = bacaInt(reader, "Tahun: ")

		for {
			buku.Rating = bacaInt(reader, "Rating (0-10): ")
			if buku.Rating >= 0 && buku.Rating <= 10 {
				break
			}
			fmt.Println("Rating harus antara 0 dan 10. Coba lagi.")
		}

		*pustaka = append(*pustaka, buku)
	}
}

func CetakTerfavorit(pustaka DaftarBuku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}
	favorit := pustaka[0]
	for _, buku := range pustaka {
		if buku.Rating > favorit.Rating {
			favorit = buku
		}
	}
	fmt.Printf("\nBuku Terfavorit: %s oleh %s, %s (%d) - Rating: %d\n", favorit.Judul, favorit.Penulis, favorit.Penerbit, favorit.Tahun, favorit.Rating)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var pustaka DaftarBuku

	jumlahBuku := bacaInt(reader, "Masukkan jumlah buku: ")

	if jumlahBuku <= 0 {
		fmt.Println("Jumlah buku harus lebih dari 0.")
		return
	}

	DaftarkanBuku(&pustaka, jumlahBuku, reader)
	CetakTerfavorit(pustaka)
}

```
#### Output
![image](https://github.com/user-attachments/assets/4d444844-366d-4d81-be91-31a8c4887325)
![image](https://github.com/user-attachments/assets/9d944b04-7973-4f6e-bd66-2a4656f5b226)

#### Deskripsi program
Program ini merupakan aplikasi berbasis teks yang digunakan untuk mendata koleksi buku dan menampilkan buku dengan rating tertinggi. Setiap buku memiliki atribut berupa ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating. Program memvalidasi input pengguna, khususnya pada atribut rating yang harus berada dalam rentang 0 hingga 10. Setelah pengguna memasukkan jumlah buku dan detailnya satu per satu, program akan mencetak buku dengan rating tertinggi sebagai buku terfavorit. Input dibaca menggunakan bufio.Reader untuk menghindari konflik antara input angka dan string, serta memastikan data yang dimasukkan bersih dan sesuai format.
