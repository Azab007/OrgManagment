package handlers

import (
	"OrgManagementApp/pkg/controllers"
	"OrgManagementApp/pkg/database/mongodb/models"
	"OrgManagementApp/pkg/utils"
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
	orgId, err := controllers.CreateOrganizationController(newOrganization)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"organization_id": orgId})
}

// GetOrganizationByIDHandler retrieves an organization by ID
func GetOrganizationByID(c *gin.Context) {
	orgID := c.Param("id")
	org, err := controllers.GetOrganizationByIDController(orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	email, err := utils.ExtractEmailFromTokenContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if utils.IsUserIDInMembers(*org, email) {
		c.JSON(http.StatusOK, org)
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

// GetAllOrganizations

func GetAllOrganizations(c *gin.Context) {

	orgs, err := controllers.GetAllOrganizationController()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := make([]models.Organization, 0)
	email, err := utils.ExtractEmailFromTokenContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, org := range orgs {
		if utils.IsUserIDInMembers(org, email) {
			res = append(res, org)
		}

	}
	c.JSON(http.StatusOK, res)
}

// Update Organization
func UpdateOrganization(c *gin.Context) {
	orgID := c.Param("id")
	var updatedOrg models.Organization
	if err := c.ShouldBindJSON(&updatedOrg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedOrg.ID = orgID
	updatedOrg.Members = nil

	err := controllers.UpdateOrganizationController(&updatedOrg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	responseOrg := models.Organization{
		ID:          updatedOrg.ID,
		Name:        updatedOrg.Name,
		Description: updatedOrg.Description,
	}
	c.JSON(http.StatusOK, responseOrg)
}

// Delete Organization
func DeleteOrganization(c *gin.Context) {
	orgID := c.Param("id")
	err := controllers.DeleteOrganizationController(orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization Deleted successfully"})
}

// Invite a member to organization using email

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
	c.JSON(http.StatusOK, gin.H{"message": "Successful invitation"})

}
