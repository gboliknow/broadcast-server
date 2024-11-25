package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

var startServerCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the WebSocket broadcast server",
	Run: func(cmd *cobra.Command, args []string) {
		server := NewBroadcastServer()
		http.HandleFunc("/ws", server.handleWebSocket)
		log.Println("Server starting on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	},
}

type BroadcastServer struct {
	clients  map[*websocket.Conn]bool
	mu       sync.Mutex // Only needed for managing multiple clients
	upgrader websocket.Upgrader
}

func NewBroadcastServer() *BroadcastServer {
	return &BroadcastServer{
		clients: make(map[*websocket.Conn]bool),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
	}
}

func (s *BroadcastServer) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP to WebSocket
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade failed: %v", err)
		return
	}

	// Add new client
	s.mu.Lock()
	s.clients[conn] = true
	s.mu.Unlock()

	// Clean up on disconnect
	defer func() {
		s.mu.Lock()
		delete(s.clients, conn)
		s.mu.Unlock()
		conn.Close()
	}()

	// Handle incoming messages
	for {
		var msg Message
		if err := conn.ReadJSON(&msg); err != nil {
			log.Printf("Read failed: %v", err)
			break
		}

		// Broadcast the message to all clients
		s.mu.Lock()
		for client := range s.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Write failed: %v", err)
				client.Close()
				delete(s.clients, client)
			}
		}
		s.mu.Unlock()

		// Print the received message in the format: `Name: Message`
		log.Printf("%s: %s", msg.Name, msg.Message)
	}
}
