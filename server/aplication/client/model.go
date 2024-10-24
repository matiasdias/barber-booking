package client

import "time"

type CreateClient struct {
	Name  *string `conversor:"nome" json:"name"`
	Email *string `conversor:"email" json:"email"`
}

type ListClients struct {
	ID       *int64     `json:"id" conversor:"id"`
	Name     *string    `json:"nome" conversor:"nome"`
	Email    *string    `json:"email" conversor:"email"`
	PassWord *string    `json:"senha" conversor:"senha"`
	CriadoEm *time.Time `json:"data_criacao" conversor:"data_criacao"`
	UpdateEm *time.Time `json:"data_atualizacao" conversor:"data_atualizacao"`
}
