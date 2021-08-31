package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func RequestLogMiddleware(ctx *gin.Context) {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	if ctx.Query("name") == "x" {
		fmt.Printf("The query name was: %s", ctx.Query("name"))
		// 记录到文件。
		f, _ := os.OpenFile("request.log", os.O_CREATE|os.O_WRONLY, 0)
		defer f.Close()
		f.WriteString("The query name was: " + ctx.Query("name") + "\n")

		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

	ctx.Next()
}
