package hoursBarber

import (
	"api/server/aplication/hoursBarber"

	"github.com/gin-gonic/gin"
)

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
