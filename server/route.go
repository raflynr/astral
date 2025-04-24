package server

import (
	"database/sql"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/raflynr/astral/config"
	"github.com/raflynr/astral/helper"
	"github.com/raflynr/astral/internal/auth"
)

func MiddlewareJWT() fiber.Handler {
	secretKey := config.NewConfig().JWT.Secret

	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(helper.NewError("Missing or invalid Authorization header", nil))
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := helper.VerifyJWT(tokenStr, secretKey)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(helper.NewError("Invalid or expired token", nil))
		}

		c.Locals("username", claims.Username)

		return c.Next()
	}
}

func NewRoute(db *sql.DB, validate *validator.Validate, app *fiber.App) {
	repository := auth.NewAuthRepository(db)
	service := auth.NewAuthService(repository, validate)
	handler := auth.NewAuthHandler(service)

	api := app.Group("/api")

	auth := api.Group("/auth")
	{
		auth.Post("/register", handler.Register)
		auth.Post("/login", handler.Login)
	}

	profile := api.Group("/profile")
	{
		profile.Use(MiddlewareJWT())
	}

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Health")
	})
}
