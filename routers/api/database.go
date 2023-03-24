package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/models"
	"net/http"
	"strconv"
)

type user struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

func TestSave(c *gin.Context) {
	var u user
	c.ShouldBindJSON(&u)
	models.Create(u.Name, u.Sex)
}

func TestFind(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)

	u := models.Find(id)

	c.JSON(http.StatusOK, u)
}
