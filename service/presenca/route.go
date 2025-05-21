package presenca

import (
	"log"
	"net/http"
	"strconv"

	"github.com/compermane/gontabilizador/types"
	"github.com/compermane/gontabilizador/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.PresencaStore
}

func NewHandler(store types.PresencaStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) PresencaRoutes(router *mux.Router) {
	router.HandleFunc("/registrar-presenca", h.handlePresenca).Methods("POST")
}

func (h *Handler) handlePresenca(w http.ResponseWriter, r *http.Request) {
	log.Println("[handlePresenca] executed")

	var payload types.RegisterPresencaPayload
	if err := r.ParseForm(); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	id_ensaio, _ := strconv.Atoi(r.FormValue("ensaio_id"))
	id_ritmistas := r.Form["presentes"]

	for _, id := range id_ritmistas {
		id_int, _ := strconv.Atoi(id)
		log.Printf("[handlePresenca] %v %v %v %v\n", r.Form["ensaio_id"], id_ensaio, id, id_int)
		payload = types.RegisterPresencaPayload{
			IDEnsaio: id_ensaio,
			IDRitmista: id_int,
			Presente: true,
		}

		err := h.store.CreatePresenca(types.Presenca{
			IDRitmista: payload.IDRitmista,
			IDEnsaio: payload.IDEnsaio,
			Presente: payload.Presente,
		})

		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
	}

	log.Printf("[handlePresenca] succesfully executed on ensaio %v and ritmista %v\n", payload.IDEnsaio, payload.IDRitmista)
	http.Redirect(w, r, "/presencas", http.StatusSeeOther)
}
