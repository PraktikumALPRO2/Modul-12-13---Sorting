<p align="center">
  <strong>LAPORAN PRAKTIKUM</strong>
  <br>
  <strong>ALGORITMA DAN PEMROGRAMAN 2</strong>
  <br>
</p>

<br>

<p align="center">
  <strong>MODUL XII</strong>
  <br>
  <strong>PENGURUTAN DATA</strong>
  <br>
</p>

<br>

<p align="center">
  <img src="https://github.com/user-attachments/assets/296eb3c2-bd6b-401f-8256-3661150ec39e" alt="Logo" width="200"/>
</p>

<br>

<p align="center">
  <strong>Disusun Oleh :</strong>
  <br>
  Abdee Alamsyah Noer Siyam
  <br>
  2311102247
  <br>
  S1-IF11-05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu :</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong>
  <br>
  <strong>FAKULTAS INFORMATIKA</strong>
  <br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong>
  <br>
  <strong>2024</strong>
</p>


## <strong> Dasar Teori </strong>

<strong><h3>SORTING DATA</h3></strong>

Sorting data adalah proses pengurutan elemen dalam sebuah array atau list sehingga elemen-elemen tersebut terorganisir dalam urutan tertentu, baik secara ascending (menaik) maupun descending (menurun). Dua algoritma yang umum digunakan untuk sorting adalah Selection Sort dan Insertion Sort. Masing-masing algoritma memiliki prinsip kerja dan efisiensi yang berbeda.

<strong><h3>SELECTION SORT</h3></strong>

Selection Sort adalah algoritma pengurutan yang bekerja dengan cara memilih elemen terkecil dari bagian yang belum terurut dan menukarnya dengan elemen pertama dari bagian tersebut. Proses ini diulang untuk setiap elemen dalam array hingga seluruh array terurut.

<strong><h4>LANGKAH-LANGKAH</h4></strong>

Berikut adalah langkah-langkah algoritma selection sort

1. Mulai dari indeks pertama, anggap elemen pada indeks tersebut sebagai elemen terkecil.
2. Bandingkan elemen terkecil dengan elemen lainnya di bagian yang belum terurut.
3. Jika ditemukan elemen yang lebih kecil, perbarui indeks elemen terkecil.
4. Setelah menyelesaikan pencarian, tukar elemen terkecil dengan elemen pada posisi awal.
5. Ulangi langkah 1 hingga 4 untuk seluruh array.

<strong><h4>IMPLEMENTASI DALAM BAHASA GO</h4></strong>

Berikut adalah contoh implementasi algoritma Selection Sort dalam bahasa Go:

```go
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
```

<strong><h3>INSERTION SORT</h3></strong>

Insertion Sort bekerja dengan cara membangun urutan yang terurut satu per satu. Pada setiap iterasi, algoritma mengambil satu elemen dari bagian yang belum terurut dan menyisipkannya ke posisi yang tepat di antara elemen-elemen yang sudah terurut.ut.

<strong><h4>LANGKAH-LANGKAH</h4></strong>

Berikut adalah langkah-langkah algoritma selection sort

1. Anggap bahwa elemen pertama sudah terurut.
2. Ambil elemen berikutnya dan simpan sebagai kunci (key).
3. Bandingkan kunci dengan elemen-elemen di bagian yang sudah terurut dari belakang ke depan.
4. Geser elemen-elemen yang lebih besar dari kunci satu posisi ke kanan untuk memberikan ruang bagi kunci.
5. Sisipkan kunci ke posisi yang tepat.
6. Ulangi langkah 2 hingga seluruh array terurut.

<strong><h4>IMPLEMENTASI DALAM BAHASA GO</h4></strong>

Berikut adalah contoh implementasi algoritma Insertion Sort dalam bahasa Go:

```go
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
```


## <strong> GUIDED </strong>

### <h2>GUIDED 1</h2>

#### Source Code

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
		arr[i] = arr[minIdx]
		arr[minIdx] = arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	rumah := make([]string, n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		houses := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}

		selectionSort(houses)
		result := ""
		for j, angka := range houses {
			if j > 0 {
				result += " "
			}
			result += fmt.Sprint(angka)
		}

		rumah[i] = result
	}
	fmt.Println("\nHasil Setelah diurutkan : ")
	for _, line := range rumah {
		fmt.Println(line)
	}
}
```
#### Screenshot

![image](https://github.com/user-attachments/assets/399a41e8-4881-48c7-b396-97b187b5bac5)

#### Deksripsi Program

Program ini mengimplementasikan algoritma Selection Sort untuk mengurutkan data dalam urutan naik. Program menerima input berupa beberapa grup angka, di mana setiap grup mewakili daftar angka dari rumah yang berbeda. Setiap grup angka diurutkan menggunakan Selection Sort, kemudian hasil pengurutan untuk setiap grup dicetak secara terpisah. Tujuan utama program adalah mengurutkan setiap daftar angka secara individual dan menampilkan hasil akhir dalam bentuk daftar yang sudah terurut.

### <h2>GUIDED 2</h2>

#### Source Code

```go
package main

