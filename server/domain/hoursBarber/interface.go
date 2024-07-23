package hoursBarber

import (
	"api/server/infrastructure/persistence/hoursBarber"

	"github.com/gin-gonic/gin"
)

type IHoursBarber interface {
	Create(ctx *gin.Context, hours *hoursBarber.HoursBarber) (err error)
	CheckConflitHoursBarber(ctx *gin.Context, hours *hoursBarber.HoursBarber) (conflit bool, err error)
	List(ctx *gin.Context) (hoursBarbers []hoursBarber.ListHoursBarber, err error)
	CreateHoursBarberException(ctx *gin.Context, hoursException *hoursBarber.HoursBarberException) (err error)
	MarkReservationAsPending(ctx *gin.Context, BarberID *int64, hoursExeptionID *string) (err error)
	HoursExecptionExists(ctx *gin.Context, hoursException *hoursBarber.HoursBarberException) (exists bool, err error)
}
