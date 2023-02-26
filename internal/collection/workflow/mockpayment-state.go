package workflow

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"net/http"
	"time"
)

// basic skeleton below; replace MockPayment with your own ..
//const (
//	MockPaymentID = "MockPayment"
//)

type MockPaymentState struct {
	iwf.DefaultStateIdAndOptions
}

// GetStateId is optional with above annotation
//func (b MockPaymentState) GetStateId() string {
//	return MockPaymentID
//}

// MockResponse if fields are not pub will not transfer across ..
type MockResponse struct {
	Status int
	Body   json.RawMessage
}

func (b MockPaymentState) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	fmt.Println("MOCKPAY_START")
	// Setup first time
	v := make([]MockResponse, 0)
	//var mr MockResponse
	var initialized bool
	var status string
	input.Get(&status)
	if status != "" {
		var l bool
		persistence.GetStateLocal("init", &l)
		spew.Dump(l)
		fmt.Println("Got STATUS:", status)
		initialized = true
		persistence.GetDataObject("mocker", &v)
		spew.Dump(v)
	} else {
		fmt.Println("non-empty Status; expect to see alert")
		persistence.GetStateLocal("init", &initialized)
		spew.Dump(initialized)
	}

	if !initialized {
		// Put in some dummy data
		v = append(v, MockResponse{
			Status: http.StatusOK,
			Body:   nil,
		})
		v = append(v, MockResponse{
			Status: http.StatusInternalServerError,
			Body:   nil,
		})
		//persistence.SetStateLocal("mock", v)
		persistence.SetDataObject("mocker", v)
		persistence.SetStateLocal("init", true)
		//persistence.SetStateLocal("mr", MockResponse{
		//
		//})
		//persistence.SetDataObject("mocker", MockResponse{
		//	Status: http.StatusOK,
		//	Body:   nil,
		//})
	}
	time.Sleep(time.Second)
	spew.Dump(v)
	// Block to simulate slowness
	// Decide to return an unexpected error
	return iwf.EmptyCommandRequest(), nil
}

func (b MockPaymentState) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	fmt.Println("MOCKPAY_DECIDE")
	v := make([]MockResponse, 0)
	var mocker MockResponse
	//persistence.GetStateLocal("mock", &v)
	persistence.GetDataObject("mocker", &v)
	//persistence.GetDataObject("mocker", &mocker)
	spew.Dump(mocker)
	fmt.Println("=== BELOW After GetStateLocal =========")
	spew.Dump(v)
	//persistence.GetDataObject("mock", &v)
	//spew.Dump(v)
	var initialized bool
	fmt.Println("non-empty Status; expect to see alert")
	persistence.GetStateLocal("init", &initialized)
	spew.Dump(initialized)

	if len(v) == 0 {
		fmt.Println("END!!!")
		// Once data all pop out .. then can end ..
		return iwf.GracefulCompleteWorkflow(1), nil
	}
	time.Sleep(time.Second)
	//else pop and try agan ..
	v = v[1:]
	fmt.Println("=== BELOW After Pop =========")
	spew.Dump(v)
	//persistence.SetStateLocal("mock", v)
	persistence.SetDataObject("mocker", v)

	//persistence.SetStateLocal("init", true)
	return iwf.SingleNextState(MockPaymentState{}, "init"), nil
}

// GetStateOptions is Optional ..
//func (b MockPaymentState) GetStateOptions() *iwfidl.WorkflowStateOptions {
//
//	iwfidl.NewNullableWorkflowStateOptions(nil)
//	return nil
//}
