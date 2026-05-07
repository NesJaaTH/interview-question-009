package handler

import (
	"example.com/interview-question-009/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

// AuthHandler handles HTTP requests for authentication.
type AuthHandler struct {
	useCase usecase.AuthUseCase
}

// NewAuthHandler constructs an AuthHandler with the given use case.
func NewAuthHandler(uc usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{useCase: uc}
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login handles POST /api/auth/login.
// Returns 400 for missing fields, 401 for invalid credentials, 200 with token + user on success.
func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req loginRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if req.Username == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "username and password are required"})
	}

	token, user, err := h.useCase.Login(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":           user.ID,
			"username":     user.Username,
			"display_name": user.DisplayName,
		},
	})
}
