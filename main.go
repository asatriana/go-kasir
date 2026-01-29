package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Produk represents a product in the cashier system
type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

/* =========================
   IN-MEMORY DATA
========================= */
// In-memory storage (sementara, nanti ganti database)
var produk = []Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
	{ID: 3, Nama: "kecap", Harga: 12000, Stok: 20},
}

var categories = []Category{
	{ID: 1, Name: "Makanan", Description: "Kategori makanan"},
	{ID: 2, Name: "Minuman", Description: "Kategori minuman"},
}

/*
	=========================
	  PRODUK HANDLERS

=========================
*/
func getProdukByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID dari URL path
	// URL: /api/produk/123 -> ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// Cari produk dengan ID tersebut
	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	// Kalau tidak found
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// PUT localhost:8080/api/produk/{id}
func updateProduk(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// loop produk, cari id, ganti sesuai data dari request
	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// DELETE localhost:8080/api/produk/{id}
func deleteProduk(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// loop produk cari ID, dapet index yang mau dihapus
	for i, p := range produk {
		if p.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			produk = append(produk[:i], produk[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

/* =========================
   CATEGORY HANDLERS
========================= */

// GET /categories/{id}
func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for _, c := range categories {
		if c.ID == id {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(c)
			return
		}
	}

	http.Error(w, "Category belum ada", http.StatusNotFound)
}

// PUT /categories/{id}
func updateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var cUpdate Category
	if err := json.NewDecoder(r.Body).Decode(&cUpdate); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range categories {
		if categories[i].ID == id {
			cUpdate.ID = id
			categories[i] = cUpdate

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(cUpdate)
			return
		}
	}

	http.Error(w, "Category belum ada", http.StatusNotFound)
}

// DELETE /categories/{id}
func deleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for i, c := range categories {
		if c.ID == id {
			categories = append(categories[:i], categories[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]string{"message": "sukses delete"})
			return
		}
	}

	http.Error(w, "Category belum ada", http.StatusNotFound)
}

/*
	=========================
	  MAIN

=========================
*/
func main() {

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	/* ========= PRODUK ========= */

	// GET localhost:8080/api/produk
	// POST localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)

		} else if r.Method == "POST" {
			// baca data dari request
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			// masukkin data ke dalam variable produk
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(produkBaru)
		}
	})

	// GET, PUT, DELETE localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukByID(w, r)
		} else if r.Method == "PUT" {
			updateProduk(w, r)
		} else if r.Method == "DELETE" {
			deleteProduk(w, r)
		}
	})

	/* ========= CATEGORY ========= */

	// GET /categories
	// POST /categories
	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(categories)
			return
		}

		if r.Method == http.MethodPost {
			var cNew Category
			if err := json.NewDecoder(r.Body).Decode(&cNew); err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			cNew.ID = len(categories) + 1
			categories = append(categories, cNew)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(cNew)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	// GET/PUT/DELETE /categories/{id}
	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getCategoryByID(w, r)
			return
		}
		if r.Method == http.MethodPut {
			updateCategory(w, r)
			return
		}
		if r.Method == http.MethodDelete {
			deleteCategory(w, r)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)

}
