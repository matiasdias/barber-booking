package client

import "time"

type Client struct {
	ID       *int64  `conversor:"id"`
	Name     *string `conversor:"nome"`
	Email    *string `conversor:"email"`
	Contato  *string `conversor:"contato"`
	PassWord *string `conversor:"senha"`
}

type Clients struct {
	ID       *int64     `json:"id" conversor:"id"`
	Name     *string    `json:"nome" conversor:"nome"`
	Email    *string    `json:"email" conversor:"email"`
	Contato  *string    `json:"contato" conversor:"contato"`
	PassWord *string    `json:"senha" conversor:"senha"`
	CriadoEm *time.Time `json:"criado_em" conversor:"criado_em"`
	UpdateEm *time.Time `json:"updated_em" conversor:"updated_em"`
}
