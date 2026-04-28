package helloworld

import (
	"context"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
)

type Greeting struct {
	TemporalClient client.Client
}

func NewGreeting(temporalClient client.Client) *Greeting {
	return &Greeting{
		TemporalClient: temporalClient,
	}
}

func (g *Greeting) Activity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)

	tempClient := g.TemporalClient
	// trigger a workflow
	workflowOptions := client.StartWorkflowOptions{
		ID:        "standalone_activity_helloworld_workflowID",
		TaskQueue: "standalone-activity-helloworld",
	}

	we, err := tempClient.ExecuteWorkflow(ctx, workflowOptions, "GreetingSample")
	if err != nil {
		logger.Error("Unable to execute workflow", "error", err)
		return "", err
	}
	logger.Info("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	return "Hello " + name + "!", nil
}
