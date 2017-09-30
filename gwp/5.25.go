package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html")
	t.ExecuteTemplate(w, "layout", "")
}

func process2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html")
	t.ExecuteTemplate(w, "layout", "aa")
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)
	server.ListenAndServe()
}
