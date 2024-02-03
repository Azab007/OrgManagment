package routes

import (
	"OrgManagementApp/pkg/api/handlers"
	middlewares "OrgManagementApp/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupOrganizationRoutes(router *gin.Engine) {
	organizationGroup := router.Group("/organization").Use(middlewares.AuthMiddleware())
	{

		// Create a new organization
		organizationGroup.POST("/", handlers.CreateOrganization)

		// Get organization by ID
		organizationGroup.GET("/:id", handlers.GetOrganizationByID)

		// Get all organizations
		organizationGroup.GET("/", handlers.GetAllOrganizations)

		// Update organization by ID
		organizationGroup.PUT("/:id", handlers.UpdateOrganization)

		// Delete organization by ID
		organizationGroup.DELETE("/:id", handlers.DeleteOrganization)

		// Invite user to organization
		organizationGroup.POST("/:id/invite", handlers.InviteToOrganization)

	}
}
