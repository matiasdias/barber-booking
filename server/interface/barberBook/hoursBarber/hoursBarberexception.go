package hoursBarber

import (
	"api/server/aplication/hoursBarber"

	"github.com/gin-gonic/gin"
)

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
