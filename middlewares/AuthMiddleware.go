package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		// 去掉 "Bearer " 前缀
		tokenString = tokenString[len("Bearer "):]

		claims, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		role, _ := (*claims)["role"].(string)
		if len(allowedRoles) == 0 || contains(allowedRoles, role) {
			c.Next()
			return
		}

	}
}

func contains(allowedRoles []string, role string) bool {
	for _, r := range allowedRoles {
		if r == role {
			return true
		}
	}
	return false
}
