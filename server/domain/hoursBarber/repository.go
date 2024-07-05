package hoursBarber

import (
	"api/server/infrastructure/persistence/hoursBarber"
	"api/server/infrastructure/persistence/hoursBarber/postgres"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type repository struct {
	pg *postgres.PGHoursBarber
}

func newRepository(db *sql.DB) *repository {
	return &repository{
		pg: &postgres.PGHoursBarber{DB: db},
	}
}

func (r *repository) Create(ctx *gin.Context, hours *hoursBarber.HoursBarber) error {
	return r.pg.Create(ctx, hours)
}

func (r *repository) CheckConflitHoursBarber(ctx *gin.Context, hours *hoursBarber.HoursBarber) (conflti bool, err error) {
	return r.pg.CheckConflitHoursBarber(ctx, hours)
}

func (r *repository) List(ctx *gin.Context) (hoursBarbers []hoursBarber.ListHoursBarber, err error) {
	return r.pg.List(ctx)
}
