package barber

import "time"

type Barber struct {
	Name    *string `conversor:"nome"`
	Contato *string `conversor:"contato"`
}

type Barbers struct {
	ID       *int64     `json:"id" conversor:"id"`
	Name     *string    `json:"nome" conversor:"nome"`
	Contato  *string    `json:"contato" conversor:"contato"`
	CriadoEm *time.Time `json:"data_criacao" conversor:"data_criacao"`
	UpdateEm *time.Time `json:"data_atualizacao" conversor:"data_atualizacao"`
}
