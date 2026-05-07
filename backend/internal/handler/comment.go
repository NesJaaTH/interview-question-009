package handler

import (
	"strconv"

	"example.com/interview-question-009/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

// CommentHandler handles HTTP requests related to comments.
type CommentHandler struct {
	useCase usecase.CommentUseCase
}

// NewCommentHandler constructs a CommentHandler with the given use case.
func NewCommentHandler(uc usecase.CommentUseCase) *CommentHandler {
	return &CommentHandler{useCase: uc}
}

// createCommentRequest is the expected JSON body for POST /:postID/comments.
type createCommentRequest struct {
	Content string `json:"content"`
}

// GetComments handles GET /:postID/comments.
// Returns 400 if the post id is not a valid integer.
func (h *CommentHandler) GetComments(c fiber.Ctx) error {
	postID, err := strconv.ParseUint(c.Params("postID"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid post id"})
	}

	comments, err := h.useCase.ListComments(uint(postID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(comments)
}

// CreateComment handles POST /:postID/comments.
// Returns 400 for a non-integer post id, malformed JSON body, or empty content.
// Returns 201 with the created comment on success.
func (h *CommentHandler) CreateComment(c fiber.Ctx) error {
	postID, err := strconv.ParseUint(c.Params("postID"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid post id"})
	}

	var req createCommentRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if req.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "content is required"})
	}

	author, _ := c.Locals("displayName").(string)
	comment, err := h.useCase.AddComment(uint(postID), req.Content, author)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(comment)
}

// DeleteComment handles DELETE /:postID/comments/:commentID.
// Returns 400 if the comment id is not a valid integer, 204 on success.
func (h *CommentHandler) DeleteComment(c fiber.Ctx) error {
	commentID, err := strconv.ParseUint(c.Params("commentID"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid comment id"})
	}

	if err := h.useCase.RemoveComment(uint(commentID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
