package template

import (
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"github.com/iworkflowio/iwf-golang-sdk/iwf"
)

// basic skeleton below; replace __REPLACE__ with your own ..
const (
	__REPLACE__ID = "__REPLACE__"
)

type __REPLACE__State struct{}

func (b __REPLACE__State) GetStateId() string {
	return __REPLACE__ID
}

func (b __REPLACE__State) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	return nil, nil
}

func (b __REPLACE__State) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	//var i int
	return nil, nil
}

func (b __REPLACE__State) GetStateOptions() *iwfidl.WorkflowStateOptions {

	iwfidl.NewNullableWorkflowStateOptions(nil)
	return nil
}
