package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t := template.New("template5.1.html")
	t2, _ := t.ParseFiles("template5.1.html")
	t2.Execute(w, "hello world2")
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
