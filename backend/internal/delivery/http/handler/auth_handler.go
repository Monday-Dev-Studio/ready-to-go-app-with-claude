package handler

import (
	"errors"
	"log/slog"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"app/internal/domain"
	"app/internal/usecase"
	"app/pkg/middleware"
	"app/pkg/response"
)

type AuthHandler struct {
	auth             *usecase.AuthUsecase
	validate         *validator.Validate
	refreshExpiryDays int
}

func NewAuthHandler(auth *usecase.AuthUsecase, refreshExpiryDays int) *AuthHandler {
	return &AuthHandler{
		auth:              auth,
		validate:          validator.New(),
		refreshExpiryDays: refreshExpiryDays,
	}
}

type registerRequest struct {
	Name     string `json:"name"     validate:"required,min=2,max=100"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type loginRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type userResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req registerRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "invalid request body")
	}
	if err := h.validate.Struct(req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	user, err := h.auth.Register(c.Context(), usecase.RegisterInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, domain.ErrEmailAlreadyTaken) {
			return response.Conflict(c, "email already taken")
		}
		slog.Error("register user", "err", err)
		return response.InternalError(c)
	}

	return response.Created(c, userResponse{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req loginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "invalid request body")
	}
	if err := h.validate.Struct(req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	tokens, err := h.auth.Login(c.Context(), usecase.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			return response.Unauthorized(c)
		}
		slog.Error("login", "err", err)
		return response.InternalError(c)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		HTTPOnly: true,
		Secure:   c.Protocol() == "https",
		SameSite: "Lax",
		MaxAge:   h.refreshExpiryDays * 24 * 60 * 60,
		Path:     "/",
	})

	return response.OK(c, fiber.Map{
		"access_token": tokens.AccessToken,
		"user": userResponse{
			ID:        tokens.User.ID.String(),
			Name:      tokens.User.Name,
			Email:     tokens.User.Email,
			CreatedAt: tokens.User.CreatedAt,
		},
	})
}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return response.Unauthorized(c)
	}

	accessToken, err := h.auth.RefreshTokens(c.Context(), refreshToken)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidToken) {
			return response.Unauthorized(c)
		}
		slog.Error("refresh token", "err", err)
		return response.InternalError(c)
	}

	return response.OK(c, fiber.Map{"access_token": accessToken})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HTTPOnly: true,
		Secure:   c.Protocol() == "https",
		SameSite: "Lax",
		MaxAge:   -1,
		Path:     "/",
	})
	return response.Message(c, "logged out")
}

func (h *AuthHandler) Me(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	user, err := h.auth.GetUser(c.Context(), userID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return response.NotFound(c, "user")
		}
		slog.Error("get user", "err", err)
		return response.InternalError(c)
	}

	return response.OK(c, userResponse{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	})
}
