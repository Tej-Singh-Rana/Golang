package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Status ok")
	}

	http.HandleFunc("/ready", h1)

	log.Fatal(http.ListenAndServe(":9595", nil))
}
