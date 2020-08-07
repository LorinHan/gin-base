package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Rest struct {
	Status int `json:"status"`
	// 加上json的话，返回时的反序列化会转为指定的字符串，这里就变成小写的d开头了
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func New(c *gin.Context, status int, data interface{}, message string) {
	c.JSON(http.StatusOK, &Rest{status, data, message})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Rest{http.StatusOK, data, "success"})
}

func Error(c *gin.Context, message string) {
	c.JSON(http.StatusOK, &Rest{http.StatusInternalServerError, nil, message})
}
