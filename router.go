package main

import (
	"ginTest/controllers"
	"ginTest/utils/auth"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.POST("/login", controllers.UserCtl.JwtLogin)
	user := r.Group("user", auth.AuthMiddleware("role")) // 中间件可以放在普通路由也可以放在Group上
	{
		user.GET("/needAuth", controllers.UserCtl.NeedAuth)
		//user.POST("/login4/:username/:password", controllers.UserCtl.Login4)
		//user.GET("/find/page/:page", controllers.UserCtl.FindUsersNative)
	}
}
