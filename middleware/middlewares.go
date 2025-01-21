package middleware

import (
	v1 "groot/controller/tasks/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JwtAuthMiddleware 验证 token
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := v1.TokenVerify(c); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
