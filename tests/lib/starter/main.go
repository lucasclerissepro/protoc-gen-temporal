package main

import (
	"context"
	"fmt"

	"github.com/lucasclerissepro/protoc-gen-temporal/tests/lib/workflow"
	"go.temporal.io/sdk/client"
)

func main() {
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		fmt.Printf("failed to create client: %s", err)
	}
	defer temporalClient.Close()
	// ...
	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: "greeter.v1",
	}

	ids := []string{}

	for i := 0; i < 100; i++ {

		workflowRun, err := temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, workflow.GreeterWorkflow, "lucas")
		if err != nil {
			fmt.Printf("failed to start workflow: %s", err)
		}

	  ids = append(ids, workflowRun.GetRunID())
	}

  fmt.Println(ids)

}
