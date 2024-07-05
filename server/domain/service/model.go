package service

import "time"

type Services struct {
	ID    *int64   `conversor:"id"`
	Name  *string  `conversor:"nome"`
	Price *float64 `conversor:"preco"`
}

type ListService struct {
	ID       *int64     `json:"id" conversor:"id"`
	Name     *string    `json:"nome" conversor:"nome"`
	Price    *float64   `json:"preco" conversor:"preco"`
	CriadoEm *time.Time `json:"criado_em" conversor:"criado_em"`
	UpdateEm *time.Time `json:"updated_em" conversor:"updated_em"`
}
