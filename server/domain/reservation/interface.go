package reservation

import (
	"api/server/infrastructure/persistence/reservation"

	"github.com/gin-gonic/gin"
)

type IReservation interface {
	Create(ctx *gin.Context, reser *reservation.Reservation) (err error)
	CheckConflictReservation(ctx *gin.Context, reser *reservation.Reservation) (err error)
	List(ctx *gin.Context) (reservations []reservation.ReservationList, err error)
	CheckExceptionForBarber(ctx *gin.Context, barberID *int64, dataReservation *string) (existe bool, err error)
	UpdateReservation(ctx *gin.Context, reservationID *int64, reser *reservation.Reservation) (err error)
}
