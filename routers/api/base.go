package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func TestGet(c *gin.Context) {
	log.Println("test")
	c.String(http.StatusOK, "test")
}

func TestGetQuery(c *gin.Context) {
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})

}

func TestGetParam(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

type bean struct {
	Name string `form:"name" uri:"name" json:"name"`
}

func TestGetQuery1(c *gin.Context) {
	var name bean
	err := c.ShouldBindQuery(&name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
	}

	c.JSON(http.StatusOK, name)

}

func TestGetParam1(c *gin.Context) {
	var name bean
	err := c.ShouldBindUri(&name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
	}
	c.JSON(http.StatusOK, name)
}

func TestPost(c *gin.Context) {
	var name bean
	c.ShouldBindJSON(&name)

	c.JSON(http.StatusOK, name)
}
