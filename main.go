package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/conf"
	"go-gin-example/routers"
)

func init() {
	conf.InitConfig()
}

func main() {
	r := gin.Default()

	routers.InitRouter(r)

	port := ":" + conf.Config.Host.Port
	r.Run(port)
}