package collection

import (
	"app/internal/collection/workflow"
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

// Summary of FSM
// Objective: Attempt collection at various times in the day
//	but must not have overlaps for each unique customer by CustomerID
//	Rules + strategy for collection can dynamically change occasionally
//	When in SuspendedState; try out alternative collection strategy
//		try out sync first; finally only ACH ...

var client iwf.Client
var workerService iwf.WorkerService

func init() {
	var copts *iwf.ClientOptions
	// If need to customize ..
	//copts := &iwf.ClientOptions{
	//	ServerUrl:     "",
	//	WorkerUrl:     "",
	//	ObjectEncoder: nil,
	//}
	client = iwf.NewClient(workflow.GetRegistry(), copts)
	var wopts *iwf.WorkerOptions
	// If need to customize ..
	//wopts := &iwf.WorkerOptions{ObjectEncoder: nil}
	workerService = iwf.NewWorkerService(workflow.GetRegistry(), wopts)
}

func CallPayment() {
	// STart workflow if not yet ..
	// Get DataObject of currentWorkflowID
	// Pop out the next item?
	// Follow it as instructions ..
}

func BasicStartWorkflow(ctx context.Context, wf iwf.Workflow, input any) (string, error) {
	// dEBUG
	//spew.Dump(ctx)
	fmt.Println("REQID:", middleware.GetReqID(ctx))
	wfID := "mleow-2"
	// If need options?
	//wfOptions := iwf.WorkflowOptions{
	//	WorkflowIdReusePolicy: nil,
	//	WorkflowCronSchedule:  nil,
	//	WorkflowRetryPolicy: &iwfidl.RetryPolicy{
	//		InitialIntervalSeconds: nil,
	//		BackoffCoefficient:     nil,
	//		MaximumIntervalSeconds: nil,
	//		MaximumAttempts:        nil,
	//	},
	//	StartStateOptions: &iwfidl.WorkflowStateOptions{
	//		SearchAttributesLoadingPolicy: nil,
	//		DataObjectsLoadingPolicy:      nil,
	//		CommandCarryOverPolicy:        nil,
	//		StartApiTimeoutSeconds:        nil,
	//		DecideApiTimeoutSeconds:       nil,
	//		StartApiRetryPolicy:           nil,
	//		DecideApiRetryPolicy:          nil,
	//	},
	//	InitialSearchAttributes: []iwfidl.SearchAttribute{
	//		{Key: nil,
	//			StringValue:      nil,
	//			IntegerValue:     nil,
	//			DoubleValue:      nil,
	//			BoolValue:        nil,
	//			StringArrayValue: nil,
	//			ValueType:        nil,
	//		},
	//	},
	//}
	runID, err := client.StartWorkflow(ctx, wf, wfID, 3600, input, nil)
	if err != nil {
		// If already running; not fatal, just no-op
		if !iwf.IsWorkflowAlreadyStartedError(err) {
			spew.Dump(err)
			return "", err
		}
	}

	return runID, nil
}

func BasicInvokeStartHandler(ctx context.Context, req iwfidl.WorkflowStateStartRequest) (*iwfidl.WorkflowStateStartResponse, error) {
	//spew.Dump(req)
	resp, err := workerService.HandleWorkflowStateStart(ctx, req)
	if err != nil {
		spew.Dump(err)
		return nil, err
	}

	return resp, nil
}

func BasicInvokeDecideHandler(ctx context.Context, req iwfidl.WorkflowStateDecideRequest) (*iwfidl.WorkflowStateDecideResponse, error) {
	// spew.Dump(req)
	resp, err := workerService.HandleWorkflowStateDecide(ctx, req)
	if err != nil {
		spew.Dump(err)
		return nil, err
	}

	return resp, nil
}
