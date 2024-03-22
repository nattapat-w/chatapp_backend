package message

import (
	"log"
	"strconv"

	"github.com/gofiber/websocket/v2"
	"github.com/nattapat-w/chatapp/core/message/service"
	"github.com/nattapat-w/chatapp/core/user/dto"
)

type MessageController struct {
	messageService service.MessageService
}

func NewMessageController(messageService service.MessageService) *MessageController {
	return &MessageController{messageService: messageService}
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

func (mc *MessageController) CreateMessage(c *websocket.Conn) {
	// Add client to clients map
	userData := c.Locals("userData").(*dto.UserDataDTO)

	chatID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return
	}

	clients[c] = true
	defer func() {
		delete(clients, c)
		c.Close()
	}()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return
		}

		if err := mc.messageService.CreateMessage(string(msg), userData.ID, uint(chatID)); err != nil {
			log.Println("Error creating message:", err)
			return
		}
		// Broadcast the received message to all clients
		broadcast <- msg
	}
}
func StartBroadcastingMessages() {
	go func() {
		for {
			msg := <-broadcast
			for client := range clients {
				if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
					log.Println("Write error:", err)
					return
				}
			}
		}
	}()
}
