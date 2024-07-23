package barberBook

import (
	barberShop "api/server/interface/barberBook/baberShop"
	"api/server/interface/barberBook/barber"
	"api/server/interface/barberBook/client"
	"api/server/interface/barberBook/hoursBarber"
	"api/server/interface/barberBook/reservation"
	"api/server/interface/barberBook/service"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {

	r.POST("client/create", client.Create)
	r.POST("barber/create", barber.Create)
	r.POST("service/create", service.Create)
	r.POST("barberShop/create", barberShop.Create)
	r.POST("hoursBarber/create", hoursBarber.Create)
	r.POST("reservation/create", reservation.Create)
	r.POST("hoursBarberException/create", hoursBarber.CreateException)

	r.GET("client/list", client.List)
	r.GET("barber/list", barber.List)
	r.GET("service/list", service.List)
	r.GET("barberShop/list", barberShop.List)
	r.GET("hoursBarber/list", hoursBarber.List)
	r.GET("reservation/list", reservation.List)
}
