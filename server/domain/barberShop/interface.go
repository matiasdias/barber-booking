package barberShop

import (
	"api/server/infrastructure/persistence/barberShop"

	"github.com/gin-gonic/gin"
)

type IBarberShop interface {
	Create(ctx *gin.Context, b *barberShop.BarberShop) (err error)
	List(ctx *gin.Context) (barberShops []barberShop.ListBarberShop, err error)
}
