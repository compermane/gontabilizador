package render

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	path := filepath.Join("static/templates", tmpl)

	t, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, "Erro ao carregar o template", http.StatusInternalServerError)
		log.Println("Erro no template:", err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		log.Println("Erro ao executar template:", err)
	}
}
func HomeRender(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}