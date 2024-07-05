package barberShop

import (
	"api/server/infrastructure/persistence/barberShop"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Service struct {
	repo IBarberShop
}

// GetService retorna um servico para acesso a funções de auxilio à lógica de negócio
func GetService(r IBarberShop) *Service {
	return &Service{repo: r}
}

// GetRepository retorna um repositório para acesso à camada de dados
func GetRepository(db *sql.DB) IBarberShop {
	return newRepository(db)
}

func (s *Service) Create(ctx *gin.Context, b *BarberShop) (err error) {
	dados := &barberShop.BarberShop{
		ID:               b.ID,
		Name:             b.Name,
		Cidade:           b.Cidade,
		Rua:              b.Rua,
		NumeroResidencia: b.NumeroResidencia,
		PontoReferencia:  b.PontoReferencia,
		Contato:          b.Contato,
	}

	return s.repo.Create(ctx, dados)
}

func (s *Service) List(ctx *gin.Context) (barberShops []ListBarberShop, err error) {
	barberShop, err := s.repo.List(ctx)
	if err != nil {
		return
	}
	barberShops = make([]ListBarberShop, len(barberShop))
	for i := range barberShop {
		var shop ListBarberShop
		shop.ID = barberShop[i].ID
		shop.Name = barberShop[i].Name
		shop.Cidade = barberShop[i].Cidade
		shop.Rua = barberShop[i].Rua
		shop.NumeroResidencia = barberShop[i].NumeroResidencia
		shop.PontoReferencia = barberShop[i].PontoReferencia
		shop.Contato = barberShop[i].Contato
		shop.CriadoEm = barberShop[i].CriadoEm
		shop.UpdateEm = barberShop[i].UpdateEm
		barberShops[i] = shop
	}
	return
}
