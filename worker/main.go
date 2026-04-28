package main

import (
	"log"

	helloworld "github.com/temporalio/samples-go/standalone-activity/activity"
	greetings "github.com/temporalio/samples-go/standalone-activity/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/contrib/envconfig"
	"go.temporal.io/sdk/worker"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(envconfig.MustLoadDefaultClientOptions())
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "standalone-activity-helloworld", worker.Options{})

	greetingActivity := *helloworld.NewGreeting(c)
	w.RegisterActivity(greetingActivity.Activity)

	w.RegisterWorkflow(greetings.GreetingSample)
	activities := &greetings.Activities{Name: "Temporal", Greeting: "Hello"}
	w.RegisterActivity(activities)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
