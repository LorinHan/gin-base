package router

import (
	ct "gin-base/controllers"
	"gin-base/middlewares"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.POST("/login", ct.JwtLogin)
	user := r.Group("user", middlewares.Auth()) // 中间件可以放在普通路由也可以放在Group上
	{
		user.GET("/needAuth", ct.NeedAuth)
		//user.POST("/login4/:username/:password", controllers.UserCtl.Login4)
		//user.GET("/find/page/:page", controllers.UserCtl.FindUsersNative)
	}
}
