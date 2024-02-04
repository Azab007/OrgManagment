package controllers

import (
	"OrgManagementApp/pkg/database"
	"OrgManagementApp/pkg/database/mongodb/models"
	"OrgManagementApp/pkg/database/mongodb/repository"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func CreateOrganizationController(org models.Organization) (string, error) {
	// Validate org data
	if org.Name == "" || org.Description == "" {
		return "", fmt.Errorf("Invalid organization data")
	}
	org.ID = uuid.NewString()
	db := database.GetDB()
	orgRepo := repository.NewOrganizationRepository(db)

	err := orgRepo.CreateOrganization(&org)
	if err != nil {
		return "", fmt.Errorf("Error creating organization: %v", err)
	}

	return org.ID, nil

}

func GetOrganizationByIDController(orgID string) (*models.Organization, error) {
	db := database.GetDB()
	orgRepo := repository.NewOrganizationRepository(db)

	org, err := orgRepo.GetOrganizationByID(orgID)
	if err != nil {
		return nil, fmt.Errorf("Error getting organization: %v", err)
	}

	return org, nil
}
func GetAllOrganizationController() ([]models.Organization, error) {
	db := database.GetDB()
	orgRepo := repository.NewOrganizationRepository(db)

	orgs, err := orgRepo.GetAllOrganizations()
	if err != nil {
		return nil, fmt.Errorf("Error getting organizations: %v", err)
	}

	return orgs, nil
}
func UpdateOrganizationController(org *models.Organization) error {
	db := database.GetDB()
	orgRepo := repository.NewOrganizationRepository(db)

	err := orgRepo.UpdateOrganization(org)
	if err != nil {
		return fmt.Errorf("Error updating organization: %v", err)
	}

	return nil
}
func DeleteOrganizationController(orgID string) error {
	db := database.GetDB()
	orgRepo := repository.NewOrganizationRepository(db)

	err := orgRepo.DeleteOrganization(orgID)
	if err != nil {
		return fmt.Errorf("Error deleting organization: %v", err)
	}

	return nil
}
func InviteToOrganizationController(orgID string, newMember *models.Member) error {
	db := database.GetDB()
	userRepo := repository.NewUserRepository(db)
	orgRepo := repository.NewOrganizationRepository(db)

	// Retrieve the existing organization
	existingOrg, err := orgRepo.GetOrganizationByID(orgID)
	if err != nil {
		return err
	}
	existingUser, err := userRepo.GetUserByEmail(newMember.Email)
	if err != nil {
		return err
	}
	newMember.Name = existingUser.Name

	// Check if the member already exists in the organization
	for _, member := range existingOrg.Members {
		if member.Email == newMember.Email {
			return errors.New("Member with the same email already exists in the organization")
		}
	}
	// Add the new member to the organization
	existingOrg.Members = append(existingOrg.Members, *newMember)

	// Update the organization in the repository
	err = orgRepo.UpdateOrganization(existingOrg)
	if err != nil {
		return err
	}

	return nil
}
