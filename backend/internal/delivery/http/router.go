package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"app/internal/delivery/http/handler"
	"app/internal/repository"
	"app/internal/usecase"
	"app/pkg/config"
	"app/pkg/database"
	"app/pkg/middleware"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewApp(cfg *config.Config, db *pgxpool.Pool) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{"success": false, "error": err.Error()})
		},
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORS.AllowedOrigin,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	_ = database.RunMigrations(db, "/app/migrations")

	userRepo := repository.NewUserRepository(db)
	authUC := usecase.NewAuthUsecase(
		userRepo,
		cfg.JWT.AccessSecret,
		cfg.JWT.RefreshSecret,
		cfg.JWT.AccessExpiryMinutes,
		cfg.JWT.RefreshExpiryDays,
	)
	authHandler := handler.NewAuthHandler(authUC, cfg.JWT.RefreshExpiryDays)

	app.Get("/health", handler.Health)

	authRoutes := app.Group("/api/auth")
	authRoutes.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 1 * time.Minute,
	}))
	authRoutes.Post("/register", authHandler.Register)
	authRoutes.Post("/login", authHandler.Login)
	authRoutes.Post("/refresh", authHandler.Refresh)
	authRoutes.Post("/logout", authHandler.Logout)

	protected := app.Group("/api", middleware.RequireAuth(cfg.JWT.AccessSecret))
	protected.Get("/users/me", authHandler.Me)

	return app
}
