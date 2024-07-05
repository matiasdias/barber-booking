package reservation

import (
	"api/server/aplication/reservation"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var (
		err error
		req reservation.CreateReservation
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err = reservation.Create(c.Copy(), &req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Reservation created successfully",
	})

}

func List(c *gin.Context) {
	var (
		err error
	)
	reservations, err := reservation.List(c.Copy())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, reservations)
}
