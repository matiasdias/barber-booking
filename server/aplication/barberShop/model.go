package barberShop

import "time"

type CreateBarberShop struct {
	Name             *string `conversor:"nome" json:"nome" binding:"required"`
	Cidade           *string `conversor:"cidade" json:"cidade" binding:"required"`
	Rua              *string `conversor:"rua" json:"rua" binding:"required"`
	NumeroResidencia *int64  `conversor:"numero_residencia" json:"numero_residencia" binding:"required"`
	PontoReferencia  *string `conversor:"ponto_referencia" json:"ponto_referencia" binding:"required"`
	Contato          *string `conversor:"contato" json:"contato" binding:"required"`
}

type ListBarbserShop struct {
	ID               *int64     `json:"id" conversor:"id"`
	Name             *string    `json:"nome" conversor:"nome"`
	Cidade           *string    `json:"cidade" conversor:"cidade"`
	Rua              *string    `json:"rua" conversor:"rua"`
	NumeroResidencia *int64     `json:"numero_residencia" conversor:"numero_residencia"`
	PontoReferencia  *string    `json:"ponto_referencia" conversor:"ponto_referencia"`
	Contato          *string    `json:"contato" conversor:"contato"`
	CriadoEm         *time.Time `json:"data_criacao" conversor:"data_criacao"`
	UpdateEm         *time.Time `json:"data_atualizacao" conversor:"data_atualizacao"`
}
