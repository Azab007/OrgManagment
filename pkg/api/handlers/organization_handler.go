package handlers

import (
	"OrgManagementApp/pkg/controllers"
	"OrgManagementApp/pkg/database/mongodb/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrganization handler to add an organization
func CreateOrganization(c *gin.Context) {
	// Parse organization information from the request (example: using JSON)
	var newOrganization models.Organization
	if err := c.BindJSON(&newOrganization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the controller method to handle user signup
	err := controllers.CreateOrganizationController(newOrganization)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization Created successfully"})
}

// GetOrganizationByIDHandler retrieves an organization by ID
func GetOrganizationByID(c *gin.Context) {
	orgID := c.Param("id")

	org, err := controllers.GetOrganizationByIDController(orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, org)
}
func GetAllOrganizations(c *gin.Context) {

	orgs, err := controllers.GetAllOrganizationController()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orgs)
}

func UpdateOrganization(c *gin.Context) {
	orgID := c.Param("id")
	var updatedOrg models.Organization
	if err := c.ShouldBindJSON(&updatedOrg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedOrg.ID = orgID
	err := controllers.UpdateOrganizationController(&updatedOrg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedOrg)
}

func DeleteOrganization(c *gin.Context) {
	orgID := c.Param("id")
	err := controllers.DeleteOrganizationController(orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization Deleted successfully"})
}
func InviteToOrganization(c *gin.Context) {
	orgID := c.Param("id")
	// Parse member information from the request (example: using JSON)
	var newMember models.Member
	if err := c.ShouldBindJSON(&newMember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := controllers.InviteToOrganizationController(orgID, &newMember)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}
