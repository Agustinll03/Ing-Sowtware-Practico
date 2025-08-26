package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// --- conectar a la base ---
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// --- 👇 acá agregás la creación de la tabla ---
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS items (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL
        );
    `)
	if err != nil {
		log.Fatal("error creando tabla:", err)
	}
	log.Println("Tabla 'items' lista ✅")

	// --- endpoints ---
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/api/items", itemsHandler)

	log.Println("Servidor escuchando en puerto", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch r.Method {
	case http.MethodPost:
		var it Item
		if err := json.NewDecoder(r.Body).Decode(&it); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := db.QueryRow(
			"INSERT INTO items(name) VALUES($1) RETURNING id", it.Name,
		).Scan(&it.ID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(it)
		return

	case http.MethodGet:
		rows, err := db.Query("SELECT id, name FROM items ORDER BY id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		all := []Item{} // 👈 array vacío (no nil) → se serializa como []
		for rows.Next() {
			var it Item
			if err := rows.Scan(&it.ID, &it.Name); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			all = append(all, it)
		}
		if rows.Err() != nil {
			http.Error(w, rows.Err().Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(all)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
