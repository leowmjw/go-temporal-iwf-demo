package main

import (
	"app/internal/subscription/workflow"
	"context"
	"fmt"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"time"
)

const (
	taskqueue = "tq.subscription"
)

// startTemporalWorker run own wroker for payment purpose ..
func startTemporalWorker() (closeFunc func()) {
	// COnfig for client to be extracted here ..
	c, err := client.Dial(client.Options{})
	if err != nil {
		panic(err)
	}
	// Make worjer has a slow shutdown ..
	w := worker.New(c, taskqueue, worker.Options{
		WorkerStopTimeout:       time.Second * 10,
		LocalActivityWorkerOnly: true,
	})
	w.RegisterWorkflow(workflow.SubscriptionWorkflow{})
	pact := workflow.PaymentActivities{}
	w.RegisterActivity(&pact)

	// For test only
	go func() {
		wfr, xerr := c.ExecuteWorkflow(context.Background(),
			client.StartWorkflowOptions{},
			workflow.SubscriptionWorkflow{}, nil)
		if xerr != nil {
			panic(xerr)
		}
		gerr := wfr.Get(context.Background(), nil)
		if gerr != nil {
			panic(gerr)
		}
	}()

	// Start is async ..
	serr := w.Start()
	if serr != nil {
		panic(serr)
	}

	return func() {
		fmt.Println("Closing Temporal Worker ...")
		w.Stop()
	}
}
