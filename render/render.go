package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/compermane/gontabilizador/types"
)

type PageHandler struct {
	RitmistaStore types.RitmistaStore
	EnsaioStore   types.EnsaioStore
	PresencaStore types.PresencaStore
}

func NewPageHandler(rs types.RitmistaStore, es types.EnsaioStore, ps types.PresencaStore) *PageHandler {
	return &PageHandler{
		RitmistaStore: rs,
		EnsaioStore: es,
		PresencaStore: ps,
	}
}

func (ph *PageHandler) Home(w http.ResponseWriter, r *http.Request) {
	ritmistas, err := ph.RitmistaStore.GetAllRitmistas()
	if err != nil {
		http.Error(w, "Erro ao buscar ritmistas", http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join("static", "templates", "home.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Println("Erro ao carregar template:", err)
		http.Error(w, "Erro interno no servidor", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, ritmistas)
	if err != nil {
		log.Println("Erro ao renderizar template:", err)
		http.Error(w, "Erro ao renderizar página", http.StatusInternalServerError)
	}
}