package client

import "time"

type Client struct {
	Name                  *string `conversor:"nome"`
	Email                 *string `conversor:"email"`
	RefreshToken          *string `conversor:"refresh_token"`
	RefreshTokenExpiresAt *int64  `conversor:"refresh_token_expires_at"`
}

type Clients struct {
	ID       *int64     `json:"id" conversor:"id"`
	Name     *string    `json:"nome" conversor:"nome"`
	Email    *string    `json:"email" conversor:"email"`
	CriadoEm *time.Time `json:"data_criacao" conversor:"data_criacao"`
	UpdateEm *time.Time `json:"data_atualizacao" conversor:"data_atualizacao"`
}
