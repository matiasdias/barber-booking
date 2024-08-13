package barber

import "time"

type CreateBarber struct {
	Name    *string `conversor:"nome" json:"nome" binding:"required"`
	Contato *string `conversor:"contato" json:"contato" binding:"required"`
}

type ListBarbers struct {
	ID       *int64     `json:"id" conversor:"id"`
	Name     *string    `json:"nome" conversor:"nome"`
	Contato  *string    `json:"contato" conversor:"contato"`
	CriadoEm *time.Time `json:"data_criacao" conversor:"data_criacao"`
	UpdateEm *time.Time `json:"data_atualizacao" conversor:"data_atualizacao"`
}
