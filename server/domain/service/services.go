package service

import (
	"api/server/infrastructure/persistence/service"
	"api/server/utils"
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

type Service struct {
	repo IService
}

// GetService retorna um servico para acesso a funções de auxilio à lógica de negócio
func GetService(r IService) *Service {
	return &Service{repo: r}
}

// GetRepository retorna um repositório para acesso à camada de dados
func GetRepository(db *sql.DB) IService {
	return newRepository(db)
}

func (s *Service) Create(ctx *gin.Context, services *Services) (err error) {
	dados := &service.Services{
		ID:       services.ID,
		Name:     services.Name,
		Price:    services.Price,
		Duration: services.Duration,
	}
	return s.repo.Create(ctx, dados)
}

func (s *Service) List(ctx *gin.Context) (services []ListService, err error) {
	servicee, err := s.repo.List(ctx)
	if err != nil {
		return
	}
	services = make([]ListService, len(servicee))
	for i := range servicee {
		var s ListService
		s.ID = servicee[i].ID
		s.Name = servicee[i].Name
		s.Price = servicee[i].Price
		s.Duration = servicee[i].Duration
		s.CriadoEm = servicee[i].CriadoEm
		s.UpdateEm = servicee[i].UpdateEm
		services[i] = s
	}
	return
}

func (s *Service) ValidadeService(service *Services) (*FormartDuration, error) {
	if service.Name == nil || *service.Name == "" {
		return nil, errors.New("missing or invalid Name")
	}

	if service.Price == nil || *service.Price == 0.0 {
		return nil, errors.New("missing or invalid Price")
	}
	duration, err := utils.ParseDuration(service.Duration)
	if err != nil {
		return nil, err
	}

	durationStrFormatted := duration.String()

	format := &FormartDuration{
		Duration: &durationStrFormatted,
	}

	return format, nil

}
