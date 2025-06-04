package main

import (
	"log"
	pizza "temporal-102-go/samples/custom-search-attributes"

	"go.temporal.io/sdk/client"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
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

	w := worker.New(c, pizza.TaskQueueName, worker.Options{})

	w.RegisterWorkflow(pizza.PizzaWorkflow)
	w.RegisterActivity(pizza.GetDistance)
	w.RegisterActivity(pizza.SendBill)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}

}
