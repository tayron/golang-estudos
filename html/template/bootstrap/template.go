package bootstrap

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("pages/*.html"))

// LoadView -
func LoadView(w http.ResponseWriter, view string, parameters interface{}) {
	template.Must(templates.ParseGlob("pages/layout/*.html"))
	err := templates.ExecuteTemplate(w, view, parameters)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
