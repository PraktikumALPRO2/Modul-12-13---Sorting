<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL XII & XIII</strong></h2>
<h2 align="center"><strong> PENGURUTAN DATA </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Yesika Widiyani / 2311102195 <br>
  S1 IF11-05
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
# DASAR TEORI
## 1. Selection Sort
Selection Sort adalah algoritma pengurutan yang bekerja dengan cara membagi array menjadi dua bagian: bagian yang sudah terurut dan bagian yang belum terurut. Algoritma ini secara berulang-ulang memilih elemen terkecil dari bagian yang belum terurut dan menempatkannya ke posisi yang benar dalam bagian yang sudah terurut.

**Ciri-ciri Selection Sort:**
- Proses utama adalah *seleksi* elemen terkecil (atau terbesar, tergantung kebutuhan).
- Pengurutan dilakukan secara in-place (tidak membutuhkan ruang tambahan).
- Tidak stabil karena elemen dengan nilai yang sama dapat bertukar posisi relatif.

**Keunggulan:**
1. Implementasi Sederhana: Mudah dipahami dan diimplementasikan karena logikanya hanya mencari elemen terkecil/besar dan menukarnya.
2. Ramah Memori: Menggunakan sorting in-place sehingga hanya membutuhkan ruang tambahan sebesar \(O(1)\).
3. Tidak Bergantung pada Urutan Awal: Kompleksitas waktu tidak berubah meskipun data sudah terurut sebagian.

**Kekurangan:**
1. Kompleksitas Waktu Tinggi: 
   - Selalu membutuhkan \(O(n^2)\) perbandingan, bahkan jika data sudah terurut.
2. Tidak Stabil: Jika ada elemen dengan nilai yang sama, urutan relatifnya dapat berubah.
3. Kurang Efisien untuk Data Besar: Tidak cocok untuk dataset besar dibandingkan algoritma lain seperti Merge Sort atau Quick Sort.

## 2. Insertion Sort
**Insertion Sort** adalah algoritma pengurutan yang bekerja dengan cara menyusun elemen satu per satu. Setiap elemen dari array yang belum terurut disisipkan ke dalam posisi yang benar di bagian array yang sudah terurut.

**Ciri-ciri Insertion Sort:**
- Proses utama adalah *penyisipan* elemen pada posisi yang benar.
- Stabil, artinya elemen dengan nilai yang sama mempertahankan urutan relatifnya.
- Efisien untuk array kecil atau hampir terurut.

**Keunggulan:**
1. Stabil: Urutan relatif elemen yang memiliki nilai sama tetap terjaga.
2. Efisien untuk Dataset Kecil: Lebih cepat dibandingkan algoritma lain untuk array kecil.
3. Efisien untuk Data Hampir Terurut: Waktu eksekusi menjadi \(O(n)\) jika array hampir terurut.
4. Ramah Memori: Sorting dilakukan in-place, hanya membutuhkan ruang tambahan \(O(1)\).

**Kekurangan:**
1. Kompleksitas Tinggi untuk Data Acak:
   - Kompleksitas waktu terburuk tetap \(O(n^2)\).
2. Kurang Efisien untuk Data Besar: Tidak cocok untuk dataset besar dibandingkan algoritma dengan kompleksitas lebih rendah seperti Quick Sort atau Heap Sort.

------

# GUIDED
## Guided - 1
### Study Case
Hercules, preman terkenal seantero Ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort.**

