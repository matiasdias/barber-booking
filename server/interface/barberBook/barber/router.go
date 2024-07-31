package barber

import "github.com/gin-gonic/gin"

func RouterBarber(r *gin.RouterGroup) {
	r.POST("create", Create)
	r.GET("list", List)

}
