package template

import (
	"app/internal/template/workflow"
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

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

func BasicStartWorkflow(ctx context.Context, wf iwf.Workflow, input any) (string, error) {
	// dEBUG
	//spew.Dump(ctx)
	fmt.Println("REQID:", middleware.GetReqID(ctx))
	wfID := "mleow-0"
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
		// TODO: If already runing; show handle it .. what is underlying error type?
		spew.Dump(err)
		return "", err
	}

	return runID, nil
}

func BasicInvokeStartHandler(ctx context.Context, req iwfidl.WorkflowStateStartRequest) (*iwfidl.WorkflowStateStartResponse, error) {
	spew.Dump(req)
	resp, err := workerService.HandleWorkflowStateStart(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func BasicInvokeDecideHandler(ctx context.Context, req iwfidl.WorkflowStateDecideRequest) (*iwfidl.WorkflowStateDecideResponse, error) {
	// spew.Dump(req)
	resp, err := workerService.HandleWorkflowStateDecide(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
