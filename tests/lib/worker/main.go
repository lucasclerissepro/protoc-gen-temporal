package main

import (
	"fmt"

	"github.com/lucasclerissepro/protoc-gen-temporal/tests/lib/activity"
	"github.com/lucasclerissepro/protoc-gen-temporal/tests/lib/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})

	if err != nil {
    fmt.Printf("error when configuring client: %s", err)
	}
	defer c.Close()

	w := worker.New(c, "greeter.v1", worker.Options{})
	w.RegisterWorkflow(workflow.GreeterWorkflow)
	w.RegisterActivity(activity.LogString)
	err = w.Run(worker.InterruptCh())
	if err != nil {
    fmt.Printf("error when configuring worker: %s", err)
		return
	}
}
