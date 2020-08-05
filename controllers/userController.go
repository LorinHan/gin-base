package controllers

import (
	"fmt"
	"gin-base/models"
	"gin-base/mysql"
	"gin-base/utils/auth"
	"gin-base/utils/rest"
	"github.com/gin-gonic/gin"
)

type userCtl struct {}

var UserCtl = &userCtl{}

/**
	jwt颁发
 */
func (u *userCtl) JwtLogin(c *gin.Context) {
	fmt.Println(mysql.DB)
	var user models.User
	c.ShouldBindJSON(&user) // application-json
	// 校验用户名和密码是否正确
	if user.Username == "admin" && user.Password == "ff1234" {
		// 生成Token
		tokenString, _ := auth.GenToken(user.Username)
		c.JSON(rest.Success(tokenString))
		return
	}
	c.JSON(rest.New(2002, nil, "鉴权失败"))
}
/**
	测试获取jwt参数
 */
func (u *userCtl) NeedAuth(c *gin.Context) {
	user := c.MustGet("user").(*auth.JwtUser)
	c.JSON(rest.Success(&user))
}