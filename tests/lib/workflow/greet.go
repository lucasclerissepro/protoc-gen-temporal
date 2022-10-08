package workflow

import (
	"fmt"
	"time"

	"github.com/lucasclerissepro/protoc-gen-temporal/tests/lib/activity"
	"go.temporal.io/sdk/workflow"
)

// GreeterWorkflow simply greets user with a given name
func GreeterWorkflow(ctx workflow.Context, name string) error {
  ao := workflow.ActivityOptions{
		StartToCloseTimeout: 60 * time.Second,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)
  future := workflow.ExecuteActivity(ctx, activity.LogString, fmt.Sprintf("Hello %s", name))
  future.Get(ctx, nil)

	return nil
}
