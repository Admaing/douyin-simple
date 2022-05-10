package routers

import (
	"douyijn-simple/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	apiV1 := r.Group("douyin")
	apiV1.GET("/", controller.Hello)
	apuV1.GET("/feed/", cont)
	return r
}
