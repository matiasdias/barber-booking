package barberShop

import (
	"api/server/aplication/barberShop"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var (
		req barberShop.CreateBarberShop
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err = barberShop.BarberShopCreate(c.Copy(), &req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Barber shopadded successfully",
	})
}

func List(c *gin.Context) {
	var (
		err error
	)
	barberShops, err := barberShop.ListBshop(c.Copy())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, barberShops)
}
