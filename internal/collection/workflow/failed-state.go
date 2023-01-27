package workflow

import (
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

// basic skeleton below; replace Failed with your own ..
//const (
//	FailedID = "Failed"
//)

type FailedState struct {
	iwf.DefaultStateIdAndOptions
}

// GetStateId is optional with above annotation
//func (b FailedState) GetStateId() string {
//	return FailedID
//}

// Start for Failed triggers change in strategy; wfID + strategyID
func (b FailedState) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	fmt.Println("INIT_START")
	return iwf.EmptyCommandRequest(), nil
}

func (b FailedState) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	fmt.Println("INIT_DECIDE")
	return iwf.GracefulCompleteWorkflow(1), nil
}

// GetStateOptions is Optional ..
//func (b FailedState) GetStateOptions() *iwfidl.WorkflowStateOptions {
//
//	iwfidl.NewNullableWorkflowStateOptions(nil)
//	return nil
//}
