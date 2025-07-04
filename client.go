package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	connection *websocket.Conn
	manager    *Manager

	// egress is used to avoid concurrent writes to the connection. 
	egress chan []byte
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		egress:     make(chan []byte, 100),
	}
}

func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	for {
		messageType, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("error reading message:", err)
			}
			break
		}

		for wsclient := range c.manager.clients {
			wsclient.egress <- payload
		}
		log.Println(messageType)
		log.Println("received message:", string(payload))
	}
}

func (c *Client) writeMessages() {
    defer c.manager.removeClient(c)

    for msg := range c.egress {
        if err := c.connection.WriteMessage(websocket.TextMessage, msg); err != nil {
            log.Println("error writing message:", err)
            return
        }
        log.Println("sent message:", string(msg))
    }
	
    c.connection.WriteMessage(websocket.CloseMessage, nil)
}

