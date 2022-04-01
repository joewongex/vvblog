package middleware

import (
	"fmt"
	"net/http"
	"vvblog/errors/vcode"
	"vvblog/errors/verror"
	"vvblog/vlog"

	"github.com/gin-gonic/gin"
)

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			r := recover()
			if r == nil {
				return
			}

			statusCode := http.StatusInternalServerError
			if err, ok := r.(error); ok {
				code := verror.Code(err)
				if code == vcode.CodeInvalidParameter || code == vcode.CodeInvalidRequest || code == vcode.CodeBusinessValidationFailed {
					statusCode = http.StatusBadRequest
				} else if code == vcode.CodeNotFound {
					statusCode = http.StatusNotFound
				} else if code == vcode.CodeNotAuthorized {
					statusCode = http.StatusUnauthorized
				} else if code == vcode.CodeNil {
					// 类似runtime error，并没有在之前被捕获
					err = verror.WrapCode(vcode.CodeInternalError, err)
				}

				if statusCode == http.StatusInternalServerError {
					// %+s只打印堆栈信息
					if verr, ok := err.(*verror.Error); ok {
						vlog.Errorf("%s\n%+s", verr.Error(), verr.Current())
					} else {
						vlog.Error(err)
					}
				}

				c.JSON(statusCode, gin.H{
					"error": fmt.Sprintf("%-v", err),
				})
			}
		}()

		c.Next()
	}
}
