package barber

import (
	"api/server/aplication/barber"

	"github.com/gin-gonic/gin"
)

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
