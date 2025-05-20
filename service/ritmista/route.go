package ritmista

import (
	"fmt"
	"net/http"

	"github.com/compermane/gontabilizador/types"
	"github.com/compermane/gontabilizador/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.RitmistaStore
}

func NewHandler(store types.RitmistaStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegistroRoutes(router *mux.Router) {
	router.HandleFunc("/registro", h.handleRegistro).Methods("POST")
}


func (h *Handler) handleRegistro(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterRitmistaPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	_, err := h.store.GetRitmistaByName(payload.Nome)

	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Ritmista com nome %v já existe\n", payload.Nome))
	}

	err = h.store.CreateRitmista(types.Ritmista{
		Nome: payload.Nome,
		Modulo: payload.Modulo,
		Naipe: payload.Naipe,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}