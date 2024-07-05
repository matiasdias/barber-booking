package barber

import (
	"api/server/infrastructure/persistence/barber"

	"github.com/gin-gonic/gin"
)

type IBarber interface {
	Create(ctx *gin.Context, barber *barber.Barber) (err error)
	List(ctx *gin.Context) (barbers []barber.Barbers, err error)
}
