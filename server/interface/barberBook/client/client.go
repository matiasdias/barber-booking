package client

import (
	"api/server/aplication/client"

	"github.com/gin-gonic/gin"
)

// CreateClient godoc
// @Summary Create client
// @Description Cria um novo cliente
// @Tags client
// @Accept  json
// @Produce  json
// @Param client body CreateClient true "Create client"
// @Success 200 "Sem conteúdo"
// @Router /barber/client/create [post]
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

// ListUsers godoc
// @Summary List os clientes da barbearia
// @Description Lista todos os clientes da barbearia
// @Tags client
// @Accept  json
// @Produce  json
// @Success 200 {array} client.ListClients
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
