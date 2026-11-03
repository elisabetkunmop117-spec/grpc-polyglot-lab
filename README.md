 

## D. LANGKAH-LANGKAH PRAKTIKUM

## Langkah 1. Memeriksa Instalasi Perangkat Lunak

**Tujuan:** Memastikan perangkat lunak yang diperlukan telah tersedia sebelum project dibuat.

1. Buka Windows PowerShell atau terminal VS Code.
2. Jalankan perintah berikut satu per satu:

```powershell
git --version
python --version
node -v
npm -v
go version
protoc --version
```

3. Periksa hasil keluaran versi setiap perangkat lunak.

**Analisis:**

<p align="justify">Perintah tersebut digunakan untuk memeriksa apakah Git, Python, Node.js, npm, Go, dan Protocol Buffers Compiler telah terpasang serta dapat dikenali oleh sistem. Jika terminal menampilkan informasi versi, berarti perangkat lunak dapat dipanggil melalui terminal.</p>

**Dokumentasi:** `dokumentasi/gambar/01-pemeriksaan-instalasi.png`

## Langkah 2. Mengatur Environment Variable Protoc

1. Buka menu Start Windows.
2. Cari dan buka **Edit the system environment variables**.
3. Pilih tab **Advanced**.
4. Klik **Environment Variables**.
5. Pada bagian User variables, pilih variabel `Path`.
6. Klik **Edit**, kemudian pilih **New**.
7. Masukkan lokasi folder `bin` tempat `protoc.exe` disimpan. Contoh:

```text
D:\protoc-25.9-win64\bin
```

8. Klik OK sampai semua jendela tertutup.
9. Buka terminal baru dan jalankan:

```powershell
protoc --version
```

10. Pastikan versi Protobuf berhasil ditampilkan.

**Analisis:**

<p align="justify">Pengaturan Environment Variable dilakukan agar Windows dapat mengenali program <code>protoc</code> melalui terminal tanpa harus menuliskan lokasi lengkap file executable. Dengan konfigurasi tersebut, compiler dapat digunakan dari berbagai direktori project.</p>

**Dokumentasi:** `dokumentasi/gambar/02-konfigurasi-path-protoc.png`

## Langkah 3. Membuat Struktur Folder Project

1. Buka PowerShell.
2. Pindah ke drive dan direktori tempat project akan disimpan:

```powershell
D:
mkdir grpc-polyglot-lab
cd grpc-polyglot-lab
```

3. Buat folder untuk server dan client:

```powershell
mkdir go-server
mkdir python-client
mkdir node-client
```

4. Periksa struktur folder:

```powershell
tree /F
```

**Analisis:**

<p align="justify">Perintah <code>mkdir</code> digunakan untuk membuat direktori project dan direktori untuk setiap komponen aplikasi. Folder <code>go-server</code> digunakan untuk menyimpan server Go, sedangkan <code>python-client</code> dan <code>node-client</code> digunakan untuk menyimpan aplikasi client. Perintah <code>tree /F</code> menampilkan susunan folder beserta file yang terdapat di dalamnya.</p>

**Dokumentasi:** `dokumentasi/gambar/03-struktur-folder-project.png`

## Langkah 4. Membuat File user.proto

1. Pastikan terminal berada di folder utama project.
2. Buat file menggunakan perintah:

```powershell
notepad user.proto
```

3. Masukkan definisi pesan dan layanan berikut:

```protobuf
syntax = "proto3";

package user;

option go_package = "grpc-polyglot-lab/go-server/pb;pb";

message UserRequest {
  int32 user_id = 1;
}

message UserResponse {
  string name = 1;
  string email = 2;
}

service UserService {
  rpc GetUserProfile (UserRequest) returns (UserResponse);
}
```

4. Simpan file dengan nama `user.proto`.
5. Pastikan ekstensi file bukan `user.proto.txt`.

**Analisis:**

<p align="justify">File <code>user.proto</code> digunakan untuk mendefinisikan struktur permintaan <code>UserRequest</code>, struktur respons <code>UserResponse</code>, dan layanan <code>UserService</code>. File ini menjadi kontrak komunikasi yang digunakan oleh server dan client agar memiliki pemahaman struktur data yang sama.</p>

**Dokumentasi:** `dokumentasi/gambar/04-file-user-proto-awal.png`

## Langkah 5. Menyiapkan Modul dan Dependensi Go

1. Masuk ke folder server:

```powershell
cd go-server
```

2. Inisialisasi modul Go:

```powershell
go mod init grpc-polyglot-lab/go-server
```

