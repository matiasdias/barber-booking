package reservation

import (
	"github.com/gin-gonic/gin"
)

func RouterReservation(r *gin.RouterGroup) {
	r.POST("create", Create)
	r.GET("list", List)
	r.PUT("update/:id", Update)
}
