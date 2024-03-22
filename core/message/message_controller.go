package message

import (
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
	"github.com/nattapat-w/chatapp/core/message/service"
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
		fmt.Println(string(msg))
		if err := mc.messageService.CreateMessage(string(msg), 1, 1); err != nil {
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
