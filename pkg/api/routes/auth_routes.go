package routes

import (
	"OrgManagementApp/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes configures the user-related routes
func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		// Route for user signup
		userGroup.POST("/signup", handlers.SignUpHandler)
		userGroup.POST("/signin", handlers.SignInHandler)

	}
}
