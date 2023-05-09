package middleware

import (
	"errors"
	"money-tracker/configs"
	"money-tracker/models"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthenticateUser(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	// fmt.Println(c.())
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(
			http.StatusUnauthorized).JSON(
			models.Response{
				Status:  http.StatusUnauthorized,
				Message: "error",
				Data:    "Unauthorized",
			},
		)
	}

	// Extract the token from the header
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse the JWT token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Verify that the signing method is HMAC with SHA256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}

		// Return the secret key used to sign the token
		return []byte(configs.EnvJwtSecret()), nil
	})

	// Check for parsing errors
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Failed to parse JWT token",
		})
	}

	// Check if the token is valid
	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid JWT token",
		})
	}

	// Check if the token has expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Expired JWT token",
		})
	}

	// check if the token belongs to the user with the given id in the request
	reqUserId := c.Params("userId")

	// Set the user ID from the token as a context value
	claims := token.Claims.(jwt.MapClaims)

	if claims["id"] != reqUserId {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized wrong user id",
		})
	}

	c.Locals("id", claims["id"])

	// Call the next middleware in the chain
	return c.Next()
}
