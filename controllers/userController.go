package controllers

import (
	"gin-base/middlewares"
	"gin-base/models"
	"gin-base/utils/rest"
	"github.com/gin-gonic/gin"
)


/**
 * jwt颁发
*/
func JwtLogin(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user) // application-json
	// 校验用户名和密码是否正确
	if user.Username == "admin" && user.Password == "ff1234" {
		// 生成Token
		tokenString, _ := middlewares.GenToken(user.Username)
		rest.Success(c, tokenString)
		return
	}
	rest.New(c, 2002, nil, "鉴权失败")
}

/**
 * 测试获取jwt参数
*/
func NeedAuth(c *gin.Context) {
	user := c.MustGet("user").(*middlewares.JwtUser)
	rest.Success(c, user)
}
