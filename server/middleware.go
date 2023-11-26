package server

import (
	"omni/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		if bearerToken != "" {
			t := strings.Split(bearerToken, " ")
			if len(t) == 2 {
				token := t[1]
				_, err := models.ParseToken(token)
				if err != nil {
					c.JSON(401, gin.H{"error": "unauthorized"})
					c.Abort()
					return
				}
			} else {
				c.JSON(401, gin.H{"error": "unauthorized"})
				c.Abort()
				return
			}
		} else {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
