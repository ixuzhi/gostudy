package main

import (
	"fmt"
	"net/http"
)

func header(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	for k, v := range h {
		fmt.Println(k, v)
	}
	fmt.Fprintln(w, h)
}
func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/headers", header)
	server.ListenAndServe()
}
