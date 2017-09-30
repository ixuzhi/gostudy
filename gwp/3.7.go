package main

import "net/http"
import "fmt"

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}
func main() {
	hello := HelloHandler{}
	world := WorldHandler{}
	server := http.Server{
		Addr: ":8080",
	}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	server.ListenAndServe()
}
