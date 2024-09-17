package hoursBarber

import (
	"api/server/aplication/hoursBarber"

	"github.com/gin-gonic/gin"
)

// CreateException godoc
// @Summary Criação de exceções de hora de trabalho para o barbearia
// @Description Cria uma exceção de hora de trabalho para o barbeiro
// @Tags hoursBarber
// @Accept  json
// @Produce  json
// @Param hoursBarberException body CreateException true "Create hours barber exception"
// @Success 200 "Sem conteúdo"
// @Router /barber/hoursBarberException/create [post]
func CreateException(c *gin.Context) {
	var (
		req hoursBarber.CreateException
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err = hoursBarber.CreateHoursBarberExecption(c.Copy(), &req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Barber hours exception added successfully",
	})
}

func ListException(c *gin.Context) {
	var (
		err error
	)
	hoursBarbers, err := hoursBarber.ListHoursBarberException(c.Copy())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, hoursBarbers)
}
