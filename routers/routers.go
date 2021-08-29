package routers

import (
	"gin_frame/controllers/index"
	"gin_frame/controllers/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter() (router *gin.Engine) {
	router = gin.Default()
	router.GET("/", index.Index)
	router.GET("/hello", user.Index)
	return
}
