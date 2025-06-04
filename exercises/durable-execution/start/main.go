package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	durableexecution "temporal-102-go/exercises/durable-execution"
	"temporal-102-go/exercises/durable-execution/app"

	"go.temporal.io/sdk/client"
)

func main() {

	c, err := client.Dial(client.Options{
		HostPort: "127.0.0.1:7233", // Try explicit IP instead of localhost
		ConnectionOptions: client.ConnectionOptions{
			TLS: nil, // Disable TLS
		},
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "translation-workflow",
		TaskQueue: app.DurableExecutionTranslationTaskQueue,
	}

	if len(os.Args) <= 2 {
		log.Fatalln("Must specify name and language code as command-line arguments")
	}

	input := durableexecution.TranslationWorkflowInput{
		Name:         os.Args[1],
		LanguageCode: os.Args[2],
	}

	workflow, err := c.ExecuteWorkflow(context.Background(), workflowOptions, durableexecution.SayHelloGoodbye, input)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", workflow.GetID(), "RunID", workflow.GetRunID())

	// Synchronously wait for workflow completion
	var result durableexecution.TranslationWorkflowOutput
	err = workflow.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalln("Unable to format result in JSON format", err)
	}
	log.Printf("Workflow result: %s\n", string(data))
}
