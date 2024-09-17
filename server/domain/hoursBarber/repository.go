package hoursBarber

import (
	"api/server/infrastructure/persistence/hoursBarber"
	"api/server/infrastructure/persistence/hoursBarber/postgres"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type repository struct {
	pg         *postgres.PGHoursBarber
	pgExcption *postgres.PGHoursBarberException
}

func newRepository(db *sql.DB) *repository {
	return &repository{
		pg:         &postgres.PGHoursBarber{DB: db},
		pgExcption: &postgres.PGHoursBarberException{DB: db},
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

func (r *repository) CreateHoursBarberException(ctx *gin.Context, hoursException *hoursBarber.HoursBarberException) error {
	return r.pgExcption.CreateHoursBarberException(ctx, hoursException)
}

func (r *repository) MarkReservationAsPending(ctx *gin.Context, BarberID *int64, hoursExeptionID *string) (marked bool, err error) {
	return r.pgExcption.MarkReservationAsPending(ctx, BarberID, hoursExeptionID)
}

func (r *repository) HoursExecptionExists(ctx *gin.Context, hoursException *hoursBarber.HoursBarberException) (exists bool, err error) {
	return r.pgExcption.HoursExecptionExists(ctx, hoursException)
}

func (r *repository) ListExeption(ctx *gin.Context) (hoursBarbers []hoursBarber.ListHoursBarberExeption, err error) {
	return r.pgExcption.ListExeption(ctx)
}
