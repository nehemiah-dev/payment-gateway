// internal/middleware/recovery.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				zap.L().Error(
					"panic recovered",
					zap.Any("error", err),
					zap.String("request_id", ctx.GetString("request_id")),
				)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "internal server error",
				})
			}
		}()
		ctx.Next()
	}
}
