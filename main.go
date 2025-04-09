package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users := []string{"Alice", "Bob", "Charlie"}
		json.NewEncoder(w).Encode(users)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var creds struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		json.NewDecoder(r.Body).Decode(&creds)
		if creds.Username == "admin" && creds.Password == "password123" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"token": "dummy-token"}`))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		products := []struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Price int    `json:"price"`
		}{
			{1, "Product A", 100},
			{2, "Product B", 200},
		}
		json.NewEncoder(w).Encode(products)
	})

	fmt.Println("Service berjalan pada port :5000")
	http.ListenAndServe(":5000", nil)
}
