package controllers

import (
	"fmt"
	"ginTest/models"
	"ginTest/mysql"
	"ginTest/utils"
	"ginTest/utils/Rest"
	"github.com/gin-gonic/gin"
	"net/http"
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
		tokenString, _ := utils.GenToken(user.Username)
		c.JSON(http.StatusOK, Rest.Success(tokenString))
		return
	}
	c.JSON(http.StatusOK, Rest.Res(2002, nil, "鉴权失败"))
}
/**
	测试获取jwt参数
 */
func (u *userCtl) NeedAuth(c *gin.Context) {
	user := c.MustGet("user").(*utils.MyClaims)
	c.JSON(http.StatusOK, Rest.Success(&user))
}