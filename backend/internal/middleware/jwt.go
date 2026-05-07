// Package middleware contains reusable Fiber middleware for this application.
package middleware

import (
	"strings"

	"example.com/interview-question-009/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

// JWTProtected returns a Fiber handler that validates the Bearer token in the
// Authorization header. On success it stores userID and displayName in c.Locals
// for downstream handlers to read.
func JWTProtected(authUC usecase.AuthUseCase) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing or invalid token"})
		}

		claims, err := authUC.ValidateToken(strings.TrimPrefix(authHeader, "Bearer "))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		c.Locals("userID", claims.UserID)
		c.Locals("displayName", claims.DisplayName)
		return c.Next()
	}
}
