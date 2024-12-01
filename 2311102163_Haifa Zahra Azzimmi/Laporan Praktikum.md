<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL 12 & 13</strong></h2>
<h2 align="center"><strong> PENGURUTAN DATA </strong></h2> 

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Haifa Zahra Azzimmi / 2311102163<br>
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
3. [Guided](#guided)
4. [Unguided](#unguided)


## Dasar Teori

Algoritma Insertion Sort:
- Konsep Dasar: Algoritma ini bekerja dengan cara membangun larik yang diurutkan satu per satu. Pada setiap iterasi, elemen saat ini disisipkan ke dalam posisi yang sesuai dalam larik yang sudah diurutkan sebelumnya.
- Proses: Mulai dari elemen kedua, algoritma membandingkan elemen ini dengan elemen sebelumnya dan menukarnya jika perlu. Proses ini diulangi sampai semua elemen berada di tempat yang tepat.
- Kompleksitas Waktu: O(n^2) dalam kasus terburuk, karena setiap elemen mungkin perlu dibandingkan dengan semua elemen sebelumnya.
- Kelebihan: Mudah diimplementasikan dan efisien untuk larik kecil atau hampir terurut.
- Kekurangan: Kurang efisien untuk larik besar yang acak.

Algoritma Selection Sort:
- Konsep Dasar: Algoritma ini bekerja dengan cara membagi larik menjadi dua bagian: bagian yang sudah diurutkan dan bagian yang belum diurutkan. Pada setiap iterasi, algoritma menemukan elemen minimum dari bagian yang belum diurutkan dan menukarnya dengan elemen pertama dari bagian yang belum diurutkan.
- Proses: Algoritma mencari elemen terkecil dalam larik dan menukarnya dengan elemen pertama. Kemudian mencari elemen terkecil kedua dan menukarnya dengan elemen kedua, dan seterusnya.
- Kompleksitas Waktu: O(n^2) dalam semua kasus, karena setiap elemen dibandingkan dengan semua elemen lainnya.
- Kelebihan: Mudah diimplementasikan dan tidak memerlukan ruang tambahan.
- Kekurangan: Kurang efisien dibandingkan algoritma pengurutan lain untuk larik besar.

## Guided 

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumah kerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort.  

**Masukan** dimulai dengan sebuah integer 𝒏 (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi 𝒏 baris berikutnya selalu dimulai dengan sebuah integer 𝒎 (0 < m < 
1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat. 

**Keluaran** terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing masing daerah. 

### Source Code :
```go
// Haifa Zahra Azzimmi
// 2311102163
// S1 IF 11 5

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

### Output:
![image](https://github.com/user-attachments/assets/03f3b5ef-550d-4702-81d8-b70aedf92cc7)

### Full code Screenshot :
![code 10](https://github.com/user-attachments/assets/dca87620-c759-41e0-897c-7441c6fc1e9f)

### Deskripsi Program : 
Program ini mengurutkan nomor rumah kerabat di beberapa daerah menggunakan algoritma selection sort. Pertama, program meminta pengguna untuk memasukkan jumlah daerah kerabat. Setelah itu, untuk setiap daerah, program meminta jumlah nomor rumah kerabat dan berat masing-masing nomor rumah tersebut. Algoritma selection sort kemudian digunakan untuk mengurutkan nomor rumah di setiap daerah secara berurutan. Terakhir, program menampilkan nomor rumah yang sudah terurut untuk setiap daerah.

   
### 2. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya. 

**Masukan** terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array. 

**Keluaran** terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap". 

### Source Code :
```go
// Haifa Zahra Azzimmi
// 2311102163
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
![image](https://github.com/user-attachments/assets/3490294e-3824-4e97-9af1-a087935128a3)

### Full code Screenshot :
![code 11](https://github.com/user-attachments/assets/c5d2e95a-11fd-4509-adb1-40358e922eaa)

### Deskripsi Program : 
Program ini mengurutkan data integer yang dimasukkan oleh pengguna menggunakan metode Insertion Sort. Pengguna diminta memasukkan angka-angka positif dan menandai akhir input dengan angka negatif. Setelah semua data diperoleh, program mengurutkan elemen-elemen tersebut dengan menggunakan algoritma Insertion Sort, yang memposisikan setiap elemen pada tempat yang sesuai. Selain itu, program juga memeriksa apakah selisih antara elemen-elemen dalam array yang sudah diurutkan adalah konstan. Akhirnya, program menampilkan array yang telah diurutkan dan menyatakan apakah data tersebut memiliki selisih yang tetap atau tidak.


## Unguided 

### 1. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.   

**Masukan** masih persis sama seperti sebelumnya (guided 1). 

**Keluaran** terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk 
nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah. 

### Source Code :
```go
// Haifa Zahra Azzimmi
// 2311102163
// S1 IF 11 5

package main

import (
	"fmt"
	"sort"
)

func main() {
	var jumlahDaerah int
	fmt.Print("Masukkan jumlah daerah kerabat: ")
	fmt.Scan(&jumlahDaerah)

	for i := 1; i <= jumlahDaerah; i++ {
		var jumlahRumah int
		fmt.Printf("\nMasukkan jumlah rumah di daerah %d: ", i)
		fmt.Scan(&jumlahRumah)

		nomorRumah := make([]int, jumlahRumah)
		fmt.Printf("Masukkan %d nomor rumah: ", jumlahRumah)
		for j := 0; j < jumlahRumah; j++ {
			fmt.Scan(&nomorRumah[j])
		}

		var ganjil []int
		var genap []int
		for _, nomor := range nomorRumah {
			if nomor%2 == 0 {
				genap = append(genap, nomor)
			} else {
				ganjil = append(ganjil, nomor)
			}
		}

		sort.Ints(ganjil)
		sort.Sort(sort.Reverse(sort.IntSlice(genap)))

		// Tampilkan hasil
		fmt.Printf("\nNomor rumah terurut untuk daerah %d:\n", i)
		for _, nomor := range ganjil {
			fmt.Printf("%d ", nomor)
		}
		for _, nomor := range genap {
			fmt.Printf("%d ", nomor)
		}
		fmt.Println()
	}
}

```

### Output:
![image](https://github.com/user-attachments/assets/8e4b1521-6b79-4de8-a13a-13d0647517cf)

### Full code Screenshot :
![code 12](https://github.com/user-attachments/assets/173d9ef2-3b9f-49ff-b361-22e5b52ae267)

### Deskripsi Program : 
Program ini bertujuan untuk mengurutkan nomor rumah di beberapa daerah berdasarkan keparitasan (ganjil atau genap). Pertama, pengguna diminta untuk memasukkan jumlah daerah kerabat dan jumlah rumah di setiap daerah. Setelah itu, nomor rumah yang dimasukkan akan dipisahkan menjadi dua kelompok: ganjil dan genap. Kelompok ganjil diurutkan secara menaik, sementara kelompok genap diurutkan secara menurun. Pada akhirnya, program menampilkan nomor rumah yang sudah diurutkan untuk setiap daerah, dengan nomor ganjil muncul lebih dulu diikuti oleh nomor genap.

### 2.Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya? "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0. 

**Masukan** berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 
data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak 
termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313. 

**Keluaran** adalah median yang diminta, satu data per baris.

### Source Code :
```go
// Haifa Zahra Azzimmi
// 2311102163
// S1 IF 11 5

package main

import "fmt"

func sortSlice(slice []int) {
    panjangSlice := len(slice)
    for i := 0; i < panjangSlice-1; i++ {
        indeksMinimum := i
        for j := i + 1; j < panjangSlice; j++ {
            if slice[j] < slice[indeksMinimum] {
                indeksMinimum = j
            }
        }
        slice[i], slice[indeksMinimum] = slice[indeksMinimum], slice[i]
    }
}

func cariMedian(slice []int) int {
    jumlahElemen := len(slice)
    if jumlahElemen%2 == 1 {
        return slice[jumlahElemen/2]
    }
    return (slice[(jumlahElemen/2)-1] + slice[jumlahElemen/2]) / 2
}

func main() {
    var daftarAngka []int
    fmt.Print("Masukkan sejumlah angka (akhiri dengan -5313): ")

    for {
        var angka int
        _, err := fmt.Scan(&angka)

        if angka == -5313 {
            break
        }

        if err != nil {
            fmt.Println("Input tidak valid, coba lagi.")
            break
        }

        if angka == 0 {
            sortSlice(daftarAngka)
            fmt.Printf("Nilai median: %d\n", cariMedian(daftarAngka))
        } else {
            // Tambahkan angka ke slice
            daftarAngka = append(daftarAngka, angka)
        }
    }
}

```

### Output:
![image](https://github.com/user-attachments/assets/06e90ff8-17a9-4fbf-b577-5c005cd3842a)

### Full code Screenshot :
![code 13](https://github.com/user-attachments/assets/cf123bf6-c547-4b6d-a554-e7b92b1007de)

### Deskripsi Program : 
Program ini mengurutkan angka yang dimasukkan oleh pengguna dan menghitung median. Pengguna memasukkan angka sampai mereka mengetikkan -5313 untuk berhenti. Setiap angka ditambahkan ke dalam daftar. Jika angka 0 dimasukkan, program mengurutkan daftar dan menghitung median. Program ini kemudian menampilkan median dari daftar yang diurutkan. Program ini memverifikasi input dan menampilkan pesan kesalahan jika ada input yang tidak valid.


### 3.Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:

![image](https://github.com/user-attachments/assets/24629392-d82d-4936-810b-37a7a0b03b41)

**Masukan** terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan 
banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya 
adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan 
bulat yang menyatakan rating buku yang akan dicari. 

**Keluaran** terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir.


### Source Code :
```go
// Haifa Zahra Azzimmi
// 2311102163
// S1 IF 11 5

package main

import "fmt"

const maxBooks = 7919

// Struktur Buku untuk menyimpan informasi buku
type Buku struct {
	kode, judul, penulis, penerbit string
	eksemplar, tahun, rating       int
}

// Array KoleksiBuku untuk menyimpan daftar buku
type KoleksiBuku [maxBooks]Buku

// Fungsi untuk menambahkan buku ke dalam koleksi
func tambahBuku(koleksi *KoleksiBuku, jumlah *int) {
	fmt.Print("Masukkan jumlah buku yang ingin didaftarkan: ")
	fmt.Scan(jumlah)

	for i := 0; i < *jumlah; i++ {
		fmt.Printf("\nMasukkan data buku ke-%d (kode, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scan(&koleksi[i].kode, &koleksi[i].judul, &koleksi[i].penulis, &koleksi[i].penerbit,
			&koleksi[i].eksemplar, &koleksi[i].tahun, &koleksi[i].rating)
	}
}

// Fungsi untuk menampilkan buku dengan rating tertinggi
func bukuFavorit(koleksi KoleksiBuku, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Tidak ada buku yang terdaftar.")
		return
	}

	favorit := koleksi[0]
	for i := 1; i < jumlah; i++ {
		if koleksi[i].rating > favorit.rating {
			favorit = koleksi[i]
		}
	}

	fmt.Println("\nBuku Favorit:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
		favorit.judul, favorit.penulis, favorit.penerbit, favorit.tahun, favorit.rating)
}

// Fungsi untuk mengurutkan buku berdasarkan rating secara menurun
func urutkanBuku(koleksi *KoleksiBuku, jumlah int) {
	for i := 1; i < jumlah; i++ {
		buku := koleksi[i]
		j := i - 1

		for j >= 0 && koleksi[j].rating < buku.rating {
			koleksi[j+1] = koleksi[j]
			j--
		}
		koleksi[j+1] = buku
	}
}

// Fungsi untuk menampilkan 5 buku dengan rating tertinggi
func tampilkanTop5(koleksi KoleksiBuku, jumlah int) {
	fmt.Println("\n5 Buku Dengan Rating Tertinggi:")
	batas := 5
	if jumlah < 5 {
		batas = jumlah
	}

	for i := 0; i < batas; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, koleksi[i].judul, koleksi[i].rating)
	}
}

// Fungsi untuk mencari buku berdasarkan rating
func cariBukuByRating(koleksi KoleksiBuku, jumlah, rating int) {
	ditemukan := false
	fmt.Printf("\nMencari Buku dengan Rating %d:\n", rating)

	for i := 0; i < jumlah; i++ {
		if koleksi[i].rating == rating {
			ditemukan = true
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
				koleksi[i].judul, koleksi[i].penulis, koleksi[i].penerbit,
				koleksi[i].tahun, koleksi[i].eksemplar, koleksi[i].rating)
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var koleksi KoleksiBuku
	var jumlahBuku, ratingPencarian int

	// Menambahkan buku ke dalam koleksi
	tambahBuku(&koleksi, &jumlahBuku)

	// Menampilkan buku dengan rating tertinggi
	bukuFavorit(koleksi, jumlahBuku)

	// Mengurutkan buku berdasarkan rating dan menampilkan hasilnya
	urutkanBuku(&koleksi, jumlahBuku)

	// Menampilkan 5 buku dengan rating tertinggi
	tampilkanTop5(koleksi, jumlahBuku)

	// Mencari buku berdasarkan rating
	fmt.Print("\nMasukkan rating yang ingin dicari: ")
	fmt.Scan(&ratingPencarian)
	cariBukuByRating(koleksi, jumlahBuku, ratingPencarian)
}

```
### Output:
![image](https://github.com/user-attachments/assets/f9d93673-0303-4aa8-9794-d36bb9b50055)

### Full code Screenshot :
![code 14](https://github.com/user-attachments/assets/eb9cd51d-6d74-4988-9aeb-194ec9927182)

### Deskripsi Program : 
Program ini digunakan untuk mengelola koleksi buku, termasuk menambahkan buku, menemukan buku dengan rating tertinggi, mengurutkan buku berdasarkan rating, menampilkan lima buku dengan rating tertinggi, dan mencari buku berdasarkan rating tertentu. Pengguna diminta memasukkan data buku yang terdiri dari kode, judul, penulis, penerbit, eksemplar, tahun, dan rating. Buku dengan rating tertinggi ditampilkan sebagai buku favorit. Selain itu, program mengurutkan buku berdasarkan rating secara menurun dan menampilkan lima buku dengan rating tertinggi. Program juga memungkinkan pengguna untuk mencari buku dengan rating tertentu dan menampilkan informasi tentang buku-buku tersebut jika ditemukan.
