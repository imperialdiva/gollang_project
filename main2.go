package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	var users = map[string]int{
		"Tom":   1,
		"Anna":  2,
		"Fraef": 3,
	}

	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		// if err := json.NewDecoder(r.Body).Decode(users); err != nil{
		// 	fmt.Println("Failed from users get", http.StatusInternalServerError)
		// 	return
		// }
		if err := json.NewEncoder(w).Encode(users); err != nil {
			fmt.Println("Failed", http.StatusInternalServerError)
		}
	})
	/* 	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Println("slash") */
	/* }) */
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	fmt.Println("Starting server at: %s", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to start server: %s", srv.Addr)
	}
}
