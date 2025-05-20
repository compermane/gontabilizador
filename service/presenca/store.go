package presenca

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

func (s *Store) CreatePresenca(presenca types.Presenca) error {
	_, err := s.db.Query("INSERT INTO presenca VALUES (ritmista_id, ensaio_id, presente) VALUES (?, ?, ?)",
						 presenca.IDRitmista,
						 presenca.IDEnsaio,
						 presenca.Presente)

	if err != nil {
		return err
	}

	return nil
}