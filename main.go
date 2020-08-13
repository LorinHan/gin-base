package main

import (
	_ "gin-base/conf"
	"gin-base/middlewares"
	"gin-base/models"
	"gin-base/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 生产环境下开启
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	zapLogger, zapRecovery := middlewares.Log()

	r.Use(zapLogger)
	r.Use(zapRecovery)

	router.Init(r)
	r.Run(":8080")
	defer models.DBClose()
}
