package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
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

	http.ListenAndServe(":8080", nil)
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session")
	username, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/login", 302)
	}

	templates.ExecuteTemplate(w, "index.html", username)
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session")
	username, _ := session.Values["username"]

	if username != nil {
		http.Redirect(w, r, "/", 302)
	}

	templates.ExecuteTemplate(w, "login.html", username)

}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session, _ := store.Get(r, "session")

	username := r.PostForm.Get("username")
	session.Values["username"] = username

	password := r.PostForm.Get("password")
	session.Values["password"] = password

	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func logoutGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}
