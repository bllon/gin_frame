package bootstrap

import (
	"gin_frame/middleware"
	"gin_frame/routers"
	"github.com/gin-gonic/gin"
)

func App(HttpServer *gin.Engine) {
	HttpServer = gin.New()

	//使用中间件
	HttpServer.Use(gin.Recovery())
	HttpServer.Use(middleware.RequestLogMiddleware)

	//设置基本路由
	routers.BasicRouter(HttpServer)

	HttpServer.Run()
}
