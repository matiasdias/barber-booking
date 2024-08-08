package service

import "time"

type CreateService struct {
	Name     *string  `conversor:"nome" json:"nome" binding:"required"`
	Price    *float64 `conversor:"preco" json:"preco" binding:"required"`
	Duration *string  `conversor:"duracao" json:"duracao"`
}

type ListServices struct {
	ID       *int64     `json:"id" conversor:"id"`
	Name     *string    `json:"nome" conversor:"nome"`
	Price    *float64   `json:"preco" conversor:"preco"`
	Duration *string    `json:"duracao" conversor:"duracao"`
	CriadoEm *time.Time `json:"data_criacao" conversor:"data_criacao"`
	UpdateEm *time.Time `json:"data_atualizacao" conversor:"data_atualizacao"`
}
