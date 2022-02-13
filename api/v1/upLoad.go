package v1

import (
	"fmt"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpLoad(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println("我是错误！", err)
	}
	fileSize := fileHeader.Size
	fmt.Println("我是文件大小", fileSize)
	fmt.Println("我是文件", file)
	url, code := model.UpLoadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
