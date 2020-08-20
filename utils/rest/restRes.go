package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @description: restful接口，统一返回风格为 {"status": xxx, "data": xxx, "message": xxx}
 * 	除非程序报错，出现panic，否则http协议的状态码始终为200
 * 	而接口真正的状态码，在 Rest 中的 status
 * @author: Lorin
 * @time: 2020/7/19 上午11:54
 */
type Rest struct {
	Status int `json:"status"`
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
