package render

import (
	"html/template"
	"net/http"
)

func TmplRender(w http.ResponseWriter, pTmpl string, data interface{}) {

	tmplFiles := []string{
		"public/base.html",
		"public/" + pTmpl,
	}

	tmpl, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
