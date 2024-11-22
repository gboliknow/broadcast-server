# Broadcast Server

A CLI-based WebSocket broadcast server that enables real-time communication between clients. The server allows multiple clients to connect, send messages, and broadcast these messages to all connected clients, similar to a basic chat application.

## Features

- Start a WebSocket server to handle real-time communication.
- Connect clients to the server for sending and receiving messages.
- Broadcast messages from one client to all connected clients.
- Handles multiple clients connecting and disconnecting gracefully.

## Requirements

- Go 1.18 or later.
- Terminal for running CLI commands.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/gboliknow/broadcast-server.git
   cd broadcast-server

    Install dependencies:

    go mod tidy

Usage
Starting the Server

To start the WebSocket server:

go run . start

The server will listen on port 8080 by default.
Connecting a Client

To connect a client to the server:

go run . connect

You can open multiple terminal instances to connect multiple clients.
Sending Messages

    Type a message in the client terminal and press Enter.
    The message will be broadcasted to all connected clients, including the sender.

Example Interaction

    Client 1 sends:

Hello, everyone!

    Client 2 and others receive:

        Message received: Hello, everyone!

Project Structure

broadcast-server/
├── go.mod           # Go module file
├── main.go          # CLI entry point
├── server.go        # Server-side logic
└── client.go        # Client-side logic

Extending the Project

Here are some ideas to extend the functionality:

    Authentication: Require clients to authenticate with a username.
    Message History: Maintain and display a history of messages for new clients.
    Custom Ports: Allow specifying a port number when starting the server.
    Private Messaging: Implement direct messages between clients.
    Improved UI: Build a GUI client instead of a CLI-based one.

Testing

    Start the server:

go run . start

Connect multiple clients:

    go run . connect

    Send messages from one client and verify they are received by others.

    Observe server logs for connection and disconnection events.

License

This project is open-source and available under the MIT License.
