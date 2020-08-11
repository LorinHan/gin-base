package main

import (
	_ "gin-base/conf"
	"gin-base/middlewares"
	"gin-base/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 生产环境下开启
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middlewares.LoggerToFile())

	router.Init(r)
	// defer mysql.DB.Close()

	r.Run(":8080")
}
