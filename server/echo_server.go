package server

import (
	"fmt"
	"log"

	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nattapat-w/chatapp/config"
	"github.com/nattapat-w/chatapp/core/user/entities"

	"gorm.io/gorm"
)

type echoServer struct {
	app *echo.Echo
	db  *gorm.DB
	cfg *config.Config
}

func NewEchoServer(cfg *config.Config, db *gorm.DB) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *echoServer) Start() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("load .env error")
	}

	// s := echo.New()
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	s.app.Use(middleware.Logger())

	s.db.AutoMigrate(&entities.User{})

	// InitializeUserModule()

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

// func (s *echoServer) InitializeUserModule() {
// 	userRepo := repository.NewGormUserRepository(s.db)
// 	userService := service.NewUserServiceImpl(userRepo)
// 	userController := user.NewUserController(userService)

// 	userRouters := s.app.Group("/api/v1/user")
// 	userRouters.POST("/register", userController.Register)
// }
