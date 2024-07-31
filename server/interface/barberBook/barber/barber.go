package barber

import (
	"api/server/aplication/barber"

	"github.com/gin-gonic/gin"
)

// CreateBarber godoc
// @Summary Criação dos barbeiros
// @Description Cria um novo barbeiro
// @Tags barber
// @Accept  json
// @Produce  json
// @Param barber body CreateBarber true "Create barber"
// @Success 200 "Sem conteúdo"
// @Router /barber/create [post]
func Create(c *gin.Context) {

	var (
		req barber.CreateBarber
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = barber.BarberCreate(c.Copy(), &req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "barber added successfully"})
}

// ListBarber godoc
// @Summary Lista os barbeiros da barbearia
// @Description Lista todos os barbeiros da barbearia
// @Tags barber
// @Accept  json
// @Produce  json
// @Success 200 {array} ListBarbers
// @Router /barber/list [get]
func List(c *gin.Context) {
	var (
		err error
	)

	barbers, err := barber.ListBarber(c.Copy())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, barbers)
}
