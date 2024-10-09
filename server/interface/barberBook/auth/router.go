package auth

import (
	"github.com/gin-gonic/gin"
)

func RouterGoogleAuth(r *gin.RouterGroup) {
	r.GET("/login", Login)
	r.GET("callback", CallBack)
}
