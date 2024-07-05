package client

type CreateClient struct {
	ID       *int64  `conversor:"id" json:"id"`
	Name     *string `conversor:"nome" json:"nome" binding:"required"`
	Email    *string `conversor:"email" json:"email" binding:"required"`
	Contato  *string `conversor:"contato" json:"contato" binding:"required"`
	PassWord *string `conversor:"senha" json:"senha" binding:"required"`
}
