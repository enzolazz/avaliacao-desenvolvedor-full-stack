package handlers

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/back-end/internal/pubsub"
	"url-shortener/back-end/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func WebSocketHandler(ps *pubsub.RedisPubSub) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := utils.GetUserIDFromContext(c)
		if err != nil {
			return
		}

		conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("WebSocket upgrade failed:", err)
			return
		}
		defer conn.Close()

		channel := fmt.Sprintf("user:%s:updates", userID.Hex())
		unsubscribe, err := ps.Subscribe(channel, func(msg string) {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				log.Println("Failed to write WebSocket message:", err)
			}
		})
		if err != nil {
			log.Println("Failed to subscribe to Pub/Sub channel:", err)
			return
		}
		defer unsubscribe()

		log.Printf("WebSocket client connected for user: %s", userID.Hex())

		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				log.Printf("Client for user %s disconnected: %v", userID.Hex(), err)
				break
			}
		}
	}
}
