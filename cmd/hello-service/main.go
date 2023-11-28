package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	var hello_counter uint64 = 0
	r.Route("/hello", func(r chi.Router) {
		r.Get("/{name}", func(w http.ResponseWriter, r *http.Request) {
			name := chi.URLParam(r, "name")
			hello_counter++
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"Hello": name})
		})
		r.Get("/count", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"Hello count": strconv.FormatUint(hello_counter, 10)})
		})
		r.Put("/{delete_count:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			delete_count := chi.URLParam(r, "delete_count")
			delete_count_int, _ := strconv.ParseUint(delete_count, 10, 64)
			if delete_count_int > hello_counter {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			hello_counter -= delete_count_int
		})
	})

	http.ListenAndServe(":8080", r)
}
