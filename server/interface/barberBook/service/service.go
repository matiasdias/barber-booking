package service

import (
	"api/server/aplication/service"

	"github.com/gin-gonic/gin"
)

// CreateService godoc
// @Summary Criação dos serviços
// @Description Cria um novo serviço para a barbeiro e barbearia
// @Tags service
// @Accept  json
// @Produce  json
// @Param service body CreateService true "Create service"
// @Success 200 "Sem conteúdo"
// @Router /barber/service/create [post]
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

// ListService godoc
// @Summary Lista os serviços da barbearia
// @Description Lista todos os serviços ofertados pela barbearia
// @Tags service
// @Accept  json
// @Produce  json
// @Success 200 {array} ListServices
// @Router /barber/service/list [get]
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
