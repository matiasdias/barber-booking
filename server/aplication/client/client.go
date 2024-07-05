package client

import (
	"api/server/database"
	"api/server/domain/client"
	"api/server/utils"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context, cliente *CreateClient) (err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}
	defer db.Close()

	service := client.GetService(client.GetRepository(db))

	dados := &client.Client{
		ID:       cliente.ID,
		Name:     cliente.Name,
		Email:    cliente.Email,
		Contato:  cliente.Contato,
		PassWord: cliente.PassWord,
	}
	if *dados.Name == "" || *dados.Email == "" || *dados.Contato == "" || *dados.PassWord == "" {
		log.Println("Failed to create client: missing required fields")
		return errors.New("missing required fields")
	}
	contato, err := utils.FormatContact(dados.Contato)
	if err != nil {
		log.Printf("Failed to format contact: %v", err)
		return err
	}
	dados.Contato = contato

	if err := service.Create(ctx, dados); err != nil {
		log.Printf("Failed to create client: %v", err)
		return err
	}

	return nil
}

func LisClient(ctx *gin.Context) (clients []client.Clients, err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}
	defer db.Close()

	service := client.GetService(client.GetRepository(db))
	clients, err = service.List(ctx)
	if err != nil {
		log.Printf("Failed to list clients: %v", err)
		return nil, err
	}
	return
}
