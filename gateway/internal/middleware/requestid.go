package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rid := ctx.GetHeader("X-request-ID")
		if rid == "" {
			rid = uuid.NewString()
		}
		ctx.Set("request_id", rid)
		ctx.Writer.Header().Set("X-request-ID", rid)
		ctx.Next()
	}
}
