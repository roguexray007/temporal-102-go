package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	// TODO Part C: Add signals "interacting/exercises/sending-signals-client/practice"
	// to your module imports.
	// TODO Part D: Add "context" to your module imports.
	signals "temporal-102-go/exercises/sending-signals-client/practice"

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

	// TODO Part C: Use the `FullfillOrderSignal` struct type from the `signals` module
	// (i.e., the `workflow.go` file in the parent directory).
	// Create an instance of `FulfillOrderSignal` that contains `Fulfilled: true`.
	signal := signals.FulfillOrderSignal{Fulfilled: true}
	err = c.SignalWorkflow(context.Background(), "signals", "", "fulfill-order-signal", signal)
	// TODO Part D: Call `SignalWorkflow()` to send a Signal to your running Workflow.
	// It needs, as arguments, `context.Background()`, your workflow ID, your run ID
	// (which can be an empty string), the name of the signal, and the signal instance.
	// It should assign its result to `err` so that it can be checked in the next line.
	if err != nil {
		log.Fatalln("Error sending the Signal", err)
		return
	}
}
