package utils

// Example using jwt-go package
import (
	"OrgManagementApp/pkg/database/mongodb/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("somesecretjwtcode")

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
