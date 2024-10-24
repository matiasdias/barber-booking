package barberBook

import (
	"api/server/interface/barberBook/auth"
	barberShop "api/server/interface/barberBook/baberShop"
	"api/server/interface/barberBook/barber"
	"api/server/interface/barberBook/client"
	"api/server/interface/barberBook/hoursBarber"
	"api/server/interface/barberBook/reservation"
	"api/server/interface/barberBook/service"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	reservation.RouterReservation(r.Group("reservation"))
	service.RouterService(r.Group("service"))
	hoursBarber.RouterHoursBarber(r.Group("hoursBarber"))
	hoursBarber.RouterHoursExecption(r.Group("hoursBarberException"))
	client.RouterClient(r.Group("client"))
	barber.RouterBarber(r.Group(""))
	barberShop.RouterBarberShop(r.Group("barberShop"))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func AuhRouter(r *gin.RouterGroup) {
	auth.RouterGoogleAuth(r.Group(""))
	auth.RouterJwtAuth(r.Group(""))
}
