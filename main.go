package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"Status": "ok"}`))
		w.Header().Set("application/json", "Content-Type")
		http.MaxBytesReader()
	})
	server := &http.Server{Addr: ":8080", Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	server.ListenAndServe()
}

const maxBodyBytes = 1024 * 1024

func readJson(w http.ResponseWriter, r *http.Request, output any) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(output); err != nil {
	}
}
