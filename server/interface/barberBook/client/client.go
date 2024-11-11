package client

import (
	"api/server/aplication/client"

	"github.com/gin-gonic/gin"
)

// ListUsers godoc
// @Summary List os clientes da barbearia
// @Description Lista todos os clientes da barbearia
// @Tags client
// @Accept  json
// @Produce  json
// @Success 200 {array} ListClients
// @Router /barber/client/list [get]
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
