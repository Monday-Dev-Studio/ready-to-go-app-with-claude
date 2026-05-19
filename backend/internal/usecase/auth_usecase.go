package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"app/internal/domain"
	"app/pkg/middleware"
)

type AuthUsecase struct {
	users              domain.UserRepository
	accessSecret       string
	refreshSecret      string
	accessExpiryMinutes int
	refreshExpiryDays  int
}

type RegisterInput struct {
	Name     string `validate:"required,min=2,max=100"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type LoginInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type AuthTokens struct {
	AccessToken  string
	RefreshToken string
	User         *domain.User
}

func NewAuthUsecase(
	users domain.UserRepository,
	accessSecret, refreshSecret string,
	accessExpiryMinutes, refreshExpiryDays int,
) *AuthUsecase {
	return &AuthUsecase{
		users:               users,
		accessSecret:        accessSecret,
		refreshSecret:       refreshSecret,
		accessExpiryMinutes: accessExpiryMinutes,
		refreshExpiryDays:   refreshExpiryDays,
	}
}

func (u *AuthUsecase) Register(ctx context.Context, in RegisterInput) (*domain.User, error) {
	_, err := u.users.FindByEmail(ctx, in.Email)
	if err == nil {
		return nil, domain.ErrEmailAlreadyTaken
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return nil, fmt.Errorf("check email: %w", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), 12)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	now := time.Now()
	user := &domain.User{
		ID:           uuid.New(),
		Email:        in.Email,
		PasswordHash: string(hash),
		Name:         in.Name,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err = u.users.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *AuthUsecase) Login(ctx context.Context, in LoginInput) (*AuthTokens, error) {
	user, err := u.users.FindByEmail(ctx, in.Email)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, domain.ErrInvalidCredentials
		}
		return nil, fmt.Errorf("find user: %w", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.Password)); err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	accessToken, err := u.generateToken(user.ID, u.accessSecret, time.Duration(u.accessExpiryMinutes)*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("generate access token: %w", err)
	}

	refreshToken, err := u.generateToken(user.ID, u.refreshSecret, time.Duration(u.refreshExpiryDays)*24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("generate refresh token: %w", err)
	}

	return &AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}

func (u *AuthUsecase) RefreshTokens(ctx context.Context, refreshToken string) (string, error) {
	claims := &middleware.Claims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrInvalidToken
		}
		return []byte(u.refreshSecret), nil
	})
	if err != nil || !token.Valid {
		return "", domain.ErrInvalidToken
	}

	_, err = u.users.FindByID(ctx, claims.UserID)
	if err != nil {
		return "", domain.ErrInvalidToken
	}

	newAccess, err := u.generateToken(claims.UserID, u.accessSecret, time.Duration(u.accessExpiryMinutes)*time.Minute)
	if err != nil {
		return "", fmt.Errorf("generate access token: %w", err)
	}

	return newAccess, nil
}

func (u *AuthUsecase) GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return u.users.FindByID(ctx, id)
}

func (u *AuthUsecase) generateToken(userID uuid.UUID, secret string, expiry time.Duration) (string, error) {
	claims := middleware.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}
