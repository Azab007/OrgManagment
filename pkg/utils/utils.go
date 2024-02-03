package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
