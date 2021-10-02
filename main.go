package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/sites", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})

	http.ListenAndServe(":8000", nil)
}
