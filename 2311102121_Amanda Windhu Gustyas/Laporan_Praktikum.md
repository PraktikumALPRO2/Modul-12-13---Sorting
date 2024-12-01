# <h2 align="center">LAPORAN PRAKTIKUM</h2>
# <h2 align="center">ALGORITMA DAN PEMROGRAMAN 2</h2>
# <h2 align="center">MODUL </h2>
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
## Output: ![image](https://github.com/user-attachments/assets/dcad3bc1-ba98-42bf-a86b-583599c6bb07)


### 2. 






 




