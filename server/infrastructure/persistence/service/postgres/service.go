package postgres

import (
	"api/server/infrastructure/persistence/service"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type PGService struct {
	DB *sql.DB
}

func (pg *PGService) Create(ctx *gin.Context, service *service.Services) (err error) {

	query := "INSERT INTO servico ( nome, preco, duracao ) VALUES ( $1, $2, $3 ) RETURNING id"
	var serviceID int64
	err = pg.DB.QueryRowContext(ctx, query, service.Name, service.Price, service.Duration).Scan(&serviceID)
	if err != nil {
		log.Println("Erro ao inserir e consultar o ID do servico:", err)
		return err
	}
	return nil
}

func (pg *PGService) List(ctx *gin.Context) (services []service.ListService, err error) {

	query := "SELECT id, nome, preco, duracao, data_criacao, data_atualizacao FROM servico"

	rows, err := pg.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Erro ao executar a consulta:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var service service.ListService
		err = rows.Scan(&service.ID, &service.Name, &service.Price, &service.Duration, &service.CriadoEm, &service.UpdateEm)
		if err != nil {
			log.Println("Erro ao listar os consulta:", err)
			return
		}
		services = append(services, service)
	}
	return
}
