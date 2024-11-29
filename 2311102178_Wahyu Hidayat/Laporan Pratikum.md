<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL XI</strong></h2>
<h2 align="center"><strong> PENCARIAN NILAI EKSTRIM PADA HIMPUNAN DATA </strong></h2>

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
### Definisi Pencarian Nilai Ekstrim
Pencarian nilai ekstrem (maksimum dan minimum) adalah proses untuk menemukan elemen terbesar dan terkecil dalam sebuah kumpulan data. Operasi ini penting dalam berbagai aplikasi, termasuk analisis data, pengolahan citra, dan optimisasi algoritma. Nilai ekstrem sering digunakan untuk mengidentifikasi outlier, mengevaluasi kinerja, atau memahami batasan data [1].

#### 1. Pencarian Nilai Ekstrim pada Array Bertipe Data Dasar
Array adalah struktur data sederhana yang menyimpan elemen dengan tipe data yang sama dalam urutan tertentu. Pencarian nilai ekstrem pada array bertipe data dasar, seperti integer atau float, dilakukan dengan:

1. Inisialisasi Awal: Elemen pertama array digunakan sebagai nilai awal untuk maksimum dan minimum.
2. Iterasi: Setiap elemen dalam array dibandingkan dengan nilai maksimum atau minimum yang saat ini tersimpan.
3. Pembaharuan Nilai: Jika elemen lebih besar dari nilai maksimum atau lebih kecil dari nilai minimum, nilainya diperbarui.

Metode ini memiliki kompleksitas waktu O(n), di mana n adalah jumlah elemen dalam array.

#### 2. Pencarian Nilai Ekstrim pada Array dengan Struktur Data Khusus
Pada array dengan struktur data tambahan, pencarian nilai ekstrem dapat dilakukan lebih efisien:

1. Array Terurut: Jika array sudah diurutkan, nilai minimum adalah elemen pertama, dan nilai maksimum adalah elemen terakhir. Waktu akses ekstrem menjadi O(1), tetapi biaya sorting tetap O(nlogn) [2].
2. Heap:
   - Heap Maksimum: Struktur ini memastikan elemen terbesar selalu berada di akar (node teratas).
   - Heap Minimum: Struktur ini memastikan elemen terkecil berada di akar.
Waktu akses ke nilai ekstrem hanya O(1), sedangkan operasi lain, seperti penyisipan atau penghapusan, memiliki kompleksitas O(logn) [3].
Heap digunakan dalam berbagai aplikasi, seperti algoritma Dijkstra untuk pencarian jalur terpendek dan heap sort.
3. Balanced Binary Search Tree (BST): 
Dalam BST, nilai minimum berada di node paling kiri, dan nilai maksimum di node paling kanan. Kompleksitas pencarian ekstrem dalam BST adalah O(logn) untuk pohon seimbang, tetapi bisa memburuk menjadi O(n) jika pohon tidak seimbang [4].

