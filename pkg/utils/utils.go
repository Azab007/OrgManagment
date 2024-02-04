package utils

import (
	"OrgManagementApp/pkg/database/mongodb/models"
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// ExtractBearerToken extracts the Bearer token from the Authorization header
func ExtractBearerToken(authorizationHeader string) (string, error) {
	// Check if the Authorization header has the correct format
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid Bearer token format")
	}

	return parts[1], nil
}

var jwtSecret = []byte("somesecretjwtcode")

// VerifyAccessToken verifies the validity of the access token
// and returns the user ID if the token is valid
func VerifyAccessToken(accessToken string) (string, error) {
	// Parse the access token
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method and return the secret key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return jwtSecret, nil
	})

	// Check for parsing errors
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid access token")
	}

	// Extract user ID from claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error extracting claims from access token")
	}

	email, ok := claims["sub"].(string)
	if !ok {
		return "", fmt.Errorf("error extracting user ID from access token")
	}

	return email, nil
}

func extractEmailFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return []byte("somesecretjwtcode"), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract the email from the claims
		if email, ok := claims["sub"].(string); ok {
			return email, nil
		}
	}

	return "", fmt.Errorf("unable to extract email from token")
}
func ExtractEmailFromTokenContext(c *gin.Context) (string, error) {
	tokenString, err := getTokenFromContext(c)
	if err != nil {
		return "", err
	}

	return extractEmailFromToken(tokenString)
}

func getTokenFromContext(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	// Convert to string
	tokenStr := tokenString

	return tokenStr, nil
}

func IsUserIDInMembers(org models.Organization, emailToCheck string) bool {

	for _, member := range org.Members {
		if member.Email == emailToCheck {
			return true
		}
	}
	return false
}
