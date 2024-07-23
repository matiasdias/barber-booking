package postgres

import (
	"api/server/infrastructure/persistence/barberShop"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type PGBarberShop struct {
	DB *sql.DB
}

func (pg *PGBarberShop) Create(ctx *gin.Context, b *barberShop.BarberShop) (err error) {

	query := "INSERT INTO barbearia ( nome, cidade, rua, numero_residencia, ponto_referencia, contato ) VALUES ( $1, $2, $3, $4, $5, $6 ) RETURNING id"
	var barberShopID int64
	err = pg.DB.QueryRowContext(ctx, query, b.Name, b.Cidade, b.Rua, b.NumeroResidencia, b.PontoReferencia, b.Contato).Scan(&barberShopID)
	if err != nil {
		log.Println("Erro ao inserir e consultar o ID da barbearia:", err)
		return
	}
	return

}

func (pg *PGBarberShop) List(ctx *gin.Context) (barberShops []barberShop.ListBarberShop, err error) {

	query := "SELECT id, nome, cidade, rua, numero_residencia, ponto_referencia, contato, data_criacao, data_atualizacao FROM barbearia"

	rows, err := pg.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Erro ao consultar as barbearias:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var barberShop barberShop.ListBarberShop
		err = rows.Scan(
			&barberShop.ID,
			&barberShop.Name,
			&barberShop.Cidade,
			&barberShop.Rua,
			&barberShop.NumeroResidencia,
			&barberShop.PontoReferencia,
			&barberShop.Contato,
			&barberShop.CriadoEm,
			&barberShop.UpdateEm)
		if err != nil {
			log.Println("Erro ao listar os barbeiros:", err)
			return
		}
		barberShops = append(barberShops, barberShop)
	}
	return
}
