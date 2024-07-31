package service

type CreateService struct {
	ID       *int64   `conversor:"id" json:"id"`
	Name     *string  `conversor:"nome" json:"nome" binding:"required"`
	Price    *float64 `conversor:"preco" json:"preco" binding:"required"`
	Duration *string  `conversor:"duracao" json:"duracao"`
}
