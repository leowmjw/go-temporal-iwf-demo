package workflow

import (
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

// basic skeleton below; replace __REPLACE__ with your own ..
//const (
//	__REPLACE__ID = "__REPLACE__"
//)

type __REPLACE__State struct {
	iwf.DefaultStateIdAndOptions
}

// GetStateId is optional with above annotation
//func (b __REPLACE__State) GetStateId() string {
//	return __REPLACE__ID
//}

func (b __REPLACE__State) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	fmt.Println("INIT_START")
	return iwf.EmptyCommandRequest(), nil
}

func (b __REPLACE__State) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	fmt.Println("INIT_DECIDE")
	return iwf.GracefulCompleteWorkflow(1), nil
}

// GetStateOptions is Optional ..
//func (b __REPLACE__State) GetStateOptions() *iwfidl.WorkflowStateOptions {
//
//	iwfidl.NewNullableWorkflowStateOptions(nil)
//	return nil
//}
