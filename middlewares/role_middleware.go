package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, err := c.Cookie("role")

		if err != nil || role != allowRole {
			c.String(http.StatusForbidden, "Akses ditolak")
			c.Abort()
			return
		}

		c.Next()
	}
}
