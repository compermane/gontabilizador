package types

import "time"

type RitmistaStore interface {
	GetRitmistaByName(nome string) (*Ritmista, error)
	GetRitmistaByID(id int) (*Ritmista, error)
	CreateRitmista(Ritmista) error
}

type Ritmista struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Modulo string `json:"modulo"`
	Naipe  string `json:"naipe"`
}

type RegisterRitmistaPayload struct {
	Nome   string  `json:"nome"`
	Naipe  string  `json:"naipe"`
	Modulo string  `json:"modulo"`
}

type EnsaioStore interface {
	GetEnsaioByID(id int) (*Ensaio, error)
	CreateEnsaio(Ensaio) error
}

type Ensaio struct {
	ID   int       `json:"id"`
	Data time.Time `json:"data"`
	Nome string    `json:"nome"`
}

type RegisterEnsaioPayload struct {
	Data time.Time `json:"data"`
	Nome string    `json:"nome"`
}