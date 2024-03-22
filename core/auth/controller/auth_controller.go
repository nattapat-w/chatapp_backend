package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nattapat-w/chatapp/core/auth/model"
	"github.com/nattapat-w/chatapp/core/auth/service"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	// var loginRequest model.LoginRequest
	req := new(model.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":      fiber.ErrBadRequest.Message,
			"status_code": fiber.ErrBadRequest.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}
	// if err := c.BodyParser(&loginRequest); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	// }
	token, err := ac.authService.Login(req)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      token,
	})
}
