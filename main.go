package main

import (
	_ "gin-base/conf"
	"gin-base/router"
	"gin-base/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// 生产环境下开启
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(utils.LoggerToFile())

	router.Init(r)
	// defer mysql.DB.Close()

	r.Run(":8080")
}
