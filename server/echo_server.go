package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/nattapat-w/chatapp/config"
	"github.com/nattapat-w/chatapp/core/auth"
	"github.com/nattapat-w/chatapp/core/user"
	"github.com/nattapat-w/chatapp/core/user/entities"

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
	if err := godotenv.Load(); err != nil {
		log.Fatal("load .env error")
	}

	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Orgin, Content-Type, Accept",
	}))

	s.db.AutoMigrate(&entities.User{})

	user.InitializeUserModule(s.app, s.db)
	auth.InitializeAuthModule(s.app, s.db)

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Listen(serverUrl)
}
