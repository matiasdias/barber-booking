package main

import (
	_ "api/docs"
	"api/server/auth"
	"api/server/database"
	"api/server/interface/barberBook"
	"api/server/logger"
	"api/server/middleware"
	"api/server/token"
	"net/http"

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

	log, err = logger.ConfigLogger()
	if err != nil {
		zap.L().Fatal("Erro ao inicializar o logger: ", zap.Error(err))
		return
	}
	defer log.Sync()
	zap.ReplaceGlobals(log)

	log.Info("Aplicação iniciada",
		zap.String("versao", "1.0"),
	)

	// Conecta ao banco de dados
	database.Connection()

	// Inicializar a configuração do JWT
	err = token.InitJwt()
	if err != nil {
		zap.L().Fatal("Erro ao inicializar o JWT: ", zap.Error(err))
	}

	//Inicializar a configuração do Oauth2 do Google
	err = auth.InitAuthOauth()
	if err != nil {
		zap.L().Fatal("Erro ao inicializar o Google OAuth2: ", zap.Error(err))
	}

	// Usar um grupo de goroutines para executar o servidor
	group := errgroup.Group{}
	group.Go(func() error {
		port := database.APIConfigInfo.APIPort
		return endless.ListenAndServe(":"+port, externalRouter(log))
		//return endless.ListenAndServe(fmt.Sprintf(":%d", database.APIConfigInfo.APIPort), externalRouter(log))
	})

	if err = group.Wait(); err != nil {
		log.Error("Error while serving the application", zap.Error(err))
	}

}

func externalRouter(logg *zap.Logger) http.Handler {
	r := gin.New()
	r.Use(
		middleware.RequestLogger(logg),
	)

	barberGroup := r.Group("barber")
	barberGroup.Use(middleware.JWTAuthMiddleware())
	barberBook.Router(barberGroup)

	barberBook.AuhRouter(r.Group("auth"))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "up"})
	})

	api := r.Group("api")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found " + " : " + c.Request.URL.String()})
	})

	return r
}