import "fmt"

func insertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}


func berjarakTetap(data []int) (bool, int) {
	if len(data) < 2 {
		return true, 0
	}

	spacing := data[1] - data[0]
	for i := 1; i < len(data)-1; i++ {
		if data[i+1]-data[i] != spacing {
			return false, 0
		}
	}

	return true, spacing
}

func main() {
	var input int
	var data []int

	fmt.Println("Masukkan data (akhiri dengan bilangan negatif):")
	for {
		fmt.Scan(&input)
		if input < 0 { 
			break
		}
		data = append(data, input)
	}

	sortedData := make([]int, len(data))
	copy(sortedData, data)

	insertionSort(sortedData)

	consistent, spacing := berjarakTetap(sortedData)
	
	fmt.Println("Data setelah diurutkan:", sortedData)
	if consistent {
		fmt.Printf("Data berjarak %d\n", spacing)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}

```
#### Screenshot

![image](https://github.com/user-attachments/assets/4face4b7-7014-4f93-824b-0f56bdcc7bbe)

#### Deksripsi Program

Program ini melakukan dua operasi utama pada data yang diinputkan oleh pengguna. Pertama, data diurutkan menggunakan algoritma Insertion Sort untuk menghasilkan daftar angka dalam urutan menaik. Kedua, program memeriksa apakah data yang telah diurutkan memiliki pola jarak tetap antara elemen-elemen berturutannya. Jika pola jarak tetap ditemukan, program akan mencetak nilai jarak tersebut; jika tidak, akan ditampilkan pesan bahwa jaraknya tidak tetap. Input data dihentikan ketika pengguna memasukkan angka negatif, dan hasil akhir berupa data yang terurut serta informasi tentang pola jaraknya dicetak ke layar. Program ini berguna untuk menganalisis keteraturan dalam data numerik.

## <strong> UNGUIDED </strong>

### <h2>UNGUIDED 1</h2>

#### Question
**Latar Belakang**
Diketahui bahwa Hercules tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyebrang jalan sesedikit mungkin. Karena rumah sihir jalan ganjil dan sihir jalan genap, maka buatlah program `kerabat dekat` yang dapat membantu menghitung rumah-rumah mana yang tidak teratur membesar dan kemudian memplotkan rumah dengan nomor genap kembali terpisah.

**Masukan**
Masukan masih persis seperti sebelumnya.

**Keluaran**
Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabat yang terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.

**Contoh**
| No  | Masukan | Keluaran |
|-----|---------|----------|
| 1   | 3  5  12  15  7  18  19  189 | 3  5  7  12  15  18  19  189 |

#### Source Code

```go
package main

import "fmt"


func selectionSort(arr []int, asc bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		targetIdx := i
		for j := i + 1; j < n; j++ {
			if (asc && arr[j] < arr[targetIdx]) || (!asc && arr[j] > arr[targetIdx]) {
				targetIdx = j
			}
		}
		arr[i], arr[targetIdx] = arr[targetIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	rumah := make([]string, n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		houses := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}

		var ganjil, genap []int
		for _, house := range houses {
			if house%2 == 1 {
				ganjil = append(ganjil, house)
			} else {
				genap = append(genap, house)
			}
		}

		selectionSort(ganjil, true)  
		selectionSort(genap, false)

		result := ""
		for j, angka := range ganjil {
			if j > 0 {
				result += " "
			}
			result += fmt.Sprint(angka)
		}
		for j, angka := range genap {
			if j > 0 || len(ganjil) > 0 {
				result += " "
			}
			result += fmt.Sprint(angka)
		}

		rumah[i] = result
	}

	fmt.Println("\nHasil Setelah diurutkan:")
	for _, line := range rumah {
		fmt.Println(line)
	}
}

```
#### Screenshot

![image](https://github.com/user-attachments/assets/3a887e6e-e064-4b14-83bd-684516578006)

#### Deksripsi Program

Program ini merupakan versi pengembangan dari guided pertama dengan tambahan fitur pengelompokan angka menjadi ganjil dan genap. Angka-angka ganjil diurutkan secara ascending (menaik), sedangkan angka-angka genap diurutkan secara descending (menurun) menggunakan algoritma Selection Sort. Setelah pengelompokan dan pengurutan, hasil untuk setiap grup angka digabung kembali dan dicetak. Program ini memungkinkan pengguna untuk melihat daftar angka yang sudah terorganisir berdasarkan prioritas ganjil dan genap, dengan pola urutan yang berbeda.

### <h2>UNGUIDED 2</h2>

Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berhadapan dengan satu masalah. Problem tersebut terdiri dari dua bagian. Apa sih problemnya?

"Median adalah nilai tengah dari suatu koleksi data angka. Jika jumlah data genap, maka median adalah rata-rata dari dua nilai tengah. Untuk problem ini, semuda data merupakan bilangan bulat dan urutan dari rentang nilai tengah dibutuhkan ke bawah."

**Masukan**
Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data. Terdapat dua jenis masukan, yaitu nomor genap dan nomor ganjil.

**Contoh**
| No  | Masukan | Keluaran |
|-----|---------|----------|
| 1   |  23  11  15  -17  -5  -5313  | 11, 12  |

**Keterangan**
Sampai bilangan 0 yang pertama, data terbaca adalah 7 23 11, setelah tersusun: 7 11 23, maka median sampai bilangan 0 adalah 11.

#### Source Code

```go
package main

import "fmt"

func selectionSort(data []int) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data[j] < data[minIdx] {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
}

func calculateMedian(data []int) int {
	n := len(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2 
	}
	return data[n/2]
}

func main() {
	var input int
	var data []int
	var groups [][]int

	fmt.Println("Masukkan data (akhiri dengan -5313):")
	for {
		fmt.Scan(&input)
		if input == -5313 {
			break
		}
		if input == 0 {
			groups = append(groups, append([]int{}, data...)) 
		} else {
			data = append(data, input)
		}
	}

	for i, group := range groups {
		selectionSort(group)
		median := calculateMedian(group)
		fmt.Printf("Median kelompok %d: %d\n", i+1, median)
	}
}
```

#### Output

![image](https://github.com/user-attachments/assets/7e883459-0a35-4b7c-b51e-574507e1aadb)

Deskripsi Program :

Program ini digunakan untuk memproses beberapa kelompok data numerik yang diinputkan oleh pengguna, menghitung median dari setiap kelompok, dan mencetak hasilnya. Pengguna memasukkan data angka, di mana angka **0** digunakan untuk memisahkan antar kelompok, dan angka **-5313** digunakan untuk mengakhiri input. Program mengurutkan setiap kelompok data menggunakan algoritma **Selection Sort**, lalu menghitung mediannya (nilai tengah) berdasarkan data yang telah terurut. Hasil median untuk setiap kelompok kemudian ditampilkan secara terpisah. Program ini bermanfaat untuk menganalisis statistik sederhana dari kumpulan data tersegmentasi.

### <h2>UNGUIDED 3</h2>

#### Question
Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misinya terdefinisi struct dan array seperti berikut ini:

```go
const nMax : integer = 7919
type Buku = < 
  id, judul, penulis, penerbit : string
  eksamplr, tahun, rating : integer >

type DaftarBuku = array [ 1 .. nMax] of Buku
DaftarBuku : DaftarBuku
nBuku : integer
```

**Masukan** terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masing adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan banyak data buku yang dicari.

**Keluaran** terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah data buku dengan judul yang degan rating tertinggi, sedangnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir.

#### Source Code

```go
package main

import (
	"fmt"
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

type DaftarBuku struct {
	Pustaka  []Buku
	NPustaka int
}

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	pustaka.NPustaka = n
	pustaka.Pustaka = make([]Buku, n)
	for i := 0; i < n; i++ {
		fmt.Println("Masukkan ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating:")
		fmt.Scanln(&pustaka.Pustaka[i].ID, &pustaka.Pustaka[i].Judul, &pustaka.Pustaka[i].Penulis,
			&pustaka.Pustaka[i].Penerbit, &pustaka.Pustaka[i].Eksemplar, &pustaka.Pustaka[i].Tahun, &pustaka.Pustaka[i].Rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku) {
	fmt.Println("\nData Buku:")
	for _, buku := range pustaka.Pustaka {
		fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
			buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Rating)
	}
}

func UrutBuku(pustaka *DaftarBuku) {
	n := pustaka.NPustaka
	for i := 1; i < n; i++ {
		key := pustaka.Pustaka[i]
		j := i - 1
		for j >= 0 && pustaka.Pustaka[j].Rating < key.Rating {
			pustaka.Pustaka[j+1] = pustaka.Pustaka[j]
			j--
		}
		pustaka.Pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku) {
	fmt.Println("\n5 Buku dengan Rating Tertinggi:")
	for i := 0; i < 5 && i < pustaka.NPustaka; i++ {
		fmt.Printf("Judul: %s, Rating: %d\n", pustaka.Pustaka[i].Judul, pustaka.Pustaka[i].Rating)
	}
}

func CariBuku(pustaka DaftarBuku, r int) {
	left, right := 0, pustaka.NPustaka-1
	found := false
	for left <= right {
		mid := (left + right) / 2
		if pustaka.Pustaka[mid].Rating == r {
			found = true
			buku := pustaka.Pustaka[mid]
			fmt.Printf("\nBuku ditemukan:\nID: %s, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
				buku.ID, buku.Judul, buku.Penulis, buku.Penerbit, buku.Eksemplar, buku.Tahun, buku.Rating)
			break
		} else if pustaka.Pustaka[mid].Rating < r {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu.")
	}
}

func main() {
	var pustaka DaftarBuku
	var n int

	fmt.Println("Masukkan jumlah buku:")
	fmt.Scanln(&n)

	DaftarkanBuku(&pustaka, n)

	CetakTerfavorit(pustaka)

	UrutBuku(&pustaka)

	Cetak5Terbaru(pustaka)

	var rating int
	fmt.Println("\nMasukkan rating buku yang ingin dicari:")
	fmt.Scanln(&rating)
	CariBuku(pustaka, rating)
}

```

#### Output

![image](https://github.com/user-attachments/assets/e2192227-c62d-46f5-a53e-23e626d9f55e)

Deskripsi Program :
Program ini merupakan sistem pengelolaan data perpustakaan sederhana yang memungkinkan pengguna untuk mendaftarkan buku, mengurutkan buku berdasarkan rating, menampilkan daftar buku, dan mencari buku dengan rating tertentu. Pengguna dapat memasukkan data buku berupa ID, judul, penulis, penerbit, jumlah eksemplar, tahun penerbitan, dan rating. Data buku yang telah terdaftar dapat ditampilkan kembali untuk dilihat. Selanjutnya, buku diurutkan berdasarkan rating dalam urutan menurun menggunakan algoritma Insertion Sort, sehingga mempermudah dalam menemukan buku dengan rating terbaik. Setelah proses pengurutan, program akan menampilkan hingga 5 buku dengan rating tertinggi. Selain itu, pengguna juga dapat mencari buku berdasarkan rating tertentu menggunakan metode pencarian biner, dan jika buku ditemukan, informasi lengkap buku tersebut akan ditampilkan. Program ini dirancang untuk membantu perpustakaan mengorganisasi data buku dan memberikan kemudahan dalam pencarian buku favorit berdasarkan rating.

## <strong> Kesimpulan </strong>

## Kesimpulan

Baik Selection Sort maupun Insertion Sort memiliki kelebihan dan kekurangan masing-masing tergantung pada kondisi data yang akan diurutkan. Selection Sort lebih sederhana namun kurang efisien dibandingkan Insertion Sort pada data kecil atau hampir terurut. Sebaliknya, Insertion Sort lebih fleksibel dan efisien dalam banyak situasi praktis.

## <strong> Referensi </strong>

#### [1] Retnoningsih, Endang. (2018).Algoritma Pengurutan Data (Sorting) Dengan Metode Insertion Sort dan Selection Sort. Diakses dari https://download.garuda.kemdikbud.go.id/article.php?article=909504&title=Algoritma+Pengurutan+Data+Sorting+Dengan+Metode+Insertion+Sort+dan+Selection+Sort&val=11043

#### [2] Christopher, Finn. (2022). Algoritma Selection Sort di Python. Diakses dari https://binus.ac.id/bandung/2023/10/algoritma-selection-sort-di-python/

#### [3] Chandak, Gaurav. (2022). Sorting Algorithms (Bubble Sort, Insertion Sort, Selection Sort). Diakses dari https://workat.tech/problem-solving/tutorial/sorting-algorithms-bubble-insertion-selection-sort-veubp86w3e1r
