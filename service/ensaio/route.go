package ensaio

import (
	"net/http"

	"github.com/compermane/gontabilizador/types"
	"github.com/compermane/gontabilizador/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.EnsaioStore
}

func NewHandler(store types.EnsaioStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegistroRoutes(router *mux.Router) {
	router.HandleFunc("/criar-ensaio", h.handleCriarEnsaio).Methods("POST")
}

func (h *Handler) handleCriarEnsaio(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterEnsaioPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	err := h.store.CreateEnsaio(types.Ensaio{
		Data: payload.Data,
		Nome: payload.Nome,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}