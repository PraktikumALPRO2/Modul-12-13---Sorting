
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
  Mutia Rani Zahra Meilani
  <br>
  2311102182
  <br>
  S1IF1105
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

## <strong> DASAR TEORI </strong>

<span style="font-size:16px"><strong> ── PENGERTIAN SORTING</strong></span>
<br>
Pengurutan (sorting) merupakan proses menyusun elemen-elemen dari sebuah himpunan data ke dalam urutan tertentu, baik berdasarkan kriteria numerik, alfabet, atau kriteria lainnya. Tujuan utama dari pengurutan adalah untuk mempermudah pencarian, analisis, atau manipulasi data

<span style="font-size:16px"><strong> ── ALGORITMA SELECTION SORT</strong></span>
<br>

Selection sort adalah algoritma pengurutan yang bekerja dengan cara memilih elemen terkecil dari sisa data dan menukarnya dengan elemen pertama. Proses ini diulang untuk elemen berikutnya hingga seluruh data terurut. Pada setiap iterasi, algoritma ini mencari nilai terkecil dari bagian yang belum terurut dan menukarnya dengan elemen pertama dari bagian tersebut12.

**Langkah-langkah Algoritma Selection Sort**
- Mulai dengan indeks pertama (i = 0).
- Temukan elemen terkecil di antara elemen yang belum terurut (dari i hingga N-1).
- Tukar elemen terkecil tersebut dengan elemen di posisi i.
- Ulangi langkah 2 dan 3 untuk indeks berikutnya hingga seluruh elemen terurut.

Contoh Pseudocode:
```go
for i from 0 to N-1 do
    minIndex = i
    for j from i+1 to N do
        if array[j] < array[minIndex] then
            minIndex = j
    swap array[i] with array[minIndex]
```

<span style="font-size:16px"><strong> ── ALGORITMA INSERTION SORT</strong></span>
<br>

Insertion sort adalah algoritma yang mengurutkan data dengan cara membangun sublist yang sudah terurut secara bertahap. Setiap elemen baru akan disisipkan ke dalam posisi yang tepat dalam sublist tersebut. Konsep ini mirip dengan cara orang bermain kartu, di mana kartu disisipkan satu per satu ke dalam tumpukan yang sudah terurut12.

**Langkah-langkah Algoritma Insertion Sort**
- Mulai dari elemen kedua (indeks 1), anggap elemen pertama sudah terurut.
- Ambil elemen berikutnya dan bandingkan dengan elemen sebelumnya.
- Geser elemen-elemen yang lebih besar ke kanan untuk memberi ruang pada elemen baru.
- Sisipkan elemen baru pada posisi yang tepat.
- Ulangi langkah 2-4 hingga seluruh data terurut.

Contoh Pseudocode:
```go
for i from 1 to N-1 do
    key = array[i]
    j = i - 1
    while j >= 0 and array[j] > key do
        array[j + 1] = array[j]
        j = j - 1
    array[j + 1] = key
```

## <strong>  GUIDED </strong>

### ── Guided 1

#### Source Code

```go
package main

import (
	"fmt"
)

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

		fmt.Printf("Nomor rumah terurut untuk daerah %d : ", i+1)
		for _, num := range arr {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
```

