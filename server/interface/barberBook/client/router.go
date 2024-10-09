package client

import "github.com/gin-gonic/gin"

func RouterClient(r *gin.RouterGroup) {
	r.GET("list", List)
}
