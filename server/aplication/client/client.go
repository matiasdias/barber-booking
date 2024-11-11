package client

import (
	"api/server/database"
	"api/server/domain/client"
	"api/server/token"
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var ErrCLientExists = errors.New("client already exists")

func CreateClientFromGoogle(ctx *gin.Context, userInfo *CreateClient) error {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}
	defer db.Close()

	var (
		service = client.GetService(client.GetRepository(db))
		exists  bool
	)
	exists, err = service.FindByEmail(ctx, userInfo.Email)
	if err != nil {
		log.Printf("Failed to check if client exists: %v", err)
		return err
	}

	if exists {
		log.Println("Cliente já cadastrado, atualizando o refresh token...")

		refreshTokenUpdate, err := token.GenerateRefreshToken(userInfo.Email)
		if err != nil {
			log.Printf("Failed to generate refresh token: %v", err)
			return err
		}
		err = service.UpdateRefreshToken(ctx, userInfo.Email, &refreshTokenUpdate)
		if err != nil {
			log.Printf("Failed to update refresh token: %v", err)
			return err
		}
		return ErrCLientExists
	} else {
		if userInfo.Name == nil || *userInfo.Name == "" || userInfo.Email == nil || *userInfo.Email == "" {
			log.Println("Failed to create client: missing required fields")
			return errors.New("missing required fields")
		}
		refreshToken, err := token.GenerateRefreshToken(userInfo.Email)
		if err != nil {
			log.Printf("Failed to generate refresh token: %v", err)
			return err
		}
		dados := &client.Client{
			Name:         userInfo.Name,
			Email:        userInfo.Email,
			RefreshToken: &refreshToken,
		}
		if err = service.Create(ctx, dados); err != nil {
			log.Printf("Failed to create client: %v", err)
			return err
		}
	}

	return nil
}

func LisClient(ctx *gin.Context) (clients []*ListClients, err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}
	defer db.Close()
	var (
		service = client.GetService(client.GetRepository(db))
		dados   []client.Clients
	)

	if dados, err = service.List(ctx); err != nil {
		log.Printf("Failed to list clients: %v", err)
		return nil, err
	}

	for i := range dados {
		cli := &ListClients{
			ID:       dados[i].ID,
			Name:     dados[i].Name,
			Email:    dados[i].Email,
			CriadoEm: dados[i].CriadoEm,
			UpdateEm: dados[i].UpdateEm,
		}
		clients = append(clients, cli)
	}
	return
}

func IsRefreshTokenValid(ctx *gin.Context, refreshToken string) (string, error) {
	claims, err := token.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}
	email := claims.Email

	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return "", err
	}
	defer db.Close()

	var (
		service = client.GetService(client.GetRepository(db))
	)

	storedToken, err := service.GetRefreshTokenByEmail(ctx, &email)
	if err != nil || storedToken != refreshToken {
		return "", fmt.Errorf("invalid or expired token")
	}
	return email, nil
}