#### Output
![image](https://github.com/user-attachments/assets/88bcacdb-823e-4f6b-8b6d-12c2a92c3328)

#### Deskripsi Program :
Program ini mengurutkan nomor rumah di beberapa daerah menggunakan algoritma selection sort. Fungsi `selectionSort` mencari elemen terkecil dalam array dan menukarnya dengan elemen pertama, kemudian berulang untuk elemen berikutnya. Program menerima input jumlah daerah (n), lalu untuk setiap daerah, meminta jumlah nomor rumah dan nomor-nomor tersebut. Setelah diurutkan, nomor rumah setiap daerah ditampilkan dalam urutan yang telah disortir. Proses ini memastikan nomor rumah di setiap daerah ditampilkan secara terorganisir.

### ── Guided 2

#### Source Code

```go
package main

import (
	"fmt"
)

func insertionSort(arr []int, n int) {
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

	fmt.Print("Masukkan data integer (Akhiri dengan bilangan negatif) : ")
	for {
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		arr = append(arr, num)
	}

	n := len(arr)

	insertionSort(arr, n)

	isConstant, difference := isConstantDifference(arr, n)

	fmt.Print("Array setelah diurutkan : ")
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	fmt.Print("Data berjarak")
	if isConstant {
		fmt.Printf(" %d", difference)
	}
	fmt.Println()
}
```

#### Output
![image](https://github.com/user-attachments/assets/d7b8f8f0-6ca7-45fe-ac39-2da445f99d07)

#### Deskripsi Program :
Program ini mengurutkan data integer menggunakan algoritma insertion sort dan memeriksa apakah elemen dalam array memiliki selisih konstan. Fungsi `insertionSort` menyisipkan elemen ke posisi yang tepat dalam urutan yang sudah disortir, sedangkan fungsi `isConstantDifference` memverifikasi konsistensi selisih antar elemen array yang sudah diurutkan. Pengguna memasukkan data integer hingga bilangan negatif, yang menandakan akhir input. Program kemudian menampilkan array yang telah diurutkan dan, jika berlaku, selisih konstan antar elemen.

## <strong>  UNGUIDED </strong>

### ── Unguided 1

#### Study Case

Berdasarkan yang dijelaskan Hercules, jika nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.

Format **Masukan** masih persis sama seperti sebelumnya.

**Keluaran** terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.

| No | Masukan                                      | Keluaran    |
|----|---------------------------------------------|-------------|
| 1  | 3                                           | 1 7 9 13 2 |
|    | 5 2 1 7 9 13                                | 15 27 39 75 133 189 |
|    | 6 189 15 27 39 75 133                       | 1 9 4      |
|    | 3 4 9 1                                     |             |

**Keterangan** : Terdapat 3 daerah dalam contoh masukan. Baris kedua berisi campuran bilangan ganjil dan genap. Baris berikutnya hanya berisi bilangan ganjil, dan baris terakhir hanya berisi bilangan genap.

**Petunjuk** :
Waktu pembacaan data, bilangan ganjil dan genap dipisahkan ke dalam dua array yang berbeda, untuk kemudian masing-masing diurutkan tersendiri.
Atau, tetap disimpan dalam satu array, diurutkan secara keseluruhan. Tetapi pada waktu pencetakan, mulai dengan mencetak semua nilai ganjil lebih dulu, kemudian setelah selesai cetaklah semua nilai genapnya.
#### Source Code

```go
package main

import "fmt"

func slice_data(arr []int) ([]int, []int) {
	var ganjil, genap []int
	for _, num := range arr {
		if num%2 == 0 {
			genap = append(genap, num)
		} else {
			ganjil = append(ganjil, num)
		}
	}
	return ganjil, genap
}

func asc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func desc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	var n int
	fmt.Print("Masukkan Jumlah Daerah (n) : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("\nMasukkan Jumlah Rumah Daerah %d : ", i+1)
		var m int
		fmt.Scan(&m)

		arr := make([]int, m)
		fmt.Printf("Masukkan %d Nomor Rumah Daerah %d : ", m, i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		ganjil, genap := slice_data(arr)

		desc(genap)
		asc(ganjil)

		fmt.Printf("Nomor Rumah Terurut [ Ganjil(+), Genap(-) ] Daerah %d : ", i+1)
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

#### Output
![image](https://github.com/user-attachments/assets/f59f0a29-da43-4369-be6d-ca43343d4fc2)

#### Deskripsi Program :
Program ini memisahkan nomor rumah di setiap daerah ke dalam dua kelompok berdasarkan paritasnya: ganjil dan genap. Fungsi `slice_data` memisahkan elemen array ke dalam dua array terpisah untuk ganjil dan genap. Fungsi `asc` dan `desc` mengurutkan array dalam urutan menaik dan menurun menggunakan algoritma bubble sort. Program meminta pengguna memasukkan jumlah daerah, nomor rumah di setiap daerah, lalu menampilkan nomor ganjil dalam urutan menaik diikuti nomor genap dalam urutan menurun untuk setiap daerah.

### ── Unguided 2

#### Study Case

Kompetisi pemrograman yang baru saja berlalu dilkuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim bertomba untuk menyelesaikan sebanyak mingkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?

_"Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkkan ke bawah."_

Buatlah program median yang mencetaK nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.

**Masukan** berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313.
Keluaran adalah median yang diminta, satu data per baris.

| No | Masukan                                      | Keluaran    |
|----|---------------------------------------------|-------------|
| 1  | 7 23 11 0 5 19 2 29 3 13 17 0 -5313         | 11          |
|    |                                             | 12          |

**Keterangan** :
Sampai bilangan 0 yang pertama, data terbaca adalah 7 23 11, setelah tersusun: 7 11 23, maka median saat itu adalah 11.
Sampai bilangan 0 yang kedua, data adalah 7 23 11 5 19 2 29 3 13 17, setelah diurutkan: 2 3 5 7 11 13 17 19 23 29, karena ada 10 data, genap, maka median adalah (11+13)/2=12.

**Petunjuk** :
Untuk setiap data bukan 0 (dan bukan marker -5313541) simpan ke dalam array, Dan setiap kali menemukan bilangan 0, urutkanlah data yang sudah tersimpan dengan menggunakan metode insertion sort dan ambil mediannya.
#### Source Code

```go
package main

import "fmt"

func sort(data []int) {
  n := len(data)
  for i := 0; i < n-1; i++ {
    minIdx := i
    for j := i + 1; j < n; j++ {
      if data[j] < data[minIdx] {
        minIdx = j
      }
    }
    // Tukar elemen
    data[i], data[minIdx] = data[minIdx], data[i]
  }
}

func median(data []int) int {
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

  fmt.Println("Masukkan Data (akhiri dengan -5313) ")
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
    sort(group)
    median := median(group)
    fmt.Printf("Median %d : %d\n", i+1, median)
  }
}
```

#### Output
![image](https://github.com/user-attachments/assets/f728b256-a761-408f-8f9f-0c7abb8059c8)

#### Deskripsi Program :
Program ini mengelompokkan data integer ke dalam beberapa grup, mengurutkannya menggunakan algoritma selection sort, dan menghitung median dari setiap grup. Fungsi `sort` mengurutkan elemen grup, sementara fungsi `median` menghitung nilai tengah grup berdasarkan panjangnya (genap atau ganjil). Pengguna memasukkan data integer yang dipisahkan dengan 0 untuk menandai grup baru, dan `-5313` untuk mengakhiri input. Program menampilkan median dari setiap grup yang telah diurutkan.

### ── Unguided 3

#### Study Case

Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:
```go
const nMax : integer = 7919
type Buku = <
id, judul, penulis, penerbit : string
eksemplar, tahun, rating : integer >
type DaftarBuku = array [ 1 .. nMax] of Buku
Pustaka : DaftarBuku
nPustaka: integer
```
**Masukan** terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.

**Keluaran** terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir.

Lengkapi subprogram-subprogram dibawah ini, sesuai dengan I.S. dan F.S yang diberikan.

```go
procedure DaftarkanBuku(in/out pustaka : DaftarBuku, n : integer)
{I.S. sejumlah n data buku telah siap para piranti masukan
F.S. n berisi sebuah nilai, dan pustaka berisi sejumlah n data buku}

procedure CetakTerfavorit(in pustaka : DaftarBuku, in n : integer)
{I.S. array pustaka berisi n buah data buku dan belum terurut
F.S. Tampilan data buku (judul, penulis, penerbit, tahun)
terfavorit, yaitu memiliki rating tertinggi}

procedure UrutBuku( in/out pustaka : DaftarBuku, n : integer )
{I.S. Array pustaka berisi n data buku
F.S. Array pustaka terurut menurun/mengecil terhadap rating.
Catatan: Gunakan metoda Insertion sort}

procedure Cetak5Terbaru( in pustaka : DaftarBuku, n integer )
{I.S. pustaka berisi n data buku yang sudah terurut menurut rating
F.S. Laporan 5 judul buku dengan rating tertinggi
Catatan: Isi pustaka mungkin saja kurang dari 5}

procedure CariBuku(in pustaka : DaftarBuku, n : integer, r : integer)
{I.S. pustaka berisi n data buku yang sudah terurut menurut rating
F.S. Laporan salah satu buku (judul, penulis, penerbit, tahun,
eksemplar, rating) dengan rating yang diberikan. Jika tidak ada buku
dengan rating yang ditanyakan, cukup tuliskan "Tidak ada buku dengan
rating seperti itu". Catatan: Gunakan pencarian biner/belah dua.}
```
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

func TambahBuku(pustaka *DaftarBuku) {
	var n int
	fmt.Print("Masukkan Jumlah Buku : ")
	fmt.Scan(&n)

	pustaka.NPustaka = n
	pustaka.Pustaka = make([]Buku, n)

	for i := 0; i < n; i++ {
		fmt.Printf("\nMasukkan Data Buku %d! \n", i+1)
		fmt.Print("ID : ")
		fmt.Scan(&pustaka.Pustaka[i].ID)
		fmt.Print("Judul : ")
		fmt.Scan(&pustaka.Pustaka[i].Judul)
		fmt.Print("Penulis : ")
		fmt.Scan(&pustaka.Pustaka[i].Penulis)
		fmt.Print("Penerbit : ")
		fmt.Scan(&pustaka.Pustaka[i].Penerbit)
		fmt.Print("Eksemplar : ")
		fmt.Scan(&pustaka.Pustaka[i].Eksemplar)
		fmt.Print("Tahun : ")
		fmt.Scan(&pustaka.Pustaka[i].Tahun)
		fmt.Print("Rating : ")
		fmt.Scan(&pustaka.Pustaka[i].Rating)
	}
}

func CetakBuku(pustaka DaftarBuku) {
	fmt.Println("\n= DATA BUKU = ")
	fmt.Printf("%-5s %-25s %-20s %-20s %-10s %-10s %-5s\n", "ID", "JUDUL", "PENULIS", "PENERBIT", "EKSEMPLAR", "TAHUN", "RATING")
	for _, buku := range pustaka.Pustaka {
		fmt.Printf("%-5s %-25s %-20s %-20s %-10d %-10d %-5d\n",
			buku.ID, buku.Judul, buku.Penulis, buku.Penerbit, buku.Eksemplar, buku.Tahun, buku.Rating)
	}
}

func UrutkanBuku(pustaka *DaftarBuku) {
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

func CetakBukuTeratas(pustaka DaftarBuku) {
	fmt.Println("\n= 5 BUKU TERATAS BERDASARKAN RATING = ")
	fmt.Printf("%-25s %-5s\n", "JUDUL", "RATING")
	for i := 0; i < 5 && i < pustaka.NPustaka; i++ {
		fmt.Printf("%-25s %-5d\n", pustaka.Pustaka[i].Judul, pustaka.Pustaka[i].Rating)
	}
}

func CariBuku(pustaka DaftarBuku) {
	var rating int
	fmt.Print("\nMasukan Rating :")
	fmt.Scan(&rating)

	left, right := 0, pustaka.NPustaka-1
	found := false
	for left <= right {
		mid := (left + right) / 2
		if pustaka.Pustaka[mid].Rating == rating {
			found = true
			buku := pustaka.Pustaka[mid]
			fmt.Printf("\n= BUKU DENGAN RATING %d = \n", rating)
			fmt.Printf("%-5s %-25s %-20s %-20s %-10s %-10s %-5s\n", "ID", "JUDUL", "PENULIS", "PENERBIT", "EKSEMPLAR", "TAHUN", "RATING")
			fmt.Printf("%-5s %-25s %-20s %-20s %-10d %-10d %-5d\n",
				buku.ID, buku.Judul, buku.Penulis, buku.Penerbit, buku.Eksemplar, buku.Tahun, buku.Rating)
			break
		} else if pustaka.Pustaka[mid].Rating < rating {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if !found {
		fmt.Println("\nTidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var pustaka DaftarBuku


	TambahBuku(&pustaka)

	UrutkanBuku(&pustaka)

	CetakBuku(pustaka)

	CetakBukuTeratas(pustaka)

	CariBuku(pustaka)
}
```

#### Output
![image](https://github.com/user-attachments/assets/70ec02b5-22a5-493e-875a-50eff2bf2a73)

#### Deskripsi Program :
Program ini mengelola data buku dengan menyediakan fitur untuk menambah, mencetak, mengurutkan, dan mencari data buku berdasarkan rating. Fungsi `TambahBuku` meminta pengguna memasukkan data buku, sementara fungsi `UrutkanBuku` mengurutkan daftar buku berdasarkan rating menggunakan insertion sort. Fungsi `CetakBuku` mencetak semua buku, sedangkan `CetakBukuTeratas` menampilkan 5 buku dengan rating tertinggi. Fungsi `CariBuku` menggunakan pencarian biner untuk menemukan buku dengan rating tertentu

## <strong> REFERENSI </strong>

#### [1] Ramadhani, A. (2015). Pengurutan (Sorting). Diakses dari https://media.neliti.com/media/publications/414701-none-931dccaa.pdf
#### [2] Yuliana, S. (2020). Struktur Data - Bab IV: Pengurutan (Sorting). Diakses dari https://yuliana.lecturer.pens.ac.id/Struktur%20Data%20C/Teori%20SD%20-%20pdf/Data%20Structure%20-%20Bab%206.pdf
#### [3] Munir, R. (2008). Algoritma Pengurutan Dalam Pemrograman. Diakses dari https://informatika.stei.itb.ac.id/~rinaldi.munir/Matdis/2008-2009/Makalah2008/Makalah0809-022.pdf)ac.id.
