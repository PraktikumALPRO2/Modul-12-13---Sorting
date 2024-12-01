# <h2 align="center">LAPORAN PRAKTIKUM</h2>
# <h2 align="center">ALGORITMA DAN PEMROGRAMAN 2</h2>
# <h2 align="center">MODUL 12-13</h2>
# <h2 align="center">SORTING</h2>
<p align="center">
    <img src="https://github.com/user-attachments/assets/3ccfed0b-72d1-4349-ac08-c4dce379c827" alt="Gambar">
</p>
 <h3  align="center" >Disusun Oleh : </h3>
<p align="center">Syamsul Adam - 2311102144</p>
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
## Output:![Screenshot 2024-12-01 224400](https://github.com/user-attachments/assets/abc428e0-25db-4ccd-86a1-e81471b07c39)

Program ini berfungsi untuk mengurutkan sejumlah array (daftar angka) dalam urutan menaik menggunakan algoritma Selection Sort. 
Program menerima input berupa jumlah array dan elemen-elemen setiap array, kemudian mencetak hasil pengurutan masing-masing array.



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
## Output:![Screenshot 2024-12-01 224057](https://github.com/user-attachments/assets/51b24bb2-4a15-4efd-ab75-d7cd2a8359f0)


Program ini dirancang untuk menerima sejumlah bilangan bulat positif, mengurutkan data dalam urutan menaik menggunakan algoritma Sort Insertion, dan kemudian memeriksa apakah jarak antar elemen tetap. Ketika pengguna memasukkan bilangan negatif, program berakhir.

## III. UNGUIDED 

### 1. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya di ujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.<br/> Format Masukan: masih persis sama seperti sebelumnya.<br/> Keluaran: terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.<br/> ![image](https://github.com/user-attachments/assets/950d0183-c431-448e-8ce8-b03a3651ca03) <br/> Keterangan: Terdapat 3 daerah dalam contoh masukan. Baris kedua berisi campuran bilangan ganjil dan genap. Baris berikutnya hanya berisi bilangan ganjil, dan baris terakhir hanya berisi bilangan genap.<br/> 
### Petunjuk:
### - Waktu pembacaan data, bilangan ganjil dan genap dipisahkan ke dalam dua array yang berbeda, untuk kemudian masing-masing diurutkan tersendiri.
### - Atau, tetap disimpan dalam satu array, diurutkan secara keseluruhan. Tetapi pada waktu pencetakan, mulai dengan mencetak semua nilai ganjil lebih dulu, kemudian setelah selesai cetaklah semua nilai genapnya.

```go
package main

import (
	"fmt"
	"sort"
)

// Fungsi untuk memisahkan dan mengurutkan nomor ganjil dan genap
func prosesNomorRumah(nomorRumah []int) ([]int, []int) {
	// Slice untuk ganjil dan genap
	ganjil := []int{}
	genap := []int{}

	// Memisahkan ganjil dan genap
	for _, nomor := range nomorRumah {
		if nomor%2 == 0 {
			genap = append(genap, nomor)
		} else {
			ganjil = append(ganjil, nomor)
		}
	}

	// Mengurutkan nomor rumah
	sort.Sort(sort.Reverse(sort.IntSlice(ganjil))) // Ganjil diurutkan menurun
	sort.Ints(genap)                              // Genap diurutkan menaik

	return ganjil, genap
}

func cetakNomorRumah(wilayah int, ganjil []int, genap []int) {
	fmt.Printf("\nNomor rumah terurut wilayah %d:\n", wilayah)
	fmt.Print("Ganjil: ")
	for _, nomor := range ganjil {
		fmt.Printf("%d ", nomor)
	}
	fmt.Print("\nGenap: ")
	for _, nomor := range genap {
		fmt.Printf("%d ", nomor)
	}
	fmt.Println()
}

func main() {
	var totalWilayah int

	// Input jumlah wilayah
	fmt.Print("Masukkan jumlah wilayah: ")
	fmt.Scanln(&totalWilayah)

	// Memproses setiap wilayah
	for wilayah := 1; wilayah <= totalWilayah; wilayah++ {
		var totalRumah int

		fmt.Printf("\nJumlah rumah di wilayah %d: ", wilayah)
		fmt.Scanln(&totalRumah)

		// Membuat slice untuk menyimpan nomor rumah
		nomorRumah := make([]int, totalRumah)
		fmt.Printf("Masukkan %d nomor rumah (pisahkan dengan spasi): ", totalRumah)
		for i := 0; i < totalRumah; i++ {
			fmt.Scan(&nomorRumah[i])
		}

		// Memproses pengelompokan dan pengurutan
		ganjil, genap := prosesNomorRumah(nomorRumah)

		// Menampilkan hasil
		cetakNomorRumah(wilayah, ganjil, genap)
	}
}

```
## Output:![Screenshot 2024-12-01 220642](https://github.com/user-attachments/assets/6faabafe-bf5f-446d-8630-9b7cc900a93b)

Program ini memisahkan dan mengurutkan nomor rumah berdasarkan jenisnya (ganjil atau genap) untuk beberapa wilayah. Hasilnya adalah nomor rumah ganjil yang diurutkan dari besar ke kecil dan nomor rumah genap yang diurutkan secara dari kecil ke besar.


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
	"fmt"
	"sort"
)

