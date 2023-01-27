package workflow

import (
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

// basic skeleton below; replace Init with your own ..
//const (
//	InitID = "Init"
//)

type InitState struct {
	iwf.DefaultStateIdAndOptions
}

// GetStateId is optional with above annotation
//func (b InitState) GetStateId() string {
//	return InitID
//}

// Start pulls the latest dynamic strategy as config; as input, normal no input
func (b InitState) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	fmt.Println("INIT_START")
	return iwf.EmptyCommandRequest(), nil
}

func (b InitState) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	fmt.Println("INIT_DECIDE")
	return iwf.GracefulCompleteWorkflow(1), nil
}

// GetStateOptions is Optional ..
//func (b InitState) GetStateOptions() *iwfidl.WorkflowStateOptions {
//
//	iwfidl.NewNullableWorkflowStateOptions(nil)
//	return nil
//}
