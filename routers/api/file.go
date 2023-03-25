package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UploadFile(c *gin.Context) {

	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dst := "./" + file.Filename

	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("%s upload success", file.Filename))

}

func DownloadFile(c *gin.Context) {
	fileName := c.Query("fileName")
	dst := "./" + fileName

	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(dst)
}
