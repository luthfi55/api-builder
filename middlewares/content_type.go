package middlewares

import (
	"github.com/gin-gonic/gin"
)

func SetJSONContentTypeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		ctx.Next()
	}
}
