package main

import (
	"github.com/gin-gonic/gin"
	_ "go-gin-example/cache"
	"go-gin-example/conf"
	_ "go-gin-example/conf"
	_ "go-gin-example/models"
	"go-gin-example/routers"
)

func main() {
	r := gin.Default()

	routers.InitRouter(r)

	port := ":" + conf.Config.Host.Port
	r.Run(port)
}
