package main

import (
	"log"

	queries "temporal-102-go/exercises/querying-workflows/practice"

	"log/slog"
	"os"

	"go.temporal.io/sdk/client"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
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

	w := worker.New(c, "queries", worker.Options{})

	w.RegisterWorkflow(queries.Workflow)
	w.RegisterActivity(queries.Activity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
