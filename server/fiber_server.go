package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"github.com/joho/godotenv"
	"github.com/nattapat-w/chatapp/config"
	"github.com/nattapat-w/chatapp/core/auth"
	"github.com/nattapat-w/chatapp/core/chat"
	chatEntities "github.com/nattapat-w/chatapp/core/chat/entities"
	chatRepo "github.com/nattapat-w/chatapp/core/chat/repository"
	chatServ "github.com/nattapat-w/chatapp/core/chat/service"
	"github.com/nattapat-w/chatapp/core/message"
	messageEntites "github.com/nattapat-w/chatapp/core/message/entities"
	messageRepo "github.com/nattapat-w/chatapp/core/message/repostiory"
	messageServ "github.com/nattapat-w/chatapp/core/message/service"
	"github.com/nattapat-w/chatapp/core/middleware"
	"github.com/nattapat-w/chatapp/core/user"
	"github.com/nattapat-w/chatapp/core/user/dto"
	"github.com/nattapat-w/chatapp/core/user/entities"
	"github.com/nattapat-w/chatapp/core/user/repository"
	"github.com/nattapat-w/chatapp/core/user/service"
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

// var clients = make(map[*websocket.Conn]bool)
// var broadcast = make(chan []byte)

func (s *fiberServer) Start() {
	// Initialize module(feature) here

	if err := godotenv.Load(); err != nil {
		log.Fatal("load .env error")
	}

	s.app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "*",
		AllowHeaders:     "Orgin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	s.app.Use(middleware.LoggingMiddleware)

	auth.InitializeAuthModule(s.app, s.db)
	// user.InitializeUserModule(s.app, s.db)
	userRepo := repository.NewGormUserRepository(s.db)
	userService := service.NewUserServiceImpl(userRepo)
	userController := user.NewUserController(userService)

	userRouters := s.app.Group("/api/v1/user")
	userRouters.Post("/register", userController.Register)

	chatRepository := chatRepo.NewGormChatRepository(s.db)
	chatService := chatServ.NewChatSerivceImpl(chatRepository)
	chatController := chat.NewChatController(chatService)
	chatRouters := s.app.Group("/api/v1/chat")

	messageRepository := messageRepo.NewGormMessageRepository(s.db)
	messageService := messageServ.NewMessageSerivceImpl(messageRepository)
	messageController := message.NewMessageController(messageService)
	messageRotuers := s.app.Group("/api/v1/message")

	//protected
	s.app.Use(middleware.CheckToken)
	message.StartBroadcastingMessages()
	messageRotuers.Get("/ws/:id", websocket.New(messageController.CreateMessage))
	userRouters.Get("/users", userController.GetAllUser)
	chatRouters.Post("/", chatController.CreateChat)
	s.app.Get("/protected", func(c *fiber.Ctx) error {
		// Retrieve user data from the request context
		userData := c.Locals("userData").(*dto.UserDataDTO)

		return c.JSON(userData)
	})
	s.db.AutoMigrate(&entities.User{}, &chatEntities.Chat{}, &messageEntites.Message{}, &chatEntities.ChatOwners{})

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Listen(serverUrl)
}
