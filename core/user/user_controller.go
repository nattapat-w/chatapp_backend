package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nattapat-w/chatapp/core/user/model"
	"github.com/nattapat-w/chatapp/core/user/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register(c *fiber.Ctx) error {
	var registerRequest model.RegisterRequest
	if err := c.BodyParser(&registerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	response := uc.userService.Register(&registerRequest)
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (uc *UserController) GetAllUser(c *fiber.Ctx) error {
	response, err := uc.userService.GetAllUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("ERROR")
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}
