package main

import (
	"net/http"

	"github.com/google/uuid"
)

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "content/index.html")
}
func newGame(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/new", newGame)
	http.ListenAndServe(":3000", nil)
}
