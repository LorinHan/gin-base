package middlewares

import (
	"crypto"
	"errors"
	"gin-base/conf"
	"gin-base/utils/rest"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type JwtUser struct {
	Username string `json:"username"`
	Role     string
	jwt.StandardClaims
}

func init() {
	// 方式1：用jwt中不包含的签名方法创建签名对象（如sha1），可以参照源码自定义：
	SigningMethodSHA1 = &jwt.SigningMethodHMAC{"SHA1", crypto.SHA1}
	jwt.RegisterSigningMethod(SigningMethodSHA1.Alg(), func() jwt.SigningMethod {
		return SigningMethodSHA1
	})
}

var TokenExpireDuration = time.Hour * conf.Jwt.Expired
var MySecret = []byte(conf.Jwt.Secret)
var SigningMethodSHA1 *jwt.SigningMethodHMAC

// GenToken，生成token
func GenToken(username string) (string, error) {
	c := JwtUser{
		username,
		"role",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
		},
	}

	token := jwt.NewWithClaims(SigningMethodSHA1, c)
	// 方式2：使用jwt库中已有的签名方法创建签名对象
	// token := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*JwtUser, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &JwtUser{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtUser); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// AuthMiddleware 基于JWT的认证中间件
func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 从请求头中获取token
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			rest.Error(c, "请求头中auth为空")
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			rest.Error(c, "请求头中auth格式有误")
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			rest.Error(c, "无效的token")
			c.Abort()
			return
		}
		// 验证角色
		// if mc.Role != role {
		//	rest.New(c, 403, nil, "没有"+role+"角色权限")
		//	c.Abort()
		//	return
		// }
		// 将当前请求的user信息保存到请求的上下文c上
		c.Set("user", mc)
		c.Next() // 后续的处理函数可以用过 c.Get("user").(*auth.JwtUser) 来获取当前请求的用户信息
	}
}
