package utils

import (
	"money-tracker/configs"
	"money-tracker/models"

	"github.com/golang-jwt/jwt"
)

func GetJwtToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      user.Id,
		"email":   user.Email,
		"name":    user.Name,
		"picture": user.PhotoUrl,
	})
	secretKey := []byte(configs.EnvJwtSecret())
	signedToken, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
