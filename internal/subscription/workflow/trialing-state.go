package workflow

import (
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

//const (
//	TrialingID = "Trialing"
//)

type TrialingState struct {
	iwf.DefaultStateIdAndOptions
}

//func (b TrialingState) GetStateId() string {
//	return TrialingID
//}

func (b TrialingState) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	fmt.Println("INIT_START")

	// While Trailing; if:
	// 	Half - Send call-to-action
	//	Last 10% - Send final reminder
	return iwf.EmptyCommandRequest(), nil
}

func (b TrialingState) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	fmt.Println("INIT_DECIDE")
	//return iwf.GracefulCompleteWorkflow(1), nil
	// Process Timer
	//	switch to PASTDUE ..

	// Process Signal
	signal := commandResults.GetSignalCommandResultByChannel("subscription-action")
	if signal.Status == iwfidl.RECEIVED {
		signal.SignalValue.Get(nil)
		// wants to pay; moved to active, call Temporal ..
		//	moved to active .. collect
		// wants to cancel; switch to CANCELED
	}
	// If finish all reminder; go to the ActiveState
	//	After Trial time; and collected First Payment Success
	//	After Trial time; and First Payment FAILED

	// Update reminder; go again, the next round ..
	return iwf.MultiNextStates(&Init_State{}, TrialingState{}), nil
}

//func (b TrialingState) GetStateOptions() *iwfidl.WorkflowStateOptions {
//
//	iwfidl.NewNullableWorkflowStateOptions(nil)
//	return nil
//}
