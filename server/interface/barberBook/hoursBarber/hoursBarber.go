package hoursBarber

import (
	"api/server/aplication/hoursBarber"

	"github.com/gin-gonic/gin"
)

// CreateHoursBarber godoc
// @Summary Criação dos horários de trabalho para o barbeiro
// @Description Cria os horários de trabalho para o barbeiro
// @Tags hoursBarber
// @Accept  json
// @Produce  json
// @Param hoursBarberException body CreateHoursBarber true "Create hours barber"
// @Success 200 "Sem conteúdo"
// @Router /barber/hoursBarber/create [post]
func Create(c *gin.Context) {
	var (
		req hoursBarber.CreateHoursBarber
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err = hoursBarber.Create(c.Copy(), &req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "barber hours added successfully",
	})
}

// hoursBarber.ListHoursBarber godoc
// @Summary Lista todas os horarios de trabalho
// @Description Lista todos os horários de trabalho para o barbeiro
// @Tags hoursBarber
// @Accept  json
// @Produce  json
// @Success 200 {array} hoursBarber.ListHoursBarber
// @Router /barber/hoursBarber/list [get]
func List(c *gin.Context) {
	var (
		err error
	)
	hoursBarbers, err := hoursBarber.ListHourBarber(c.Copy())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, hoursBarbers)
}
