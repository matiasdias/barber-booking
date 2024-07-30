package barberShop

import "github.com/gin-gonic/gin"

func RouterBarberShop(r *gin.RouterGroup) {
	r.POST("create", Create)
	r.GET("list", List)
}
