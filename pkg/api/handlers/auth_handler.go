package handlers

import (
	"OrgManagementApp/pkg/controllers"
	"OrgManagementApp/pkg/database/mongodb/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUpHandler handles user signup requests
func SignUpHandler(c *gin.Context) {
	// Parse user information from the request (example: using JSON)
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the controller method to handle user signup
	err := controllers.SignUpController(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully"})
}

// SignInHandler handles user signup requests
func SignInHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	authenticatedUser, isAuthenticated, err := controllers.AuthenticateUser(user)
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Authentication failed"})
		return
	}

	if isAuthenticated {
		// Authentication successful, proceed with further actions
		accessToken, refreshToken, err := controllers.GenerateTokens(*authenticatedUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Successfully singed in", "access_token": accessToken, "refresh_token": refreshToken})
	} else {
		// Authentication failed
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

// RefreshTokenRequest is a struct to represent the request body for the refresh token endpoint
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// RefreshTokenHandler handles refreshing access tokens
func RefreshTokenHandler(c *gin.Context) {
	var requestBody RefreshTokenRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	refreshToken := requestBody.RefreshToken
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing refresh token"})
		return
	}

	// Validate the refresh token and get a new pair of access and refresh tokens
	accessToken, newRefreshToken, err := controllers.RefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Token refreshed successfully",
		"access_token":  accessToken,
		"refresh_token": newRefreshToken,
	})
}
