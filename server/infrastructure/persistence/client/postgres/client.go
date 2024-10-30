package postgres

import (
	"api/server/infrastructure/persistence/client"
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type PGClient struct {
	DB *sql.DB
}

func (pg *PGClient) Create(ctx *gin.Context, cliente *client.Client) (err error) {

	query := "INSERT INTO cliente ( nome, email, refresh_token ) VALUES ( $1, $2, $3) RETURNING id"
	var clientID int64
	err = pg.DB.QueryRowContext(ctx, query, cliente.Name, cliente.Email, cliente.RefreshToken).Scan(&clientID)
	if err != nil {
		log.Println("Erro ao consultar ID do cliente:", err)
		return
	}
	return nil
}

func (pg *PGClient) List(ctx *gin.Context) (clients []client.Clients, err error) {

	query := "SELECT id, nome, email, data_criacao, data_atualizacao FROM cliente order by data_criacao ASC"
	rows, err := pg.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Erro ao consultar clientes:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var client client.Clients
		err = rows.Scan(&client.ID, &client.Name, &client.Email, &client.CriadoEm, &client.UpdateEm)
		if err != nil {
			log.Println("Erro ao listar os clientes:", err)
			return
		}
		clients = append(clients, client)
	}
	return
}

func (pg *PGClient) FindByEmail(ctx *gin.Context, email *string) (bool, error) {
	query := "SELECT id FROM cliente WHERE email = $1"
	var clientID int64
	err := pg.DB.QueryRowContext(ctx, query, email).Scan(&clientID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		log.Println("Erro ao consultar ID do cliente:", err)
		return false, err
	}
	return true, nil
}

func (pg *PGClient) UpdateRefreshToken(ctx context.Context, email *string, refreshToken *string) error {
	query := `
		UPDATE cliente 
		SET refresh_token = $1 
		WHERE id = (SELECT id FROM cliente WHERE email = $2) 
		RETURNING id
	`

	var id int

	err := pg.DB.QueryRowContext(ctx, query, refreshToken, email).Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to update refresh token: %v", err)
	}

	return nil
}

func (pg *PGClient) GetRefreshTokenByEmail(ctx *gin.Context, email *string) (string, error) {
	query := "SELECT refresh_token FROM cliente WHERE email = $1"
	var refreshToken string
	err := pg.DB.QueryRowContext(ctx, query, email).Scan(&refreshToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		log.Println("Erro ao consultar ID do cliente:", err)
		return "", err
	}
	return refreshToken, nil
}
