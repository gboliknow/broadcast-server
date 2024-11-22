package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "broadcast-server",
		Short: "Broadcast Server for WebSocket-based real-time messaging",
	}

	rootCmd.AddCommand(startServerCmd, connectClientCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
