package barber

type CreateBarber struct {
	ID      *int64  `conversor:"id" json:"id"`
	Name    *string `conversor:"nome" json:"nome" binding:"required"`
	Contato *string `conversor:"contato" json:"contato" binding:"required"`
}
