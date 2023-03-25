package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "go-gin-example/cache"
	"go-gin-example/conf"
	_ "go-gin-example/models"
	"go-gin-example/routers"
	"time"
)

func main() {
	r := gin.New()

	// 自定义日志格式
	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC3339),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	routers.InitRouter(r)

	port := ":" + conf.Config.Host.Port
	r.Run(port)
}
