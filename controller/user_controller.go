package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nattapat-w/chatapp/core/user/model/request"
	"github.com/nattapat-w/chatapp/core/user/port/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register(c *fiber.Ctx) error {
	var registerRequest request.RegisterRequest
	if err := c.BodyParser(&registerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	response := uc.userService.Register(&registerRequest)
	return c.Status(fiber.StatusCreated).JSON(response)
}
