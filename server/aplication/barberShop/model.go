package barberShop

type CreateBarberShop struct {
	ID               *int64  `conversor:"id" json:"id"`
	Name             *string `conversor:"nome" json:"nome" binding:"required"`
	Cidade           *string `conversor:"cidade" json:"cidade" binding:"required"`
	Rua              *string `conversor:"rua" json:"rua" binding:"required"`
	NumeroResidencia *int64  `conversor:"numero_residencia" json:"numero_residencia" binding:"required"`
	PontoReferencia  *string `conversor:"ponto_referencia" json:"ponto_referencia" binding:"required"`
	Contato          *string `conversor:"contato" json:"contato" binding:"required"`
}
