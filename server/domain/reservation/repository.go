package reservation

import (
	"api/server/infrastructure/persistence/reservation"
	"api/server/infrastructure/persistence/reservation/postgres"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type repository struct {
	pg *postgres.PGReservation
}

func newRepository(db *sql.DB) *repository {
	return &repository{
		pg: &postgres.PGReservation{DB: db},
	}
}

func (r *repository) Create(ctx *gin.Context, reser *reservation.Reservation) error {
	return r.pg.Create(ctx, reser)
}

func (r *repository) CheckConflictReservation(ctx *gin.Context, reser *reservation.Reservation) error {
	return r.pg.CheckConflictReservation(ctx, reser)
}

func (r *repository) List(ctx *gin.Context) (reservations []reservation.ReservationList, err error) {
	return r.pg.List(ctx)
}

func (r *repository) CheckExceptionForBarber(ctx *gin.Context, barberID *int64, dataReservation *string) (bool, error) {
	return r.pg.CheckExceptionForBarber(ctx, barberID, dataReservation)
}

func (r *repository) UpdateReservation(ctx *gin.Context, reservationID *int64, reser *reservation.Reservation) error {
	return r.pg.UpdateReservation(ctx, reservationID, reser)
}
