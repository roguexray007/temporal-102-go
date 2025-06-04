package main

import (
	"log"
	"log/slog"
	"os"
	durableexecution "temporal-102-go/exercises/durable-execution"
	"temporal-102-go/exercises/durable-execution/app"

	"go.temporal.io/sdk/client"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
)

func main() {
	// custom logger will come here
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

	w := worker.New(c, app.DurableExecutionTranslationTaskQueue, worker.Options{})
	w.RegisterWorkflow(durableexecution.SayHelloGoodbye)
	w.RegisterActivity(durableexecution.TranslateTerm)

	log.Println("Started worker", "TaskQueue", app.DurableExecutionTranslationTaskQueue)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
