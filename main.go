package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("main.html"))

func mainPage(w http.ResponseWriter, r *http.Request) {

	for k, v := range r.Header {
		log.Println(k, v)
	}

	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	buf.WriteTo(w)
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", mainPage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":"+port, mux)
}
