package service

import (
	"api/server/aplication/service"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var (
		req service.CreateService
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = service.Create(c.Copy(), &req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Service created successfully",
	})

}

func List(c *gin.Context) {
	var (
		err error
	)

	services, err := service.List(c.Copy())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, services)
}
