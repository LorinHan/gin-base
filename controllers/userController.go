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

// @Summary 验证token
// @Tags 用户
// @Accept json
// @Description # 本接口需要验证token
// @Description ### - token请求头为：Authorization
// @Description ### - token请求体为：Bearer + 空格 + token
// @Produce  json
// @Security token
// @Success 200 {object} rest.Rest "{"status": 200, "data": null, "message": "success"}"
// @Router /user/needAuth [get]
func NeedAuth(c *gin.Context) {
	user := c.MustGet("user").(*middlewares.JwtUser)
	rest.Success(c, user)
}
