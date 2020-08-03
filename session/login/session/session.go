package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("t0op-s3cr3t"))

// SetSessionData -
func SetSessionData(key string, value string, w http.ResponseWriter, r *http.Request) {
	sessionStore, _ := store.Get(r, "session")
	sessionStore.Values[key] = value
	sessionStore.Save(r, w)
}

// GetSessionData -
func GetSessionData(key string, w http.ResponseWriter, r *http.Request) string {
	sessionStore, err := store.Get(r, "session")

	if err != nil {
		return ""
	}

	dado, err2 := sessionStore.Values[key]

	if err2 == false {
		return ""
	}
	//return fmt.Sprintf("%s", dado)
	return dado.(string)
}

// ClearSessionData -
func ClearSessionData(w http.ResponseWriter, r *http.Request) {
	sessionStore, _ := store.Get(r, "session")
	sessionStore.Options.MaxAge = -1
	sessionStore.Save(r, w)
}
