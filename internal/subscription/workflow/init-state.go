package workflow

import (
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

// Init_ID basic skeleton below; replace Init_ with your own ..
const (
	Init_ID = "InitID"
)

type Init_State struct{}

func (b Init_State) GetStateId() string {
	return Init_ID
}

func (b Init_State) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	fmt.Println("INIT_START")
	return iwf.EmptyCommandRequest(), nil
}

func (b Init_State) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	fmt.Println("INIT_DECIDE")
	return iwf.GracefulCompleteWorkflow(1), nil
}

func (b Init_State) GetStateOptions() *iwfidl.WorkflowStateOptions {

	//iwfidl.NewNullableWorkflowStateOptions(nil)
	return nil
}
