package api

import (
	"github.com/gin-gonic/gin"
	bean2 "go-gin-example/bean"
	"go-gin-example/cache"
	"log"
	"net/http"
	"time"
)

type Info struct {
	Key      string
	InfoBody InfoBody
}

type InfoBody struct {
	Id   int
	Name string
	Date bean2.Time
}

func TestCacheSet(c *gin.Context) {
	var i Info
	c.ShouldBindJSON(&i)
	err := cache.Set(i.Key, i.InfoBody, time.Hour)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, "success")
}

func TestCacheGet(c *gin.Context) {
	key := c.Query("key")

	var i InfoBody
	cache.Get(key, &i)

	c.JSON(http.StatusOK, i)

}
