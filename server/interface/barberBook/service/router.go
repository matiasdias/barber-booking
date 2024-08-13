package service

import (
	"github.com/gin-gonic/gin"
)

func RouterService(r *gin.RouterGroup) {
	r.POST("create", Create)
	r.GET("list", List)
}
