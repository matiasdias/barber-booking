package auth

import (
	"github.com/gin-gonic/gin"
)

func RouterGoogleAuth(r *gin.RouterGroup) {
	r.GET("/login", Login)
	r.GET("callback", CallBack)
}

func RouterJwtAuth(r *gin.RouterGroup) {
	r.POST("refresh_token", RefreshToken)
}
