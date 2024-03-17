package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nattapat-w/chatapp/core/auth/controller"
	authRepo "github.com/nattapat-w/chatapp/core/auth/port/repository"
	"github.com/nattapat-w/chatapp/core/auth/port/service"
	userRepo "github.com/nattapat-w/chatapp/core/user/port/repository"
	"gorm.io/gorm"
)

func InitializeAuthModule(app *fiber.App, db *gorm.DB) {
	userRepository := userRepo.NewGormUserRepository(db)
	authRepository := authRepo.NewGormAuthRepository(db)
	authService := service.NewAuthServiceImpl(authRepository, userRepository)
	authController := controller.NewAuthController(authService)

	authRouters := app.Group("/api/v1/user")
	authRouters.Post("/login", authController.Login)
}
