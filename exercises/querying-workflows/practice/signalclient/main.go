package main

import (
	"context"
	"log"

	queries "temporal-102-go/exercises/querying-workflows/practice"

	"go.temporal.io/sdk/client"
	tlog "go.temporal.io/sdk/log"
	"log/slog"
	"os"
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

	signal := queries.FulfillOrderSignal{
		Fulfilled: true,
	}

	err = c.SignalWorkflow(context.Background(), "queries", "", "fulfill-order-signal", signal)
	if err != nil {
		log.Fatalln("Error sending the Signal", err)
		return
	}
}
