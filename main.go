package main

import (
	_ "gin-base/conf"
	"gin-base/middlewares"
	"gin-base/models"
	"gin-base/router"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func main() {
	// 生产环境下开启
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	middlewares.SetLogs()
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	router.Init(r)
	r.Run(":8080")
	defer models.DBClose()
}
