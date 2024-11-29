<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL XII-XIII -</strong></h2>
<h2 align="center"><strong> PENGURUTAN DATA </strong></h2>

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
### Definisi Pengurutan Data
Pengurutan data adalah proses mengatur elemen-elemen dalam suatu daftar berdasarkan urutan tertentu, seperti urutan menaik atau menurun. Teknik pengurutan ini digunakan secara luas di berbagai bidang untuk meningkatkan efisiensi pencarian, analisis data, dan pemrosesan lainnya.

#### 1. Ide Algoritma Selection Sort
Algoritma selection sort didasarkan pada konsep mencari elemen terkecil (atau terbesar) dari daftar yang belum diurutkan dan menempatkannya pada posisi yang sesuai. Proses ini diulangi untuk setiap elemen dalam daftar hingga seluruh daftar terurut. Prinsip dasarnya adalah membagi daftar menjadi dua bagian: bagian yang sudah diurutkan dan bagian yang belum diurutkan, lalu secara bertahap memperluas bagian yang sudah diurutkan dengan menambahkan elemen yang ditemukan【1】【2】.

#### 2. Algoritma Selection Sort
Langkah-langkah selection sort dapat dijelaskan sebagai berikut:

1. Mulai dari indeks pertama, temukan elemen terkecil di bagian daftar yang belum diurutkan.
2. Tukarkan elemen terkecil tersebut dengan elemen di posisi awal bagian yang belum diurutkan.
3. Geser batas bagian yang belum diurutkan satu elemen ke kanan.
4. Ulangi langkah 1-3 hingga semua elemen terurut.
Kompleksitas waktu selection sort dalam kasus terbaik, rata-rata, dan terburuk adalah O(n^2), di mana n adalah jumlah elemen dalam daftar. Algoritma ini tidak efisien untuk daftar berukuran besar tetapi sederhana untuk diimplementasikan【1】【3】.

#### 3. Ide Algoritma Insertion Sort
Algoritma insertion sort bekerja dengan cara membangun bagian daftar yang terurut secara bertahap. Elemen dari bagian yang belum diurutkan dimasukkan satu per satu ke dalam posisi yang tepat di bagian yang sudah diurutkan. Prinsip utamanya adalah memanfaatkan fakta bahwa bagian kecil dari daftar yang terurut dapat dipertahankan dengan biaya minimal untuk penyisipan elemen baru【4】【5】.

#### 4. Algoritma Insertion Sort
Langkah-langkah insertion sort adalah sebagai berikut:
1. Mulai dari elemen kedua dalam daftar, anggap elemen pertama sudah terurut.
2. Ambil elemen berikutnya dan bandingkan dengan elemen dalam bagian yang sudah diurutkan.
3. Geser elemen yang lebih besar ke kanan untuk membuat ruang bagi elemen yang sedang dimasukkan.
4. Tempatkan elemen tersebut pada posisi yang sesuai dalam bagian yang sudah diurutkan.
5. Ulangi langkah 2-4 untuk semua elemen dalam daftar.
Kompleksitas waktu insertion sort adalah O(n) dalam kasus terbaik (daftar sudah hampir terurut) dan O(n^2) dalam kasus terburuk (daftar terbalik). Algoritma ini efektif untuk daftar kecil dan bekerja dengan baik pada data yang hampir terurut【4】【6】.

## II. UNGUIDED
## 1. 
#### Source Code
```go

```
#### Screenshoot Source Code

#### Screenshoot Output

#### Deskripsi Program

#### Algoritma Program

#### Cara Kerja

## 2. 
#### Source Code
```go

```
#### Screenshoot Source Code

#### Screenshoot Output


#### Deskripsi Program


#### Algoritma Program


#### Cara Kerja


## 3. 
#### Source Code
```go

```

#### Screenshoot Source Code

#### Screenshoot Output


#### Deskripsi Program


#### Algoritma Program

#### Cara Kerja


### Kesimpulan
Selection sort dan insertion sort adalah algoritma pengurutan sederhana. Selection sort bekerja dengan mencari elemen terkecil lalu memindahkannya ke posisi yang tepat, sedangkan insertion sort menyisipkan elemen ke tempat yang sesuai di bagian yang sudah terurut. Keduanya cocok untuk data kecil, tetapi kurang efisien untuk data besar karena memiliki kompleksitas waktu O(n^2) pada kasus rata-rata dan terburuk.

## Referensi 
[1] T. H. Cormen, C. E. Leiserson, R. L. Rivest, and C. Stein, Introduction to Algorithms, 3rd ed. Cambridge, MA: MIT Press, 2009.

[2] R. Sedgewick and K. Wayne, Algorithms, 4th ed. Boston, MA: Addison-Wesley, 2011.

[3] M. A. Weiss, Data Structures and Algorithm Analysis in C++, 4th ed. Upper Saddle River, NJ: Pearson, 2014.

[4] S. Dasgupta, C. H. Papadimitriou, and U. Vazirani, Algorithms, New York, NY: McGraw-Hill, 2008.

[5] A. Drozdek, Data Structures and Algorithms in C++, 4th ed. Boston, MA: Cengage Learning, 2013.

[6] J. Kleinberg and É. Tardos, Algorithm Design, Boston, MA: Pearson, 2006.
