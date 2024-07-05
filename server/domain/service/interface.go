package service

import (
	"api/server/infrastructure/persistence/service"

	"github.com/gin-gonic/gin"
)

type IService interface {
	Create(ctx *gin.Context, ser *service.Services) (err error)
	List(ctx *gin.Context) (services []service.ListService, err error)
}
