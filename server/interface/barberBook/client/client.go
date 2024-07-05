package client

import (
	"api/server/aplication/client"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var (
		req client.CreateClient
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = client.Create(c.Copy(), &req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Client added successfully"})
}

func List(c *gin.Context) {
	var (
		err error
	)
	clients, err := client.LisClient(c.Copy())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, clients)
}
