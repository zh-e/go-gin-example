package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/routers/api"
	"net/http"
)

func InitRouter(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	apiGroup := r.Group("/api")
	apiGroup.GET("/test/get", api.TestGet)
	apiGroup.GET("/test/get/query", api.TestGetQuery)
	apiGroup.GET("/test/get/query1", api.TestGetQuery1)
	apiGroup.GET("/test/get/param/:name", api.TestGetParam)
	apiGroup.GET("/test/get/param1/:name", api.TestGetParam1)
	apiGroup.POST("/test/post", api.TestPost)

	apiGroup.POST("/database/create", api.TestSave)
	apiGroup.GET("/database/find", api.TestFind)

	apiGroup.POST("/cache/set", api.TestCacheSet)
	apiGroup.GET("/cache/get", api.TestCacheGet)

}
