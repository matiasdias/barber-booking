package service

import (
	"api/server/aplication/service"
	"api/server/utils"
	"net/http"

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
	var req service.CreateService

	if err := c.ShouldBindJSON(&req); err != nil {
		customErr := utils.New(400, "missing or invalid fields", err)
		c.Set("error", customErr)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chama a função de criação de serviço
	if err := service.Create(c.Copy(), &req); err != nil {
		// Cria um erro personalizado
		customErr := utils.New(400, "service creation failed", err)
		// Adiciona o erro ao contexto
		c.Set("error", customErr)
		// Retorna a resposta
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
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
