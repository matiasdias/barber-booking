package postgres

import (
	"api/server/infrastructure/persistence/client"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type PGClient struct {
	DB *sql.DB
}

func (pg *PGClient) Create(ctx *gin.Context, cliente *client.Client) (err error) {

	query := "INSERT INTO cliente ( nome, email, contato, senha ) VALUES ( $1, $2, $3, $4 ) RETURNING id"
	var clientID int64
	err = pg.DB.QueryRowContext(ctx, query, cliente.Name, cliente.Email, cliente.Contato, cliente.PassWord).Scan(&clientID)
	if err != nil {
		log.Println("Erro ao consultar ID do cliente:", err)
		return
	}
	return nil
}

func (pg *PGClient) List(ctx *gin.Context) (clients []client.Clients, err error) {

	query := "SELECT id, nome, email, contato, senha, criadoem, updatedem FROM cliente order by criadoem ASC"
	rows, err := pg.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Erro ao consultar clientes:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var client client.Clients
		err = rows.Scan(&client.ID, &client.Name, &client.Email, &client.Contato, &client.PassWord, &client.CriadoEm, &client.UpdateEm)
		if err != nil {
			log.Println("Erro ao listar os clientes:", err)
			return
		}
		clients = append(clients, client)
	}
	return
}
