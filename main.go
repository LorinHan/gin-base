package main

import (
	_ "ginTest/conf"
	"ginTest/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// 生产环境下开启
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(utils.LoggerToFile())

	InitRouter(r)
	// defer mysql.DB.Close()

	r.Run(":8080")
}
