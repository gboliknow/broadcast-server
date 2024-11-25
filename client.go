package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

type Message struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

var name string

var connectClientCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the WebSocket broadcast server as a client",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr := "ws://localhost:8080/ws"
		conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
		if err != nil {
			log.Fatalf("Connection error: %v", err)
		}
		defer conn.Close()

		// Assign name if not provided
		if name == "" {
			name = uuid.NewString()
			fmt.Printf("No name provided. Using unique identifier: %s\n", name)
		} else {
			fmt.Printf("Connecting as: %s\n", name)
		}

		// Send a "connect" message to the server
		err = conn.WriteJSON(Message{Type: "connect", Name: name, Message: "Connecting"})
		if err != nil {
			log.Fatalf("Error sending name: %v", err)
		}

		// Goroutine to handle incoming messages
		go func() {
			for {
				var msg Message
				err := conn.ReadJSON(&msg)
				if err != nil {
					log.Printf("Read error: %v", err)
					break
				}
				// Print messages in the format: `Name: Message`
				fmt.Printf("%s: %s\n", msg.Name, msg.Message)
			}
		}()

		// Handle user input
		fmt.Println("Connected to server. Type messages to send:")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			err := conn.WriteJSON(Message{Type: "message", Name: name, Message: text})
			if err != nil {
				log.Printf("Write error: %v", err)
				break
			}
		}
	},
}
