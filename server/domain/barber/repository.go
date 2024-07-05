package barber

import (
	"api/server/infrastructure/persistence/barber"
	"api/server/infrastructure/persistence/barber/postgres"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type repository struct {
	pg *postgres.PGBarber
}

func newRepository(db *sql.DB) *repository {
	return &repository{
		pg: &postgres.PGBarber{DB: db},
	}
}

func (r *repository) Create(ctx *gin.Context, b *barber.Barber) (err error) {
	return r.pg.Create(ctx, b)
}

func (r *repository) List(ctx *gin.Context) (barbers []barber.Barbers, err error) {
	return r.pg.List(ctx)
}
