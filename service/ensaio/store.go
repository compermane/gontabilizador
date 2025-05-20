package ensaio

import (
	"database/sql"

	"github.com/compermane/gontabilizador/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateEnsaio(ensaio types.Ensaio) error {
	_, err := s.db.Exec("INSERT INTO ensaio (data, nome) VALUES (?, ?)", 
						ensaio.Data, ensaio.Nome)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetEnsaioByID(id int) (*types.Ensaio, error) {
	return nil, nil
}