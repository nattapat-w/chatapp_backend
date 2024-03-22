package chat

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nattapat-w/chatapp/core/chat/service"
	"github.com/nattapat-w/chatapp/core/user/dto"
)

type ChatController struct {
	chatService service.ChatService
}

func NewChatController(chatService service.ChatService) *ChatController {
	return &ChatController{chatService: chatService}
}

func (cc *ChatController) CreateChat(c *fiber.Ctx) error {

	userData := c.Locals("userData").(*dto.UserDataDTO)
	// chatName := c.Params("chatname")

	type CreateChatRequest struct {
		UserID   uint   `json:"userID"`
		ChatName string `json:"chatname"`
	}
	var requestBody CreateChatRequest

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err := cc.chatService.CreateChat(userData.ID, requestBody.UserID, requestBody.ChatName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":         "OK",
		"status_code":    fiber.StatusOK,
		"message":        "",
		"Chat Name":      requestBody.ChatName,
		"Create By User": userData.Username,
	})
	// response := cc.chatService.CreateChat(userData)
	// return c.Status(fiber.StatusCreated).JSON(response)
}
