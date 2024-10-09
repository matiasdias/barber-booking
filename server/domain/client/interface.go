package client

import (
	"api/server/infrastructure/persistence/client"

	"github.com/gin-gonic/gin"
)

type IClient interface {
	Create(ctx *gin.Context, cliente *client.Client) (err error)
	List(ctx *gin.Context) (clients []client.Clients, err error)
	FindByEmail(ctx *gin.Context, email *string) (existe bool, err error)
}
