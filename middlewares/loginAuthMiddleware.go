package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goPro4/utils"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
			}
			jwtKey := []byte(viper.GetString("login_auth.secret_key"))
			return jwtKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("username", claims["username"])
		} else {
			utils.Info("token失效,重新登录")
			utils.Response(c, utils.RecordLogTokenExpire, utils.UserTokenExpire, utils.UserRespMsg[utils.UserTokenExpire], make(map[string]string))
			//	校验不通过阻止请求继续执行
			c.Abort()
		}
		c.Next()
	}
}
