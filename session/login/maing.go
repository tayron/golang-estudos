package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/tayron/golang-estudos/session/login/session"
)

var templates *template.Template
var store = sessions.NewCookieStore([]byte("t0op-s3cr3t"))

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/logout", logoutGetHandler).Methods("GET")

	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fs))
	http.Handle("/", r)

	log.Panic(http.ListenAndServe(":8181", nil))
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	username := session.GetSessionData("username", w, r)
	password := session.GetSessionData("password", w, r)

	if username == "" {
		http.Redirect(w, r, "/login", 302)
	}

	parameters := struct {
		Username string
		Password string
	}{
		Username: username,
		Password: password,
	}

	templates.ExecuteTemplate(w, "index.html", parameters)
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	username := session.GetSessionData("username", w, r)

	if username != "" {
		http.Redirect(w, r, "/", 302)
	}

	templates.ExecuteTemplate(w, "login.html", nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	session.SetSessionData("username", username, w, r)
	session.SetSessionData("password", password, w, r)

	http.Redirect(w, r, "/", 302)
}

func logoutGetHandler(w http.ResponseWriter, r *http.Request) {
	session.ClearSessionData(w, r)
	http.Redirect(w, r, "/login", 302)
}
