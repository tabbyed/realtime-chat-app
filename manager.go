package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

// servws - method which belongs to manager
// server provides the w and r
func (m *Manager) servWS(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection")
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	conn.Close()
}
