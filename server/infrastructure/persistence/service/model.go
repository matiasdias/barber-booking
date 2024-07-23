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
	CriadoEm *time.Time `json:"data_criacao" conversor:"data_criacao"`
	UpdateEm *time.Time `json:"data_atualizacao" conversor:"data_atualizacao"`
}
