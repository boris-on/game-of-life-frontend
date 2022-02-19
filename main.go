package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var game_tpl = template.Must(template.ParseFiles("game.html"))
var login_tpl = template.Must(template.ParseFiles("login.html"))

type PlayerInfo struct {
	Nick string
}

func gamePage(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	nick := r.Form["nick"][0]

	playerInfo := PlayerInfo{Nick: nick}

	buf := &bytes.Buffer{}
	err := game_tpl.Execute(buf, playerInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	buf.WriteTo(w)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	// for k, v := range r.Header {
	// 	log.Println(k, v)
	// }

	buf := &bytes.Buffer{}
	err := login_tpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	buf.WriteTo(w)
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/game", gamePage)
	mux.HandleFunc("/", loginPage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	fmt.Println("starting server at", port)
	http.ListenAndServe(":"+port, mux)
}
