## a. Penjelasan Singkat
Program ini menggunakan algoritma Brute Force 
Aturan mainnya adalah:
1.  **Satu per Baris & Kolom**: Setiap baris dan kolom hanya boleh diisi 1 Ratu.
2.  **Satu per Wilayah**: Setiap area (yang ditandai dengan huruf yang sama) hanya boleh diisi 1 Ratu.
3.  **Tidak Bersentuhan**: Ratu tidak boleh saling bersentuhan, baik secara tegak lurus maupun diagonal (jarak mutlak antar Ratu > 1).

Fitur utama:
- Membaca konfigurasi papan dari file `.txt`.
- Animasi visualisasi proses pencarian solusi di terminal.
- Menampilkan waktu eksekusi dan jumlah kasus yang ditinjau.
- Menyimpan solusi akhir ke file eksternal.

## b. Requirements
- **Go (Golang)**
- **Terminal/Console**

## c. Compile
1. Buka terminal di folder kode program.
2. Jalankan perintah berikut untuk execute
   go run main.go

## d. How to Use
Siapkan File Input:
Buat file teks yang berisi map boardgame menggunakan huruf. 

Contoh isi file soal.txt (4x4):

AABB
AACC
DDCC
DDEE

Interaksi Program:

Masukkan nama file: Ketik nama file.

Visualisasi: Ketik Ya jika ingin melihat animasi pencarian, atau Tidak untuk hasil instan.

Simpan: Jika solusi ditemukan, bisa menyimpannya ke file baru dengan mengetik Ya saat diminta.

## E.  identitas
Daniel Putra Rywandi S
13524143
