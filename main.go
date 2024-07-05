package main

import (
	"api/server/database"
	"api/server/interface/barberBook"
	"fmt"
	"net/http"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

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
		return endless.ListenAndServe(fmt.Sprintf(":%d", database.APIConfigInfo.APIPort), externalRouter(log))
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
	return r
}
