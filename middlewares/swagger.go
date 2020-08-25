package middlewares

import (
	"fmt"
	"gin-base/conf"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"os/exec"
)

/**
 * @description: 重新执行swag init命令以加载修改后的swag文档
 * @author: Lorin
 * @time: 2020/8/24 下午3:42
@Tags 首页
@Id 1
@Summary 获得首页数据
@Description | 项目 | 价格 | 数量 |
@Description | :-------- | --------:| :--: |
@Description | iPhone | 6000 元 | 5 |
@Description | iPad | 3800 元 | 12 |
@Description | iMac | 10000 元 | 234 |
@Produce  json
@Security ApiKeyAuth
@Param id path int true "用户id"
@Param token header string true "用户token"
@Param name query string false "用户名"
@Param img formData file false "文件"
@Param jsCode body string true "jsCode" default({"jsCode": "xxx"})
@Param 加密数据 body entity.User true "用户实体类"
@Success 200 {object} rest.Rest {"status": 200, "data": null, "message": "success"}
@Failure 200 {object} rest.Rest {"status": 500, "data": null, "message": "error"}
@Router /api/v1/{id} [get]
*/
func ReloadSwagger() {
	cmd := exec.Command("swag", "init")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("swagger初始化成功")
}

/**
 * @description: 开启跨域、映射swagger路径
 * @author: Lorin
 * @time: 2020/8/25 下午3:22
 */
func SwaggerMiddleware(r *gin.Engine) {
	r.Use(cors())
	r.StaticFile("swagger.json", "./docs/swagger.json")
	url := ginSwagger.URL("http://127.0.0.1:" + conf.Server.Port + "/swagger.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

// 处理跨域请求,支持options访问
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
