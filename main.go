package main

import (
	"gin-base/conf"
	"gin-base/middlewares"
	"gin-base/router"
	"github.com/gin-gonic/gin"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	// 每次启动执行 swag init 命令，重新编译文档
	middlewares.ReloadSwagger()
}

// @title gin项目基本框架
// @version 2.0.0
// @host 127.0.0.1:8080
// @securityDefinitions.apikey token
// @in header
// @name Authorization
// @BasePath /
func main() {
	// 生产环境下开启
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	middlewares.InitGinLogger(r)
	if conf.Server.Swagger {
		middlewares.SwaggerMiddleware(r)
	}
	router.Init(r)

	r.Run(":" + conf.Server.Port)
}
