package main

import (
	_ "gin-base/conf"
	"gin-base/middlewares"
	"gin-base/router"
	"github.com/gin-gonic/gin"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
}

func main() {
	// 生产环境下开启
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	middlewares.InitGinLogger(r)
	router.Init(r)

	r.Run(":8080")
}
