package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

// Main function in GO, like C
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welconme, %!", r.URL.Path[1:])
}

type Message struct {
	Text string
}

func about(w http.ResponseWriter, r *http.Request) {
	m := Message{"Teste"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Write(b)
}