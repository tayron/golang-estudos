package view

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/tayron/go-lang-estudos/paginacao/application/expanders"
)

func LoadViewTemplate(fileName string, w http.ResponseWriter, parameters interface{}) {
	pathToFile := fmt.Sprintf("./pages/%s.html", fileName)
	templateBase := template.New("base.html").Funcs(expanders.All())
	tplHome, error := templateBase.ParseFiles("./pages/base.html", pathToFile)
	template.Must(tplHome, error).Execute(w, parameters)
}
