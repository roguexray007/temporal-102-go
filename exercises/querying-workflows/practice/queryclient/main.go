package main

import (
	"context"
	"log"

	"log/slog"
	"os"

	"go.temporal.io/sdk/client"
	tlog "go.temporal.io/sdk/log"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: "127.0.0.1:7233", // Try explicit IP instead of localhost
		ConnectionOptions: client.ConnectionOptions{
			TLS: nil, // Disable TLS
		},
		Logger: tlog.NewStructuredLogger(
			slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelDebug,
			}))),
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// TODO Part B: Add the QueryWorkflow() call and log the result.
	// Don't forget to add "context" to your imports.
	response, err := c.QueryWorkflow(context.Background(), "queries", "", "current_state")
	if err != nil {
		log.Fatalln("Error sending the Query", err)
		return
	}
	var result string
	response.Get(&result)
	log.Println("Received Query result. Result: " + result)
}
