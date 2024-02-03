package middlewares

import (
	"OrgManagementApp/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifies the Bearer access token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Bearer token from the Authorization header
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Bearer token"})
			c.Abort()
			return
		}

		// Check if the Bearer token is well-formed
		accessToken, err := utils.ExtractBearerToken(authorizationHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Bearer token format"})
			c.Abort()
			return
		}

		// Verify the access token
		email, err := utils.VerifyAccessToken(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			c.Abort()
			return
		}

		// Attach user ID to the context for further processing
		c.Set("email", email)

		c.Next()
	}
}
