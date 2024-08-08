package barberShop

import (
	"api/server/aplication/barberShop"

	"github.com/gin-gonic/gin"
)

// CreateBarberShop godoc
// @Summary Criação das barbearias
// @Description Cria um estavelecimento
// @Tags barberShop
// @Accept  json
// @Produce  json
// @Param barberShop body CreateBarberShop true "Create barber shop"
// @Success 200 "Sem conteúdo"
// @Router /barber/barberShop/create [post]
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

// ListBarbserShop godoc
// @Summary Lista todas as barbearias
// @Description Lista todas as barbearias disponíveis para serem feitas as reservas
// @Tags barberShop
// @Accept  json
// @Produce  json
// @Success 200 {array} ListBarbserShop
// @Router /barber/barberShop/list [get]
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
