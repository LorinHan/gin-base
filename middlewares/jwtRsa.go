package middlewares

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"time"
)

var (
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
)

func init() {
	// openssl genrsa -out private.key 2048
	// openssl rsa -in private.key -pubout > public.key
	publicKeyByte, err := ioutil.ReadFile("/Users/lorin/Documents/go/gin-base/public.key")
	if err != nil {
		log.Println(err.Error())
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)
	privateKeyByte, err := ioutil.ReadFile("/Users/lorin/Documents/go/gin-base/private.key")
	if err != nil {
		log.Println(err.Error())
	}
	privateKey, _ = jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
}

func GenRsaToken(username string) (string, error) {
	c := JwtUser{
		username,
		"role",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	return token.SignedString(privateKey)
}

func ParseRsaToken(tokenString string) (*JwtUser, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtUser{}, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("验证Token的加密类型错误")
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtUser); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