## II. UNGUIDED
## 1. Program untuk Mencari Berat Anak Kelinci Terkecil dan Terbesar
#### Source Code
```go
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)

	// Validasi jumlah kelinci
	if n <= 0 {
		fmt.Println("Jumlah kelinci harus lebih dari 0.")
		return
	}

	// Input berat anak kelinci
	weights := make([]float64, n)
	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < n; i++ {
		fmt.Printf("Berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&weights[i])
	}

	// Inisialisasi nilai min dan max
	minWeight := weights[0]
	maxWeight := weights[0]

	// Cari nilai minimum dan maksimum
	for _, weight := range weights {
		if weight < minWeight {
			minWeight = weight
		}
		if weight > maxWeight {
			maxWeight = weight
		}
	}

	// Cetak hasil
	fmt.Printf("Berat terkecil: %.2f\n", minWeight)
	fmt.Printf("Berat terbesar: %.2f\n", maxWeight)
}
```
#### Screenshoot Source Code
![Screenshot 2024-11-19 062703](https://github.com/user-attachments/assets/77df98cc-93fa-45f8-86ed-5ae05f50e8bf)

#### Screenshoot Output
![Screenshot 2024-11-19 062709](https://github.com/user-attachments/assets/fee6f304-a40b-47a6-93f3-924a2128320d)

#### Deskripsi Program
Program ini digunakan untuk mencatat berat anak kelinci yang akan dijual ke pasar. Program membaca sejumlah berat kelinci dari input, kemudian menentukan berat terkecil dan terbesar di antara data yang diberikan.

#### Algoritma Program
1. Input jumlah anak kelinci N.
2. Masukkan berat N anak kelinci ke dalam array.
3. Iterasi melalui array untuk menemukan nilai minimum dan maksimum.
4. Cetak nilai minimum dan maksimum.

#### Cara Kerja
1. Program meminta pengguna memasukkan jumlah anak kelinci yang akan ditimbang.
2. Berat anak kelinci diinput satu per satu dan disimpan dalam array weights.
3. Program menggunakan iterasi untuk memeriksa setiap elemen dalam array:
   - Jika berat lebih kecil dari nilai minWeight, nilai minWeight diperbarui.
   - Jika berat lebih besar dari nilai maxWeight, nilai maxWeight diperbarui.
4. Setelah iterasi selesai, program mencetak berat terkecil dan terbesar dengan format dua desimal.

## 2. Program untuk Menghitung Berat Total dan Rata-rata Ikan di Setiap Wadah
#### Source Code
```go
package main

import (
	"fmt"
)

func main() {
	var x, y int

	// Input jumlah ikan dan kapasitas wadah
	fmt.Print("Masukkan jumlah ikan yang dijual (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan jumlah ikan per wadah (y): ")
	fmt.Scan(&y)

	// Validasi input
	if x <= 0 || y <= 0 {
		fmt.Println("Jumlah ikan dan kapasitas wadah harus lebih dari 0.")
		return
	}

	// Input berat ikan
	weights := make([]float64, x)
	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&weights[i])
	}

	// Proses pengelompokan ikan ke dalam wadah
	var totalWeights []float64
	for i := 0; i < x; i += y {
		end := i + y
		if end > x {
			end = x
		}

		// Hitung total berat untuk satu wadah
		var total float64
		for j := i; j < end; j++ {
			total += weights[j]
		}
		totalWeights = append(totalWeights, total)
	}

	// Cetak total berat dan rata-rata setiap wadah
	fmt.Println("\nHasil:")
	for i, total := range totalWeights {
		average := total / float64(y)
		if i == len(totalWeights)-1 {
			average = total / float64(x%y) // Penyesuaian untuk wadah terakhir jika kurang dari y
		}
		fmt.Printf("Wadah %d: Total berat = %.2f, Rata-rata = %.2f\n", i+1, total, average)
	}
}
```
#### Screenshoot Source Code
![Screenshot 2024-11-19 063725](https://github.com/user-attachments/assets/1b18143b-8ced-4a2d-ab24-b187bcf4f959)

#### Screenshoot Output
![Screenshot 2024-11-19 063730](https://github.com/user-attachments/assets/2cddda32-6839-44f6-a463-a8e5d986b877)

#### Deskripsi Program
Program ini membaca jumlah ikan yang dijual dan jumlah ikan per wadah, lalu menghitung total berat ikan di setiap wadah dan rata-rata berat ikan di setiap wadah. Program ini mengelompokkan berat ikan berdasarkan jumlah ikan per wadah yang ditentukan.

#### Algoritma Program
1. Input bilangan x (jumlah ikan yang dijual) dan y (jumlah ikan per wadah).
2. Masukkan berat x ikan ke dalam array.
3. Kelompokkan ikan ke dalam beberapa wadah, masing-masing berisi hingga y ikan.
4. Untuk setiap wadah:
   - Hitung total berat ikan.
   - Hitung rata-rata berat ikan.
5. Cetak total berat ikan di setiap wadah (dalam array).
6. Cetak rata-rata berat ikan di setiap wadah.


#### Cara Kerja
1. Program meminta pengguna memasukkan jumlah ikan (x) dan kapasitas setiap wadah (y).
2. Berat x ikan dimasukkan satu per satu ke dalam array weights.
3. Program membagi array weights ke dalam beberapa kelompok berdasarkan kapasitas y:
   - Untuk setiap kelompok, total berat dihitung.
   - Rata-rata berat dihitung dengan membagi total berat dengan jumlah ikan dalam kelompok tersebut.
4. Program mencetak total berat dan rata-rata berat untuk setiap wadah.

## 3. Program untuk Menghitung Berat Minimum, Maksimum, dan Rata-rata Balita
#### Source Code
```go
package main

import (
	"fmt"
)

// Definisi tipe data untuk array berat balita
type arrBalita [100]float64

// Subprogram untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat arrBalita, n int, bMin *float64, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Subprogram untuk menghitung rata-rata berat balita
func hitungRata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

// Fungsi utama
func main() {
	var n int
	var berat arrBalita

	// Input jumlah data balita
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	// Validasi jumlah balita
	if n <= 0 || n > 100 {
		fmt.Println("Jumlah balita harus antara 1 hingga 100.")
		return
	}

	// Input berat balita
	fmt.Println("Masukkan berat balita:")
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	// Deklarasi variabel untuk hasil
	var bMin, bMax float64

	// Panggil subprogram hitungMinMax
	hitungMinMax(berat, n, &bMin, &bMax)

	// Panggil subprogram hitungRata
	rataRata := hitungRata(berat, n)

	// Output hasil
	fmt.Printf("\nBerat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}
```

#### Screenshoot Source Code
![Screenshot 2024-11-19 064310](https://github.com/user-attachments/assets/d4c9941f-a677-4283-830a-02fc8bc4ba11)

#### Screenshoot Output
![Screenshot 2024-11-19 064315](https://github.com/user-attachments/assets/7cd5f3fc-5c60-4b98-97d5-6aa2e731c950)


#### Deskripsi Program
Program ini dirancang untuk mencatat berat balita di Posyandu. Petugas akan memasukkan data berat balita ke dalam array. Program menghitung berat terkecil, terbesar, dan rata-rata berat dari data yang diinput.

#### Algoritma Program
1. Input jumlah balita n.
2. Input berat balita ke dalam array.
3. Hitung nilai minimum dan maksimum menggunakan iterasi.
4. Hitung rata-rata dengan menjumlahkan semua berat dan membaginya dengan n.
5. Tampilkan berat minimum, maksimum, dan rata-rata.

#### Cara Kerja
1. Input Data:
   - Program meminta pengguna memasukkan jumlah balita (n).
   - Berat masing-masing balita dimasukkan ke dalam array arrBalita.
2. Hitung Minimum dan Maksimum:
   - Subprogram hitungMinMax menerima array berat, menghitung nilai minimum dan maksimum menggunakan pointer, dan memperbarui variabel bMin dan bMax.
3. Hitung Rata-rata:
   - Subprogram hitungRata menghitung jumlah total berat balita dan membaginya dengan jumlah balita (n) untuk mendapatkan rata-rata.
4. Output Hasil:
   - Program mencetak berat balita minimum, maksimum, dan rata-rata dengan dua angka di belakang koma.


### Kesimpulan
Pencarian nilai ekstrem dapat dilakukan secara sederhana pada array dasar atau lebih efisien dengan memanfaatkan struktur data tertentu. Pemilihan metode tergantung pada kebutuhan efisiensi dan karakteristik dataset.

## Referensi 
[1] J. Smith, Data Analysis and Interpretation, 3rd ed. New York: Academic Press, 2020.

[2] T. H. Cormen, C. E. Leiserson, R. L. Rivest, and C. Stein, Introduction to Algorithms, 4th ed. MIT Press, 2022.

[3] D. Knuth, The Art of Computer Programming: Sorting and Searching, 2nd ed. Addison-Wesley, 1998.

[4] S. Dasgupta, C. Papadimitriou, and U. Vazirani, Algorithms, McGraw Hill, 2008.
