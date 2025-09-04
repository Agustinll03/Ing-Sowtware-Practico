package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	fmt.Println("TP4 backend build OK. Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
