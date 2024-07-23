package postgres

import (
	"api/server/infrastructure/persistence/barber"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type PGBarber struct {
	DB *sql.DB
}

func (pg *PGBarber) Create(ctx *gin.Context, barber *barber.Barber) error {

	query := "INSERT INTO barbeiro ( nome, contato ) VALUES ( $1, $2 ) RETURNING id"
	var barberID int64
	err := pg.DB.QueryRowContext(ctx, query, barber.Name, barber.Contato).Scan(&barberID)
	if err != nil {
		log.Println("Erro ao inserir e consultar o ID do barbeiro:", err)
		return err
	}
	return nil
}

func (pg *PGBarber) List(ctx *gin.Context) (barbers []barber.Barbers, err error) {

	query := "SELECT id, nome, contato, data_criacao, data_atualizacao FROM barbeiro ORDER BY data_criacao ASC"
	rows, err := pg.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Erro ao consultar clientes:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var barber barber.Barbers
		err = rows.Scan(&barber.ID, &barber.Name, &barber.Contato, &barber.CriadoEm, &barber.UpdateEm)
		if err != nil {
			log.Println("Erro ao listar os barbeiros:", err)
			return
		}
		barbers = append(barbers, barber)
	}
	return
}
