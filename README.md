
# Go Kasir API 🧾

Simple REST API menggunakan **Go (Golang)** untuk simulasi sistem kasir.
Project ini dibuat **tanpa framework**, hanya menggunakan **standard library `net/http`** agar mudah dipahami oleh pemula.

> Cocok untuk belajar:
- HTTP server di Go
- REST API dasar
- CRUD (Create, Read, Update, Delete)
- Parsing URL & JSON
- Testing API dengan `curl`

---

## 🚀 Fitur

### Produk
- GET `/api/produk` → Ambil semua produk
- POST `/api/produk` → Tambah produk
- GET `/api/produk/{id}` → Detail produk
- PUT `/api/produk/{id}` → Update produk
- DELETE `/api/produk/{id}` → Hapus produk

### Category
- GET `/categories` → Ambil semua kategori
- POST `/categories` → Tambah kategori
- GET `/categories/{id}` → Detail kategori
- PUT `/categories/{id}` → Update kategori
- DELETE `/categories/{id}` → Hapus kategori

### Health Check
- GET `/health` → Cek status API

---

## 🛠️ Teknologi
- Go ≥ 1.20
- `net/http`
- `encoding/json`
- In-memory data (tanpa database)

---

## 📂 Struktur Project
go-kasir/
├── main.go
├── go.mod
└── README.md


---

## ▶️ Cara Menjalankan

### 1. Clone repository

git clone https://github.com/asatriana/go-kasir.git
cd go-kasir

### 2. Jalankan server
go run main.go

Jika berhasil, akan muncul: 
Server started on :8080
http://localhost:8080


### 3. Testing API (menggunakan curl)

Health Check
```bash
curl http://localhost:8080/health

{
  "status": "OK",
  "message": "API Running"
}
```

### a.Produk API

Ambil Semua Produk
```bash
curl http://localhost:8080/api/produk
```

Tambah Produk
```bash
curl -X POST http://localhost:8080/api/produk \
  -H "Content-Type: application/json" \
  -d '{"nama":"Teh Manis","harga":5000,"stok":15}'
```

Ambil Produk by ID
```bash
curl http://localhost:8080/api/produk/1
```

Update Produk
```bash
curl -X PUT http://localhost:8080/api/produk/1 \
  -H "Content-Type: application/json" \
  -d '{"nama":"Indomie Goreng","harga":4000,"stok":8}'
```

Hapus Produk
```bash
curl -X DELETE http://localhost:8080/api/produk/1
```

### b. Category API

Ambil Semua Category
```bash
curl http://localhost:8080/categories
```
Tambah Category
```bash
curl -X POST http://localhost:8080/categories \
  -H "Content-Type: application/json" \
  -d '{"name":"Snack","description":"Kategori snack"}'
```

Ambil Category by ID
```bash
curl http://localhost:8080/categories/1
```

Update Category
```bash
curl -X PUT http://localhost:8080/categories/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Makanan Berat","description":"Kategori makanan berat"}'
```
Hapus Category
```bash
curl -X DELETE http://localhost:8080/categories/1
```

### ⚠️ Catatan Penting
Data tidak disimpan permanen
Data akan hilang saat server restart
Tidak menggunakan database
Tidak menggunakan framework
ID masih auto-increment sederhana (berdasarkan panjang slice)

### 🧭 Rencana Pengembangan (Opsional)
 Validasi input (required field)
 Penyimpanan data ke file / database
 Middleware logging
 Authentication & Authorization
 Refactor ke handler/service (clean architecture)
 Pagination & filtering

 ### 👨‍💻 Author
Satriana Aries
Learning & experimenting with Go backend fundamentals.