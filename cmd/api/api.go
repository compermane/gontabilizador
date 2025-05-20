package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/compermane/gontabilizador/service/ritmista"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr,
		db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	
	ritmistaStore := ritmista.NewStore(s.db)
	ritmistaHandler := ritmista.NewHandler(ritmistaStore)
	ritmistaHandler.PresencaRoutes(subrouter)

	log.Println("[APIServer] Listening on ", s.addr)
	return http.ListenAndServe(s.addr, router)
}