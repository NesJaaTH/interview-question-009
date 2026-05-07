package usecase

import "example.com/interview-question-009/internal/domain"

// CommentUseCase defines the application operations available on comments.
type CommentUseCase interface {
	// ListComments returns all comments for the given post in chronological order.
	ListComments(postID uint) ([]domain.Comment, error)
	// AddComment creates a new comment on a post under the given author name.
	AddComment(postID uint, content string, author string) (*domain.Comment, error)
	// RemoveComment deletes the comment with the given primary key.
	RemoveComment(id uint) error
}

type commentUseCase struct {
	repo domain.CommentRepository
}

// NewCommentUseCase constructs a CommentUseCase backed by the given repository.
func NewCommentUseCase(repo domain.CommentRepository) CommentUseCase {
	return &commentUseCase{repo: repo}
}

func (u *commentUseCase) ListComments(postID uint) ([]domain.Comment, error) {
	return u.repo.GetByPostID(postID)
}

func (u *commentUseCase) AddComment(postID uint, content string, author string) (*domain.Comment, error) {
	comment := &domain.Comment{
		PostID:  postID,
		Author:  author,
		Content: content,
	}
	if err := u.repo.Create(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

func (u *commentUseCase) RemoveComment(id uint) error {
	return u.repo.DeleteByID(id)
}
