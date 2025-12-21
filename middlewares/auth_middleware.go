package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		login, err := c.Cookie("login")

		if err != nil || login != "true" {
			c.Redirect(http.StatusFound, "/login")

			c.Abort()
			return
		}

		c.Next()
	}

}
