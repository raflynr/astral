package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raflynr/astral/helper"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type authHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (ah *authHandler) Login(c *fiber.Ctx) error {
	dto := new(Login)

	if err := c.BodyParser(dto); err != nil {
		return err
	}

	token, err := ah.authService.Login(c.Context(), *dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.NewError(err.Error(), nil))
	}
	
	return c.Status(fiber.StatusCreated).JSON(helper.NewSuccess(token))
}

func (ah *authHandler) Register(c *fiber.Ctx) error {
	dto := new(Register)

	if err := c.BodyParser(dto); err != nil {
		return err
	}

	if err := ah.authService.Register(c.Context(), *dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.JSON(helper.NewSuccess(nil))
}
