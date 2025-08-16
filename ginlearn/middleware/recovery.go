package middleware

import (
	"ginlearn/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			// 使用 recover 捕获 panic 的函数
			// 如果发生 panic，记录错误信息并返回 500 错误
			if err := recover(); err != nil {
				logger.Log.WithFields(logrus.Fields{
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
					"err":    err,
					"ip":     c.ClientIP(),
				}).Error("服务器内部异常")

				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "服务器内部异常",
				})
				// 中止请求处理
				c.Abort()
			}
		}()
		c.Next()
	}
}
