package barberShop

import "time"

type BarberShop struct {
	ID               *int64  `conversor:"id"`
	Name             *string `conversor:"nome" `
	Cidade           *string `conversor:"cidade"`
	Rua              *string `conversor:"rua" json:"rua"`
	NumeroResidencia *int64  `conversor:"numero_residencia"`
	PontoReferencia  *string `conversor:"ponto_referencia"`
	Contato          *string `conversor:"contato"`
}
type ListBarberShop struct {
	ID               *int64     `json:"id" conversor:"id"`
	Name             *string    `json:"nome" conversor:"nome"`
	Cidade           *string    `json:"cidade" conversor:"cidade"`
	Rua              *string    `json:"rua" conversor:"rua"`
	NumeroResidencia *int64     `json:"numero_residencia" conversor:"numero_residencia"`
	PontoReferencia  *string    `json:"ponto_referencia" conversor:"ponto_referencia"`
	Contato          *string    `json:"contato" conversor:"contato"`
	CriadoEm         *time.Time `json:"criado_em" conversor:"criado_em"`
	UpdateEm         *time.Time `json:"updated_em" conversor:"updated_em"`
}
