// http handlefunc testing

package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from windows side"))
	})

	//  port & instance
	http.ListenAndServe(":9999", nil)

}