_Masukan dimulai dengan sebuah integer (01000), banyaknya daerah kerabat Hercules tinggal. Isi a baris berikutnya selalu dimulai dengan sebuah integer m (0<m< 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkalan bilangan bulat positif, nomor rumah para kerabat._

_Keluaran terdiri dari a baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing- masing daerah._

### Source Code
![carbon (11)](https://github.com/user-attachments/assets/842a08a6-4f78-4274-9f4d-76b37d30978b)

### Screenshot Output
![image](https://github.com/user-attachments/assets/a27f7238-b917-41f8-8e27-6a232020ac2a)

### Deskripsi Program
Program ini digunakan untuk mengurutkan nomor rumah kerabat di setiap daerah secara menaik menggunakan algoritma Selection Sort. Pengguna diminta untuk memasukkan jumlah daerah, jumlah nomor rumah di setiap daerah, serta daftar nomor rumahnya. Setelah data dimasukkan, program akan mengurutkan nomor rumah tersebut dan menampilkan hasilnya dalam urutan menaik untuk setiap daerah.

**Algoritma Program**  
1. Minta pengguna memasukkan jumlah daerah.  
2. Untuk setiap daerah:  
   - Minta jumlah nomor rumah dan data nomor rumah.  
   - Urutkan nomor rumah menggunakan **Selection Sort**.  
   - Tampilkan hasil nomor rumah yang sudah terurut.

**Cara Kerja Program**  
- **Input Data**: Pengguna memasukkan jumlah daerah dan daftar nomor rumah untuk setiap daerah.  
- **Proses Sorting**: Nomor rumah diurutkan menggunakan algoritma **Selection Sort**, di mana elemen terkecil dipindahkan ke posisi yang sesuai pada setiap iterasi.  
- **Output Data**: Program mencetak nomor rumah yang sudah terurut untuk setiap daerah.

## Guided - 2
### Study Case
Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda Insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.

Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array.

Keluaran terdiri dari dua haris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".

### Source Code
![carbon (12)](https://github.com/user-attachments/assets/9b4d0401-c095-49d5-a932-a400f47021f9)

### Screenshot Output
![image](https://github.com/user-attachments/assets/b7e4d02a-7635-4868-b0a3-b8db4ef77417)

### Deskripsi Program
Program ini membaca sekumpulan bilangan bulat dari pengguna hingga bilangan negatif dimasukkan sebagai penanda akhir input. Data yang dimasukkan kemudian diurutkan menggunakan algoritma Insertion Sort. Setelah diurutkan, program memeriksa apakah selisih antar elemen di dalam array adalah tetap (konstan).

Program menampilkan hasil array yang sudah terurut dan informasi apakah data memiliki selisih tetap beserta nilai selisihnya, jika ada.

**Algoritma Program**  
1. **Input Data**:  
   - Terima bilangan bulat dari pengguna hingga bilangan negatif dimasukkan.  
   - Simpan data ke dalam array.  

2. **Sorting**:  
   - Urutkan array menggunakan algoritma **Insertion Sort**.  

3. **Pemeriksaan Selisih**:  
   - Periksa apakah selisih antara setiap elemen berturut-turut dalam array adalah tetap.  
   - Jika selisih tetap, simpan nilai selisih tersebut.  

4. **Output**:  
   - Cetak array yang sudah terurut.  
   - Tampilkan status apakah data memiliki selisih tetap, dan jika iya, tunjukkan nilai selisihnya.  

**Cara Kerja Program**  
1. Program membaca input dari pengguna dan menghentikan input saat bilangan negatif dimasukkan.  
2. Data yang dikumpulkan diurutkan dengan memindahkan elemen ke posisi yang tepat menggunakan **Insertion Sort**.  
3. Setelah pengurutan, program menghitung selisih antar elemen berturut-turut:  
   - Jika semua selisih sama, data dianggap memiliki jarak tetap.  
   - Jika tidak, data dianggap tidak memiliki jarak tetap.  
4. Program menampilkan array terurut dan informasi tentang jarak elemen.

------

# UNGUIDED
## Unguided-1
### Study Case
Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari normor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.

Masukan masih persis sama seperti sebelumnya.

Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing masing daerah

### Source Code
![carbon (13)](https://github.com/user-attachments/assets/c32d6b4a-1d84-4e0c-bee6-cdce607f01f3)

### Screenshot Output
![image](https://github.com/user-attachments/assets/f11ca7fa-023a-4686-a8c2-a8a45cc287c4)

### Deskripsi
Program ini digunakan untuk mengurutkan nomor rumah kerabat di beberapa daerah sesuai dengan aturan tertentu. Nomor rumah di setiap daerah dipisahkan menjadi dua kelompok:
Nomor Ganjil: Diurutkan secara menaik (ascending).
Nomor Genap: Diurutkan secara menurun (descending).
Nomor rumah ditampilkan dengan format nomor ganjil lebih dahulu diikuti oleh nomor genap di setiap daerah.

**Algoritma Program**  
1. **Input Data**:  
   - Masukkan jumlah daerah \(n\).  
   - Untuk setiap daerah, masukkan jumlah rumah \(m\) dan daftar nomor rumah.  

2. **Pemisahan Data**:  
   - Pisahkan nomor rumah menjadi dua kelompok:  
     - **Ganjil**: Disimpan di array ganjil.  
     - **Genap**: Disimpan di array genap.  

3. **Proses Sorting**:  
   - Urutkan array ganjil secara menaik (ascending).  
   - Urutkan array genap secara menurun (descending).  

4. **Output Data**:  
   - Tampilkan nomor rumah ganjil (menaik).  
   - Diikuti nomor rumah genap (menurun).  

**Cara Kerja Program**  
1. Program membaca input pengguna untuk jumlah daerah dan daftar nomor rumah di setiap daerah.  
2. Setiap nomor rumah dipisahkan ke dalam array ganjil atau genap.  
3. Nomor rumah ganjil diurutkan dari kecil ke besar, sedangkan nomor rumah genap diurutkan dari besar ke kecil menggunakan **Selection Sort**.  
4. Program mencetak hasil urutan nomor rumah sesuai aturan: ganjil terlebih dahulu, lalu genap.

## Unguided-2
### Study case
Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan, Dari 13 problem yang diberikan, ada satu problem yang menarik, Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?

"Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.

Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data O merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313.

Keluaran adalah median yang diminta, satu data per baris.

### Source code
![carbon (14)](https://github.com/user-attachments/assets/e8132a79-1703-47ea-b2ac-2836593d3b09)

### Screenshot Output
![image](https://github.com/user-attachments/assets/9a524072-dbf7-46c7-9b4c-117de41d5d68)

### Deskripsi Program
Program ini menghitung median dari rangkaian data bilangan bulat positif yang terus bertambah, di mana perintah pencetakan median diberikan melalui masukan angka 0. Program dirancang untuk bekerja secara efisien meskipun data yang dimasukkan mencapai jutaan elemen.

**Algoritma Program**  

1. **Inisialisasi**:  
   - Buat dua heap:  
     - **MaxHeap**: Menyimpan setengah data terkecil (elemen terbesar di atas).  
     - **MinHeap**: Menyimpan setengah data terbesar (elemen terkecil di atas).  

2. **Baca Input**:  
   - Terima angka satu per satu.  
   - Jika angka:  
     - **Positif**: Masukkan ke salah satu heap dengan aturan:  
       - Masukkan ke **MaxHeap** jika lebih kecil atau sama dengan elemen terbesar di **MaxHeap**.  
       - Masukkan ke **MinHeap** jika lebih besar.  
       - Seimbangkan jumlah elemen di kedua heap.  
     - **0**: Hitung dan cetak median:  
       - Jika jumlah elemen tidak seimbang, median adalah elemen terbesar di **MaxHeap**.  
       - Jika seimbang, median adalah rata-rata elemen terbesar di **MaxHeap** dan elemen terkecil di **MinHeap** (dibulatkan ke bawah).  
     - **-5313**: Hentikan program.  

**Cara Kerja Program Secara Singkat**  

1. **Data Masuk**:  
   - Data dipisah ke **MaxHeap** (elemen kecil) dan **MinHeap** (elemen besar) agar median dapat dihitung secara dinamis.  

2. **Seimbangkan Heap**:  
   - Jika salah satu heap memiliki elemen lebih dari satu dibandingkan heap lainnya, pindahkan elemen antar heap untuk menjaga keseimbangan.  

3. **Hitung Median**:  
   - Jika **MaxHeap** lebih banyak elemen, median adalah elemen teratas **MaxHeap**.  
   - Jika kedua heap seimbang, median adalah rata-rata teratas dari **MaxHeap** dan **MinHeap**.  

4. **Efisiensi**:  
   - Operasi heap (\(O(\log n)\)) memastikan program tetap cepat meski data berjumlah besar.  

## Unguided-3
### Study Case
Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:
![image](https://github.com/user-attachments/assets/cd60a6bf-dbfd-460d-abe8-11f0b190f2de)

Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.

Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku. yang dicari sesual rating yang diberikan pada masukan baris terakhir.

### Source Code
![carbon (15)](https://github.com/user-attachments/assets/558fce90-93d6-4850-ab55-790c4a044db8)

### Screenshot Output
![image](https://github.com/user-attachments/assets/43d0bae0-5025-4863-81f0-7af897ec88a1)

### Deskripsi Program
Program ini adalah sistem pengelolaan data buku untuk perpustakaan. Program membaca data sejumlah buku, menampilkan buku dengan rating tertinggi, lima buku dengan rating tertinggi, serta mencari buku berdasarkan rating tertentu. Data buku diurutkan berdasarkan rating secara menurun untuk mempermudah pencarian dan penampilan hasil.

**Algoritma Program**

1. **Input Data Buku:**
   - Masukkan jumlah buku (**N**).
   - Masukkan data buku berupa ID, judul, penulis, penerbit, jumlah eksemplar, tahun, dan rating untuk setiap buku.

2. **Menentukan Buku Terfavorit:**
   - Iterasi seluruh data buku untuk menemukan buku dengan rating tertinggi.

3. **Mengurutkan Buku Berdasarkan Rating:**
   - Gunakan algoritma **insertion sort** untuk mengurutkan buku dari rating tertinggi ke terendah.

4. **Menampilkan Lima Buku Rating Tertinggi:**
   - Cetak judul dari lima buku pertama pada daftar yang telah diurutkan (atau kurang dari lima jika data kurang dari lima buku).

5. **Mencari Buku Berdasarkan Rating:**
   - Gunakan metode **binary search** pada data yang telah diurutkan untuk mencari buku dengan rating tertentu.
   - Jika ditemukan, tampilkan informasi buku tersebut, jika tidak, tampilkan pesan bahwa buku tidak ditemukan.

**Cara Kerja Program**

1. Program dimulai dengan meminta masukan jumlah buku dan data buku.
2. Data buku diolah untuk menemukan buku dengan rating tertinggi.
3. Data buku diurutkan menggunakan **insertion sort** untuk memudahkan pengolahan lebih lanjut.
4. Setelah pengurutan:
   - Buku dengan lima rating tertinggi ditampilkan.
   - Pencarian buku dengan rating tertentu dilakukan menggunakan **binary search**.
5. Hasil pencarian, buku terfavorit, dan daftar lima buku terbaik ditampilkan.



