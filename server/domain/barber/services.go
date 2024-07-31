package barber

import (
	"api/server/infrastructure/persistence/barber"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Service struct {
	repo IBarber
}

// GetService retorna um servico para acesso a funções de auxilio à lógica de negócio
func GetService(r IBarber) *Service {
	return &Service{repo: r}
}

// GetRepository retorna um repositório para acesso à camada de dados
func GetRepository(db *sql.DB) IBarber {
	return newRepository(db)
}

func (s *Service) Create(ctx *gin.Context, barbers *Barber) (err error) {
	dados := &barber.Barber{
		Name:    barbers.Name,
		Contato: barbers.Contato,
	}
	return s.repo.Create(ctx, dados)
}

func (s *Service) List(ctx *gin.Context) (barbers []Barbers, err error) {
	barber, err := s.repo.List(ctx)
	if err != nil {
		return
	}
	barbers = make([]Barbers, len(barber))
	for i := range barber {
		var bar Barbers
		bar.ID = barber[i].ID
		bar.Name = barber[i].Name
		bar.Contato = barber[i].Contato
		bar.CriadoEm = barber[i].CriadoEm
		bar.UpdateEm = barber[i].UpdateEm
		barbers[i] = bar
	}
	return
}
