package service

import (
	"api/server/database"
	"api/server/domain/service"
	"log"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context, services *CreateService) (err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}
	defer db.Close()

	server := service.GetService(service.GetRepository(db))

	dados := &service.Services{
		Name:     services.Name,
		Price:    services.Price,
		Duration: services.Duration,
	}

	formatDuration, err := server.ValidadeService(dados)
	if err != nil {
		log.Printf("Failed to validate service: %v", err)
		return err
	}

	dados.Duration = formatDuration.Duration

	if err := server.Create(ctx, dados); err != nil {
		log.Printf("Failed to create service: %v", err)
		return err
	}

	return
}

func List(ctx *gin.Context) (services []*ListServices, err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()

	var (
		server = service.GetService(service.GetRepository(db))
		dados  []service.ListService
	)

	if dados, err = server.List(ctx); err != nil {
		log.Printf("Failed to list services: %v", err)
		return
	}

	for i := range dados {
		ser := &ListServices{
			ID:       dados[i].ID,
			Name:     dados[i].Name,
			Price:    dados[i].Price,
			Duration: dados[i].Duration,
			CriadoEm: dados[i].CriadoEm,
			UpdateEm: dados[i].UpdateEm,
		}
		services = append(services, ser)
	}

	return
}
