package middleware

import (
	"strings"
	"vvblog/errors/vcode"
	"vvblog/errors/verror"
	"vvblog/service"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			panic(verror.NewCode(vcode.CodeMissingParameter, "缺少Authorization请求头"))
		}

		splittedHeader := strings.Split(header, " ")
		if len(splittedHeader) != 2 {
			panic(verror.NewCode(vcode.CodeInvalidParameter, "Authorization请求头格式错误"))
		}

		userClamin, err := service.Auth.ParseToken(splittedHeader[1])
		if err != nil {
			panic(err)
		}

		c.Set("username", userClamin.Username)
		c.Next()
	}
}
