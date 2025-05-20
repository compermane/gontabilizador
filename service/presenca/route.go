package presenca

import (
	"net/http"

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
	var payload types.RegisterPresencaPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	err := h.store.CreatePresenca(types.Presenca{
		IDRitmista: payload.IDRitmista,
		IDEnsaio: payload.IDRitmista,
		Presente: payload.Presente,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
