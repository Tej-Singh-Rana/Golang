// http.Handle function

package main

import (
	"fmt"
	"net/http"
)

type myHandler struct {
	salt string
}

// http.handle
func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v testing done.", mh.salt)))
}
func main() {
	http.Handle("/", &myHandler{salt: "Server"})
	http.ListenAndServe(":8585", nil)
}
