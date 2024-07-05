package service

import (
	"api/server/database"
	"api/server/domain/service"
	"errors"
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
		ID:    services.ID,
		Name:  services.Name,
		Price: services.Price,
	}
	if *dados.Name == "" || *dados.Price == 0.0 {
		log.Println("Failed to create service: missing required fields")
		return errors.New("missing required fields")
	}

	if err := server.Create(ctx, dados); err != nil {
		log.Printf("Failed to create service: %v", err)
		return err
	}

	return
}

func List(ctx *gin.Context) (services []service.ListService, err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()

	server := service.GetService(service.GetRepository(db))
	services, err = server.List(ctx)
	if err != nil {
		log.Printf("Failed to list services: %v", err)
		return
	}
	return
}
