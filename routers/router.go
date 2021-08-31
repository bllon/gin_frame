package routers

import (
	"gin_frame/controllers/index"
	"gin_frame/controllers/user"
	"github.com/gin-gonic/gin"
)

func BasicRouter(router *gin.Engine) {
	router.GET("/", index.Index)
	router.GET("/hello", user.Index)
}
