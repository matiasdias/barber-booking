package hoursBarber

import (
	"api/server/infrastructure/persistence/hoursBarber"

	"github.com/gin-gonic/gin"
)

type IHoursBarber interface {
	Create(ctx *gin.Context, hours *hoursBarber.HoursBarber) (err error)
	CheckConflitHoursBarber(ctx *gin.Context, hours *hoursBarber.HoursBarber) (conflit bool, err error)
	List(ctx *gin.Context) (hoursBarbers []hoursBarber.ListHoursBarber, err error)
}
