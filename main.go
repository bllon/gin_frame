package main

import (
	"fmt"
	"gin_frame/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

var HttpServer *gin.Engine

func main() {
	//接收命令行参数, 判断执行命令行

	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "dev" //默认读取开发环境配置
	}
	godotenv.Load(".env." + env)
	fmt.Println("env is", os.Getenv("APP_ENV"))

	bootstrap.App(HttpServer)
}