// Fungsi untuk menghitung median dari array yang sudah terurut
func cariMedian(data []int) float64 {
	n := len(data)
	if n == 0 {
		return 0
	}
	// Jika jumlah elemen ganjil
	if n%2 != 0 {
		return float64(data[n/2])
	}
	// Jika jumlah elemen genap
	tengah := n / 2
	return float64(data[tengah-1]+data[tengah]) / 2.0
}

func main() {
	var angka int
	var nilai []int
	fmt.Println("Masukkan angka:")

	for {
		fmt.Scan(&angka)

		if angka == 0 {
			if len(nilai) > 0 {
				// Urutkan data menggunakan library `sort`
				sort.Ints(nilai)
				// Hitung median
				median := cariMedian(nilai)
				fmt.Printf("%.0f\n", median)
			}
		} else if angka == -5313 {
			// Keluar dari program
			break
		} else {
			// Tambahkan angka ke daftar nilai
			nilai = append(nilai, angka)
		}
	}
}

```
## Output:![Screenshot 2024-12-01 205030](https://github.com/user-attachments/assets/69bf7873-84df-492c-84df-7878fc2a2d20)


Program ini berfungsi untuk menghitung median dari sekumpulan data yang diinputkan oleh pengguna. Program berjalan secara langsung, menerima data hingga perintah khusus diberikan, lalu mencetak hasil median dari data angka-angka yang telah dimasukkan.


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

// Konstanta untuk ukuran maksimum array
const nMax = 7919

// Struktur data untuk Buku
type Buku struct {
	ID        int
	Judul     string
	Penulis   string
	Penerbit  string
	Tahun     int
	Rating    int
}

// Tipe data untuk daftar buku
type DaftarBuku [nMax]Buku

// Fungsi utama
func main() {
	var pustaka DaftarBuku
	var nPustaka int
	var ratingCari int

	// Input jumlah buku
	fmt.Print("Masukkan jumlah buku (N): ")
	fmt.Scanln(&nPustaka)

	// Input data buku
	for i := 0; i < nPustaka; i++ {
		fmt.Printf("Masukkan data buku ke-%d:\n", i+1)
		fmt.Printf("ID: ")
		fmt.Scanln(&pustaka[i].ID)
		fmt.Printf("Judul: ")
		fmt.Scanln(&pustaka[i].Judul)
		fmt.Printf("Penulis: ")
		fmt.Scanln(&pustaka[i].Penulis)
		fmt.Printf("Penerbit: ")
		fmt.Scanln(&pustaka[i].Penerbit)
		fmt.Printf("Tahun: ")
		fmt.Scanln(&pustaka[i].Tahun)
		fmt.Printf("Rating: ")
		fmt.Scanln(&pustaka[i].Rating)
	}

	// Input rating yang akan dicari
	fmt.Print("Masukkan rating yang dicari: ")
	fmt.Scanln(&ratingCari)

	// Cetak buku terfavorit
	cetakTerfavorit(pustaka, nPustaka)

	// Cetak buku berdasarkan rating tertinggi
	cetakUrutanRating(pustaka, nPustaka)

	// Cari buku berdasarkan rating
	cariBuku(pustaka, nPustaka, ratingCari)
}

// Fungsi untuk mencetak buku terfavorit
func cetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}

	terfavorit := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].Rating > terfavorit.Rating {
			terfavorit = pustaka[i]
		}
	}

	fmt.Println("\nBuku Terfavorit:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",terfavorit.Judul, terfavorit.Penulis, terfavorit.Penerbit, terfavorit.Tahun, terfavorit.Rating)
}

// Fungsi untuk mencetak buku berdasarkan urutan rating tertinggi
func cetakUrutanRating(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}

	// Salin data ke slice untuk diurutkan
	bukuSlice := pustaka[:n]

	// Urutkan berdasarkan rating (descending)
	sort.Slice(bukuSlice, func(i, j int) bool {
		return bukuSlice[i].Rating > bukuSlice[j].Rating
	})

	fmt.Println("\nBuku Berdasarkan Rating Tertinggi:")
	for i, buku := range bukuSlice {
		fmt.Printf("%d. Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
			i+1, buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Rating)
	}
}

// Fungsi untuk mencari buku berdasarkan rating tertentu
func cariBuku(pustaka DaftarBuku, n int, rating int) {
	if n == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}

	fmt.Printf("\nMencari buku dengan rating %d:\n", rating)
	found := false
	for _, buku := range pustaka[:n] {
		if buku.Rating == rating {
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
				buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Rating)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

```
## Output:![Screenshot 2024-12-01 202903](https://github.com/user-attachments/assets/e551984f-6030-450e-bd78-bc8227d23a2b)


Program ini mengelola data perpustakaan dengan fitur pencarian, pengurutan, dan pencetakan data buku berdasarkan rating. Program menggunakan array statis DaftarBuku untuk menyimpan informasi buku seperti ID, judul, penulis, penerbit, tahun, dan rating
## KESIMPULAN
Semua program dirancang untuk menyelesaikan tugas-tugas spesifik terkait manipulasi data, seperti pengurutan, pengelompokan, analisis, atau pencarian. Mereka menunjukkan pentingnya algoritma, logika, dan struktur program yang baik untuk menghasilkan solusi yang efisien dan mudah digunakan.
