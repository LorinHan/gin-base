package router

import (
	ct "gin-base/controllers"
	"gin-base/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(r *gin.Engine) {
	r.POST("/login", ct.JwtLogin)
	user := r.Group("user", middlewares.Auth()) // 中间件可以放在普通路由也可以放在Group上
	{
		user.GET("/needAuth", ct.NeedAuth)
		// user.POST("/login4/:username/:password", controllers.UserCtl.Login4)
		// user.GET("/find/page/:page", controllers.UserCtl.FindUsersNative)
	}

	// 使用模板引擎
	r.LoadHTMLGlob("templates/*")
	r.GET("/testTemp", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", gin.H{"title": "tempTest", "content": "内容"})
	})
}
