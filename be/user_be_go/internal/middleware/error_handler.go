package middleware

import (
	"log"
	"user_be_go/internal/constants"
	"user_be_go/internal/utils"
	"user_be_go/pkg/response"

	"github.com/gin-gonic/gin"
)

// ErrorHandler 全局错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 检查是否有错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			log.Printf("HTTP_ERROR: method=%s, url=%s, error=%v",
				c.Request.Method,
				c.Request.URL.String(),
				err,
			)

			// 处理应用错误
			if appErr, ok := err.(*utils.AppError); ok {
				response.Error(c, appErr.Status, appErr.Message)
				return
			}

			// 处理验证错误
			if c.Writer.Status() == constants.StatusBadRequest {
				response.Error(c, constants.StatusBadRequest, constants.MsgInvalidParams)
				return
			}

			// 默认错误
			response.Error(c, constants.StatusInternalServerError, constants.MsgInternalError)
		}
	}
}

