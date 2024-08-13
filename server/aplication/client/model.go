package client

import "time"

type CreateClient struct {
	Name     *string `conversor:"nome" json:"nome" binding:"required"`
	Email    *string `conversor:"email" json:"email" binding:"required"`
	Contato  *string `conversor:"contato" json:"contato" binding:"required"`
	PassWord *string `conversor:"senha" json:"senha" binding:"required"`
}

type ListClients struct {
	ID       *int64     `json:"id" conversor:"id"`
	Name     *string    `json:"nome" conversor:"nome"`
	Email    *string    `json:"email" conversor:"email"`
	Contato  *string    `json:"contato" conversor:"contato"`
	PassWord *string    `json:"senha" conversor:"senha"`
	CriadoEm *time.Time `json:"data_criacao" conversor:"data_criacao"`
	UpdateEm *time.Time `json:"data_atualizacao" conversor:"data_atualizacao"`
}
