package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nattapat-w/chatapp/core/user/repository"
	"github.com/nattapat-w/chatapp/core/user/service"
	"gorm.io/gorm"
)

func InitializeUserModule(app *fiber.App, db *gorm.DB) {
	userRepo := repository.NewGormUserRepository(db)
	userService := service.NewUserServiceImpl(userRepo)
	userController := NewUserController(userService)

	userRouters := app.Group("/api/v1/user")
	userRouters.Post("/register", userController.Register)
	userRouters.Get("/users", userController.GetAllUser)
}
