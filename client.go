package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

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

		go func() {
			for {
				var msg string
				err := conn.ReadJSON(&msg)
				if err != nil {
					log.Printf("Read error: %v", err)
					break
				}
				fmt.Printf("Message received: %s\n", msg)
			}
		}()

		fmt.Println("Connected to server. Type messages to send:")

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			msg := scanner.Text()
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Printf("Write error: %v", err)
				break
			}
		}
	},
}
