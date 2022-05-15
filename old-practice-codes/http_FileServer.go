package main

import "net/http"

func main() {

	http.ListenAndServe(":9595", http.FileServer(http.Dir("public")))
}