package hoursBarber

import "github.com/gin-gonic/gin"

func RouterHoursBarber(r *gin.RouterGroup) {
	r.POST("create", Create)
	r.GET("list", List)
}

func RouterHoursExecption(r *gin.RouterGroup) {
	r.POST("create", CreateException)
}
