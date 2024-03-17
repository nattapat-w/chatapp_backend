package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nattapat-w/chatapp/controller"
	"github.com/nattapat-w/chatapp/core/user/port/repository"
	"github.com/nattapat-w/chatapp/core/user/port/service"
	"gorm.io/gorm"
)

func InitializeUserModule(app *fiber.App, db *gorm.DB) {
	userRepo := repository.NewGormUserRepository(db)
	userService := service.NewUserServiceImpl(userRepo)
	userController := controller.NewUserController(userService)

	userRouters := app.Group("/api/v1/user")
	userRouters.Post("/register", userController.Register)
}
