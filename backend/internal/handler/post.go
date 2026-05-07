// Package handler contains the Fiber HTTP handlers (interface adapter layer).
// Each handler struct receives a use-case interface via its constructor so it
// remains decoupled from concrete implementations.
package handler

import (
	"strconv"

	"example.com/interview-question-009/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

// PostHandler handles HTTP requests related to posts.
type PostHandler struct {
	useCase usecase.PostUseCase
}

// NewPostHandler constructs a PostHandler with the given use case.
func NewPostHandler(uc usecase.PostUseCase) *PostHandler {
	return &PostHandler{useCase: uc}
}

// GetPost handles GET /:postID.
// Returns 400 if the id is not a valid integer, 404 if the post does not exist.
func (h *PostHandler) GetPost(c fiber.Ctx) error {
	postID, err := strconv.ParseUint(c.Params("postID"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid post id"})
	}

	post, err := h.useCase.GetPost(uint(postID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "post not found"})
	}
	return c.JSON(post)
}
