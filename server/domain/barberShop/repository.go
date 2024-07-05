package barberShop

import (
	"api/server/infrastructure/persistence/barberShop"
	"api/server/infrastructure/persistence/barberShop/postgres"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type repository struct {
	pg *postgres.PGBarberShop
}

func newRepository(db *sql.DB) *repository {
	return &repository{
		pg: &postgres.PGBarberShop{DB: db},
	}
}

func (r *repository) Create(ctx *gin.Context, b *barberShop.BarberShop) (err error) {
	return r.pg.Create(ctx, b)
}

func (r *repository) List(ctx *gin.Context) (barberShops []barberShop.ListBarberShop, err error) {
	return r.pg.List(ctx)
}
