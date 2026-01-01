package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/blen/task_manager_api/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
			c.Abort()
			return
		}
		secretKey := os.Getenv("JWT_SECRET_KEY")

		claims, err := utils.ValidateJWT(tokenString, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		userID, ok := (*claims)["user_id"]

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id missing in token"})
			c.Abort()
			return
		}

		role, ok := (*claims)["role"]

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "role missing in token"})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Set("role", role)

		c.Next()
	}

}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")

		if !exists || role != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Admin access required"})
			return
		}

		c.Next()
	}
}
