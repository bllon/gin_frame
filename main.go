package main

import (
	"gin_frame/bootstrap"
	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func main() {
	//接收命令行参数, 判断执行命令行

	bootstrap.App(HttpServer)
}
