package main

import (
	_ "api/docs"
	"api/server/database"
	"api/server/interface/barberBook"
	"fmt"
	"net/http"
	"os"

	"github.com/fvbock/endless"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

// @title Barber Shop API
// @version 1.0
// @description This is a sample server Petstore server.

// @host localhost:5000
// @BasePath /
func main() {
	var (
		err error
		log *zap.Logger
	)
	// inicializar o logger
	defer log.Sync()
	database.Connection()

	group := errgroup.Group{}
	group.Go(func() error {
		port := os.Getenv("PORT")
		if port == "" {
			port = fmt.Sprintf("%d", database.APIConfigInfo.APIPort)
		}
		return endless.ListenAndServe(":"+port, externalRouter(log))
		//return endless.ListenAndServe(fmt.Sprintf(":%d", database.APIConfigInfo.APIPort), externalRouter(log))
	})

	if err = group.Wait(); err != nil {
		log.Error("Error while serving the application", zap.Error(err))
	} // espera que todas as rotinas adicionadas ao grupo sejam concluidas

}

func externalRouter(logg *zap.Logger) http.Handler {
	r := gin.New()
	barberBook.Router(r.Group("barber"))
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found " + " : " + c.Request.URL.String()})
	})

	api := r.Group("api")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
