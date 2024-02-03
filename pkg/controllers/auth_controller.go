package controllers

import (
	"OrgManagementApp/pkg/database"
	"OrgManagementApp/pkg/database/mongodb/models"
	"OrgManagementApp/pkg/database/mongodb/repository"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// SignUpController handles user signup
func SignUpController(user models.User) error {
	// Validate user data (add more comprehensive validation as needed)
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return fmt.Errorf("Invalid user data")
	}
	db := database.GetDB()
	userRepo := repository.NewUserRepository(db)
	// Check if the user with the same email already exists
	existingUser, err := userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return fmt.Errorf("Error checking existing user: %v", err)
	}
	if existingUser != nil {
		return fmt.Errorf("User with email %s already exists", user.Email)
	}

	// Save the user to MongoDB using the repository
	err = userRepo.SaveUser(user)
	if err != nil {
		return fmt.Errorf("Failed to sign up: %v", err)
	}

	return nil
}

func AuthenticateUser(user models.User) (*models.User, bool, error) {
	// Retrieve the user from the database based on the provided email
	db := database.GetDB()
	userRepo := repository.NewUserRepository(db)
	// Retrieve the user from the database based on the provided email
	storedUser, err := userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, false, err
	}

	// Check if the user exists and the provided password matches the stored password
	if storedUser != nil && storedUser.Password == user.Password {
		// Return the authenticated user without the password
		return &models.User{
			Name:  storedUser.Name,
			Email: storedUser.Email,
			// Omitting password for security reasons
		}, true, nil
	}

	// If the user is not found or the password doesn't match, return an authentication error
	return nil, false, fmt.Errorf("authentication failed")
}

var jwtSecret = []byte("somesecretjwtcode")

// GenerateTokens generates access and refresh tokens for the authenticated user

func generateAccessToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Access token expiration time
	return token.SignedString(jwtSecret)
}

func generateRefreshToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix() // Refresh token expiration time
	return token.SignedString(jwtSecret)
}

func GenerateTokens(user models.User) (string, string, error) {
	accessToken, err := generateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateRefreshToken(user)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

// RefreshToken refreshes the access token using a valid refresh token
func RefreshToken(refreshToken string) (string, string, error) {
	// Validate the refresh token (you should have your own validation logic)
	if isValidRefreshToken(refreshToken) {
		// If the refresh token is valid, generate a new access token and refresh token pair
		user, err := getUserFromRefreshToken(refreshToken)
		if err != nil {
			return "", "", err
		}

		accessToken, newRefreshToken, err := GenerateTokens(*user)
		if err != nil {
			return "", "", err
		}

		return accessToken, newRefreshToken, nil
	}

	// If the refresh token is invalid, return an error
	return "", "", fmt.Errorf("invalid refresh token")
}

// isValidRefreshToken checks if the refresh token is valid (e.g., not expired, not revoked)
func isValidRefreshToken(refreshToken string) bool {
	// Parse the refresh token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method and return the secret key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return jwtSecret, nil
	})

	// Check for parsing errors
	if err != nil || !token.Valid {
		return false
	}

	return true
}

func getUserFromRefreshToken(refreshToken string) (*models.User, error) {
	// Parse the refresh token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method and return the secret key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return jwtSecret, nil
	})

	// Check for parsing errors
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid refresh token")
	}

	// Extract user information from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("error extracting claims from refresh token")
	}

	// Retrieve user email from claims
	email, ok := claims["sub"].(string)
	if !ok {
		return nil, fmt.Errorf("error extracting user email from refresh token")
	}

	// In a real-world scenario, you would query the database to get the user based on the email
	// Here, we're creating a dummy user for demonstration purposes
	db := database.GetDB()
	userRepo := repository.NewUserRepository(db)
	user, err := userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