3. Pasang library gRPC dan Protobuf:

```powershell
go get google.golang.org/grpc
go get google.golang.org/protobuf
```

4. Pasang plugin generator kode Go:

```powershell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

5. Periksa plugin generator:

```powershell
where.exe protoc-gen-go
where.exe protoc-gen-go-grpc
```

**Analisis:**

<p align="justify">Perintah <code>go mod init</code> digunakan untuk membuat identitas modul Go. Perintah <code>go get</code> menambahkan library yang diperlukan untuk komunikasi gRPC dan pemrosesan Protobuf. Sementara itu, <code>go install</code> memasang plugin yang digunakan oleh compiler Protobuf untuk menghasilkan kode Go.</p>

**Dokumentasi:** `dokumentasi/gambar/05-inisialisasi-modul-go.png`

## Langkah 6. Generate Kode Protobuf dan gRPC untuk Go

1. Pastikan terminal berada di folder `go-server`.
2. Buat folder untuk menyimpan kode hasil generate:

```powershell
mkdir pb
```

3. Jalankan perintah berikut dalam satu baris:

```powershell
protoc --proto_path=.. --go_out=. --go_opt=module=grpc-polyglot-lab/go-server --go-grpc_out=. --go-grpc_opt=module=grpc-polyglot-lab/go-server ..\user.proto
```

4. Periksa hasil generate:

```powershell
Get-ChildItem pb
```

File yang dihasilkan:

```text
user.pb.go
user_grpc.pb.go
```

**Analisis:**

<p align="justify">Perintah <code>protoc</code> digunakan untuk membaca file <code>user.proto</code> dan menghasilkan kode Go. File <code>user.pb.go</code> berisi struktur pesan Protobuf, sedangkan <code>user_grpc.pb.go</code> berisi kode pendukung komunikasi layanan gRPC. Kedua file tersebut disimpan dalam folder <code>pb</code>.</p>

**Dokumentasi:** `dokumentasi/gambar/06-generate-kode-go.png`

## Langkah 7. Membuat Program Go gRPC Server

1. Pastikan terminal berada di folder `go-server`.
2. Buat file server:

```powershell
notepad server.go
```

3. Masukkan kode program server Go yang menggunakan `UserService`.
4. Simpan file dengan nama `server.go`.
5. Rapikan dependensi:

```powershell
go mod tidy
```

6. Jalankan server:

```powershell
go run server.go
```

7. Pastikan server berjalan pada port `50051`.

**Analisis:**

<p align="justify">Program server Go bertugas menerima permintaan dari client melalui protokol gRPC. Server membuka koneksi TCP pada port <code>50051</code>, membuat instance server gRPC, dan mendaftarkan layanan <code>UserService</code>. Metode <code>GetUserProfile</code> digunakan untuk menangani permintaan profil pengguna berdasarkan ID yang dikirimkan oleh client.</p>

**Dokumentasi:** `dokumentasi/gambar/07-server-go-berjalan.png`

## Langkah 8. Menyiapkan Client Python

1. Buka terminal baru.
2. Masuk ke folder `python-client`.
3. Pasang library yang diperlukan:

```powershell
python -m pip install grpcio grpcio-tools
```

4. Generate kode Python dari file `user.proto`:

```powershell
python -m grpc_tools.protoc -I.. --python_out=. --grpc_python_out=. ..\user.proto
```

5. Buat file client:

```powershell
notepad client.py
```

6. Masukkan kode client Python yang menggunakan `UserServiceStub`.
7. Simpan file dengan nama `client.py`.

**Analisis:**

<p align="justify">Library <code>grpcio</code> digunakan untuk komunikasi gRPC, sedangkan <code>grpcio-tools</code> menyediakan compiler untuk menghasilkan kode Python berdasarkan definisi Protobuf. File <code>user_pb2.py</code> berisi struktur pesan, sementara <code>user_pb2_grpc.py</code> berisi kode pendukung layanan gRPC.</p>

**Dokumentasi:** `dokumentasi/gambar/08-persiapan-client-python.png`

## Langkah 9. Menguji Client Python

1. Pastikan server Go masih berjalan pada terminal pertama.
2. Pada terminal Python, pastikan berada di folder `python-client`.
3. Jalankan:

```powershell
python client.py
```

4. Periksa apakah profil pengguna berhasil ditampilkan.

**Analisis:**

<p align="justify">Perintah <code>python client.py</code> digunakan untuk menjalankan aplikasi client Python. Client membuat koneksi ke server Go, mengirim permintaan melalui metode <code>GetUserProfile</code>, kemudian menerima dan menampilkan data profil pengguna dari server.</p>

**Dokumentasi:** `dokumentasi/gambar/09-pengujian-client-python.png`

## Langkah 10. Menyiapkan Client Node.js

1. Buka terminal baru.
2. Masuk ke folder `node-client`:

```powershell
cd node-client
```

3. Inisialisasi project Node.js:

```powershell
npm init -y
```

4. Pasang library yang diperlukan:

```powershell
npm install @grpc/grpc-js @grpc/proto-loader
```

5. Buat file client:

```powershell
notepad client.js
```

6. Masukkan kode client Node.js yang membaca file `user.proto` menggunakan `@grpc/proto-loader`.
7. Simpan file dengan nama `client.js`.

**Analisis:**

<p align="justify">Perintah <code>npm init -y</code> digunakan untuk membuat file konfigurasi <code>package.json</code>. Library <code>@grpc/grpc-js</code> digunakan untuk komunikasi gRPC, sedangkan <code>@grpc/proto-loader</code> digunakan untuk membaca definisi Protobuf saat aplikasi berjalan.</p>

**Dokumentasi:** `dokumentasi/gambar/10-persiapan-client-nodejs.png`

## Langkah 11. Menguji Client Node.js

1. Pastikan server Go masih berjalan.
2. Pada terminal Node.js, pastikan berada di folder `node-client`.
3. Jalankan:

```powershell
node client.js
```

4. Periksa hasil respons yang ditampilkan.

**Analisis:**

<p align="justify">Perintah <code>node client.js</code> digunakan untuk menjalankan aplikasi client Node.js. Program membaca skema <code>user.proto</code>, membuat koneksi ke server Go, mengirim permintaan profil pengguna, dan menampilkan respons yang diterima.</p>

**Dokumentasi:** `dokumentasi/gambar/11-pengujian-client-nodejs.png`

## Langkah 12. Menambahkan Field Phone pada user.proto

1. Buka terminal pada folder utama project.
2. Buka file Protobuf:

```powershell
notepad user.proto
```

3. Perbarui definisi `UserResponse` menjadi:

```protobuf
message UserResponse {
  string name = 1;
  string email = 2;
  string phone = 3;
}
```

4. Simpan perubahan.

**Analisis:**

<p align="justify">Field <code>phone</code> ditambahkan dengan nomor field <code>3</code>, sedangkan field <code>name</code> dan <code>email</code> tetap menggunakan nomor <code>1</code> dan <code>2</code>. Penambahan tersebut memperluas struktur respons agar server dapat mengirimkan informasi nomor telepon kepada client.</p>

**Dokumentasi:** `dokumentasi/gambar/12-penambahan-field-phone.png`

## Langkah 13. Generate Ulang Kode Protobuf

**A. Generate ulang kode Go**

1. Masuk ke folder `go-server`.
2. Jalankan:

```powershell
protoc --proto_path=.. --go_out=. --go_opt=module=grpc-polyglot-lab/go-server --go-grpc_out=. --go-grpc_opt=module=grpc-polyglot-lab/go-server ..\user.proto
```

**B. Generate ulang kode Python**

1. Buka terminal baru.
2. Masuk ke folder `python-client`.
3. Jalankan:

```powershell
python -m grpc_tools.protoc -I.. --python_out=. --grpc_python_out=. ..\user.proto
```

**Analisis:**

<p align="justify">Setiap perubahan pada definisi pesan Protobuf perlu diikuti dengan pembaruan kode hasil generate untuk bahasa yang menggunakan static code generation. Proses ini memastikan kode yang digunakan aplikasi sesuai dengan skema terbaru.</p>

**Dokumentasi:** `dokumentasi/gambar/13-generate-ulang-kode.png`

## Langkah 14. Memperbarui Server dan Client

1. Perbarui `go-server/server.go` agar respons mengisi tiga field, yaitu `Name`, `Email`, dan `Phone`.
2. Perbarui `python-client/client.py` agar menampilkan data berikut:

```python
print("=== Profil Pengguna ===")
print(f"Nama  : {response.name}")
print(f"Email : {response.email}")
print(f"Phone : {response.phone}")
```

3. Perbarui `node-client/client.js` agar menampilkan:

```javascript
console.log("=== Profil Pengguna ===");
console.log("Nama  :", response.name);
console.log("Email :", response.email);
console.log("Phone :", response.phone);
```

4. Simpan seluruh perubahan.
5. Hentikan server lama dengan `Ctrl + C`.
6. Jalankan kembali server:

```powershell
cd go-server
go run server.go
```

**Analisis:**

<p align="justify">Server harus mengisi field <code>Phone</code> pada objek respons. Client Python membaca nilai melalui <code>response.phone</code>, sedangkan client Node.js juga membaca nilai melalui <code>response.phone</code>. Dengan demikian, kedua client dapat menampilkan informasi nomor telepon yang dikirim oleh server.</p>

**Dokumentasi:** `dokumentasi/gambar/14-pembaruan-server-dan-client.png`

## Langkah 15. Pengujian Akhir

Pastikan server Go aktif pada port `50051`. Selanjutnya, jalankan client pada terminal masing-masing.

**Pengujian Python:**

```powershell
cd python-client
python client.py
```

**Pengujian Node.js:**

```powershell
cd node-client
node client.js
```

Contoh format keluaran yang diharapkan:

```text
=== Profil Pengguna ===
Nama  : Elisabet Kunmop
Email : elisabet@example.com
Phone : 000000000000
```

*Catatan: Keluaran di atas mengikuti data yang saat ini terdapat pada kode server kamu.*

**Analisis:**

<p align="justify">Pengujian dilakukan untuk memastikan server Go dapat memberikan respons yang memuat nama, email, dan nomor telepon. Keberhasilan client Python dan Node.js menampilkan ketiga data tersebut menunjukkan bahwa komunikasi gRPC dan perubahan kontrak data telah diterapkan dengan benar.</p>

**Dokumentasi:**
- `dokumentasi/gambar/15-server-go-aktif.png`
- `dokumentasi/gambar/16-hasil-client-python.png`
- `dokumentasi/gambar/17-hasil-client-nodejs.png`

## Langkah 16. Mengunggah Project ke GitHub

1. Pastikan terminal berada di folder utama `grpc-polyglot-lab`.
2. Periksa status repository:

```powershell
git status
```

3. Periksa alamat repository:

```powershell
git remote -v
```

4. Tambahkan seluruh file project:

```powershell
git add .
```

5. Buat commit:

```powershell
git commit -m "Menyelesaikan praktikum gRPC polyglot"
```

6. Unggah project ke GitHub:

```powershell
git push -u origin main
```

**Analisis:**

<p align="justify">Git digunakan untuk mengelola perubahan pada project, sedangkan GitHub digunakan sebagai tempat penyimpanan repository secara daring. Perintah <code>git add</code> menyiapkan perubahan, <code>git commit</code> menyimpan perubahan ke riwayat lokal, dan <code>git push</code> mengunggah commit ke repository GitHub.</p>

**Dokumentasi:** `dokumentasi/gambar/18-repository-github.png`

---

# E. HASIL DAN PEMBAHASAN

<p align="justify">Praktikum Laboratorium Poliglot gRPC menghasilkan sistem komunikasi sederhana yang menggunakan bahasa Go sebagai server serta Python dan Node.js sebagai client. Ketiga komponen menggunakan definisi layanan yang sama melalui file <code>user.proto</code>.</p>

<p align="justify">Server Go menyediakan layanan <code>GetUserProfile</code> pada port <code>50051</code>. Client mengirimkan permintaan menggunakan <code>user_id</code>, kemudian server mengembalikan data profil pengguna. Pada tahap evaluasi, struktur <code>UserResponse</code> diperbarui dengan menambahkan field <code>phone</code> menggunakan nomor field <code>3</code>.</p>

<p align="justify">Setelah perubahan skema dilakukan, kode Go dan Python diperbarui melalui proses generate ulang. Client Node.js membaca file <code>user.proto</code> saat runtime melalui pustaka <code>@grpc/proto-loader</code>. Server dan client kemudian disesuaikan agar informasi nomor telepon dapat dikirim dan ditampilkan.</p>

### Dokumentasi Hasil Praktikum

| No. | Dokumentasi | Keterangan |
|---|---|---|
| 1 | Server Go | Server berjalan pada port 50051. |
| 2 | Client Python | Client menampilkan profil pengguna. |
| 3 | Client Node.js | Client menampilkan profil pengguna. |
| 4 | GitHub | Project tersimpan dalam repository. |

# F. ANALISIS

### 1. Analisis Komunikasi Polyglot gRPC

<p align="justify">gRPC memungkinkan aplikasi yang dibuat menggunakan bahasa pemrograman berbeda untuk saling berkomunikasi melalui definisi layanan yang sama. Go digunakan sebagai server, sedangkan Python dan Node.js digunakan sebagai client. Kesamaan definisi pada file <code>user.proto</code> membantu menjaga kesesuaian struktur permintaan dan respons.</p>

### 2. Analisis Penambahan Field Phone

<p align="justify">Penambahan field <code>phone = 3</code> pada pesan <code>UserResponse</code> memperluas data profil pengguna tanpa mengubah nomor field yang sudah digunakan. Setelah skema diperbarui, kode hasil generate untuk Go dan Python diperbarui, server diubah agar mengirimkan nomor telepon, dan client disesuaikan agar menampilkan nilai tersebut.</p>

### 3. Perbandingan Static Code Generation dan Dynamic Runtime Loading

<p align="justify">Pada pendekatan <i>Static Code Generation</i>, kode dibuat terlebih dahulu dari file <code>user.proto</code> menggunakan compiler dan plugin yang sesuai. Dalam praktikum ini, kode Go dan Python dihasilkan sebelum aplikasi dijalankan. Apabila struktur pesan berubah, proses generate perlu dijalankan kembali agar kode yang digunakan sesuai dengan skema terbaru.</p>

<p align="justify">Pada pendekatan <i>Dynamic Runtime Loading</i>, definisi Protobuf dibaca ketika aplikasi berjalan. Client Node.js menggunakan <code>@grpc/proto-loader</code> untuk memuat file <code>user.proto</code> secara langsung. Pendekatan ini tidak memerlukan file JavaScript hasil generate secara terpisah untuk definisi pesan pada contoh client yang digunakan.</p>

| Aspek | Static Code Generation | Dynamic Runtime Loading |
|---|---|---|
| Waktu pemrosesan skema | Sebelum aplikasi dijalankan | Saat aplikasi berjalan |
| Contoh praktikum | Go dan Python | Node.js |
| Hasil proses | File kode hasil generate | Definisi dimuat dari file `.proto` |
| Ketika skema berubah | Kode perlu di-generate ulang | Skema dibaca saat runtime |
| Karakteristik | Kode hasil generate tersedia sebelum runtime | Memuat skema ketika aplikasi berjalan |

<p align="justify">Kedua pendekatan memiliki kegunaan masing-masing. Static Code Generation menyediakan kode hasil generate yang dapat digunakan oleh aplikasi, sedangkan Dynamic Runtime Loading memungkinkan aplikasi membaca skema ketika dijalankan. Pemilihan pendekatan bergantung pada kebutuhan dan rancangan aplikasi.</p>

# G. KESIMPULAN

Berdasarkan praktikum yang telah dilakukan, dapat disimpulkan bahwa:

1. gRPC dapat digunakan untuk membangun komunikasi antara aplikasi yang menggunakan bahasa pemrograman berbeda.
2. Go dapat digunakan sebagai server yang menyediakan layanan pengambilan profil pengguna melalui port `50051`.
3. Python dan Node.js dapat digunakan sebagai client untuk mengirim permintaan dan menerima respons dari server Go.
4. File `user.proto` berfungsi sebagai definisi bersama untuk struktur pesan dan layanan yang digunakan oleh server dan client.
5. Penambahan field `phone = 3` memperluas data respons sehingga informasi nomor telepon dapat dikirimkan kepada client.
6. Static Code Generation digunakan pada kode Go dan Python, sedangkan Node.js menggunakan Dynamic Runtime Loading melalui `@grpc/proto-loader`.

<p align="justify">Secara keseluruhan, praktikum ini memberikan pemahaman mengenai komunikasi antarbahasa menggunakan gRPC, penggunaan Protocol Buffers, serta perbedaan pendekatan pemuatan definisi Protobuf pada aplikasi.</p>

# H. REFERENSI

1. gRPC. *Introduction to gRPC*. https://grpc.io/docs/what-is-grpc/introduction/
2. Protocol Buffers. *Protocol Buffers Documentation*. https://protobuf.dev/
3. The Go Authors. *The Go Programming Language Documentation*. https://go.dev/doc/
4. Python Software Foundation. *Python Documentation*. https://docs.python.org/3/
5. OpenJS Foundation. *Node.js Documentation*. https://nodejs.org/docs/
6. gRPC. *gRPC for Go*. https://grpc.io/docs/languages/go/
7. gRPC. *gRPC for Python*. https://grpc.io/docs/languages/python/
8. gRPC. *gRPC for Node.js*. https://grpc.io/docs/languages/node/

---

<div align="center">

**LAPORAN PRAKTIKUM LABORATORIUM POLIGLOT gRPC**

**POLITEKNIK NEGERI LHOKSEUMAWE**

</div>
