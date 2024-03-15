package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nattapat-w/chatapp/config"
	"github.com/nattapat-w/chatapp/controller"
	"github.com/nattapat-w/chatapp/core/user/entities"
	"github.com/nattapat-w/chatapp/core/user/port/repository"
	"github.com/nattapat-w/chatapp/core/user/port/service"
	"gorm.io/gorm"
)

type fiberServer struct {
	app *fiber.App
	db  *gorm.DB
	cfg *config.Config
}

func NewFiberServer(cfg *config.Config, db *gorm.DB) Server {
	return &fiberServer{
		app: fiber.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *fiberServer) Start() {
	// Initialize module(feature) here
	// ...
	// order.InitializeOrderModule(s.app, s.db)
	// userRepo := repository.NewUserRepository(db)

	s.db.AutoMigrate(&entities.User{})

	userRepo := repository.NewGormUserRepository(s.db)
	userService := service.NewUserServiceImpl(userRepo)
	userController := controller.NewUserController(userService)
	s.app.Post("/user", userController.Register)

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Listen(serverUrl)
}
