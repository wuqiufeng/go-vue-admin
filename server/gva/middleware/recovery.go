package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
)

// RecoveryMiddleware捕获所有panic，并且返回错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//先做一下日志记录
				log.Errorf("Error:%s, Stack:%s", fmt.Sprint(err), string(debug.Stack()))
				ResponseError(c, ResCodeInnerServerError, errors.New("内部错误"))
			}
		}()
		c.Next()
	}
}
