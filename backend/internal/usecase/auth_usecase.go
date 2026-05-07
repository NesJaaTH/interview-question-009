package usecase

import (
	"errors"
	"time"

	"example.com/interview-question-009/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Claims holds the JWT payload stored inside every access token.
type Claims struct {
	UserID      uint   `json:"user_id"`
	DisplayName string `json:"display_name"`
	jwt.RegisteredClaims
}

// AuthUseCase defines authentication operations.
type AuthUseCase interface {
	// Login validates credentials and returns a signed JWT plus the user record on success.
	Login(username, password string) (token string, user *domain.User, err error)
	// ValidateToken parses and verifies a JWT string, returning its claims.
	ValidateToken(token string) (*Claims, error)
}

type authUseCase struct {
	repo      domain.UserRepository
	jwtSecret string
}

// NewAuthUseCase constructs an AuthUseCase with the given repository and JWT signing secret.
func NewAuthUseCase(repo domain.UserRepository, jwtSecret string) AuthUseCase {
	return &authUseCase{repo: repo, jwtSecret: jwtSecret}
}

func (u *authUseCase) Login(username, password string) (string, *domain.User, error) {
	user, err := u.repo.GetByUsername(username)
	if err != nil {
		// Return a generic message to avoid leaking whether the username exists.
		return "", nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	claims := &Claims{
		UserID:      user.ID,
		DisplayName: user.DisplayName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return "", nil, err
	}

	return tokenStr, user, nil
}

func (u *authUseCase) ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(u.jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	return claims, nil
}
