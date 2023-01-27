package workflow

import (
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

// basic skeleton below; replace Pending with your own ..
//const (
//	PendingID = "Pending"
//)

type PendingState struct {
	iwf.DefaultStateIdAndOptions
}

// GetStateId is optional with above annotation
//func (b PendingState) GetStateId() string {
//	return PendingID
//}

// Start for Pending means the async collection runs; most will be fast, some like ACH might be 4 days ..
func (b PendingState) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	fmt.Println("INIT_START")
	return iwf.EmptyCommandRequest(), nil
}

func (b PendingState) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	fmt.Println("INIT_DECIDE")
	// Successful will mean completion; mark the DB
	// Result as successful; set reuse to FALSE, retention 1 day
	// so if OK, complete and it cannot accidentally re-run ..
	// Failure after fix n-time will trigger a re-strategy ..
	return iwf.GracefulCompleteWorkflow(1), nil
}

// GetStateOptions is Optional ..
//func (b PendingState) GetStateOptions() *iwfidl.WorkflowStateOptions {
//
//	iwfidl.NewNullableWorkflowStateOptions(nil)
//	return nil
//}
