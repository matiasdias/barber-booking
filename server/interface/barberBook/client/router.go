package client

import "github.com/gin-gonic/gin"

func RouterClient(r *gin.RouterGroup) {
	r.POST("create", Create)
	r.GET("list", List)
}
