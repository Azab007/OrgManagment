package main

import (
	"OrgManagementApp/pkg/api/routes"
	"OrgManagementApp/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	database.Connect()

	// Set up all API routes
	routes.SetupUserRoutes(r)

	// Start the server
	r.Run(":8000")
}
