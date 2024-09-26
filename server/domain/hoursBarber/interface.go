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
	MarkReservationAsPending(ctx *gin.Context, BarberID *int64, hoursExeptionID *string) (marked bool, err error)
	HoursExecptionExists(ctx *gin.Context, hoursException *hoursBarber.HoursBarberException) (exists bool, err error)
	ListExeption(ctx *gin.Context) (hoursBarbers []hoursBarber.ListHoursBarberExeption, err error)
	DeleteExecption(ctx *gin.Context, execptionID *int64) (err error)
	MarkReservationAsActive(ctx *gin.Context, barberID *int64, dataReservation *string) (marked bool, err error)
	GetBarberIDByException(ctx *gin.Context, exceptionID *int64) (barberID *int64, dataExecao *string, err error)
}
