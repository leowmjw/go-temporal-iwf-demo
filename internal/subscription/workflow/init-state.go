package workflow

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"github.com/jackc/pgx/v5"
	"time"
)

// Init_ID basic skeleton below; replace Init_ with your own ..
const (
	Init_ID = "InitID"
)

type PostgresDB struct {
	ConnConfig *pgx.ConnConfig
	Conn       *pgx.Conn
}

type Init_State struct {
	ID      string
	history []string
	//db      *PostgresDB
}

func (b Init_State) GetStateId() string {
	return Init_ID
}

// Start will modify history; so must use pointer ..
func (b *Init_State) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	fmt.Println("INIT_START")
	// Persist when the Draft Invoice is issued ..
	// Once WF starts; it Calculates Next Billing Anchor
	// Keep track of previous history; for invoice?
	b.ID = "Dood"
	billingAnchor := fmt.Sprintf("2023%02d%02d", time.Now().Month(), time.Now().Day())
	b.history = append(b.history, billingAnchor)
	// DEBUG
	//if b.db.Conn == nil {
	//	fmt.Println("Setup ConnString in Steady ..")
	//	// connString can be passed along probably; from setup?
	//	connString := "postgres://s2admin:password@127.0.0.1:5432/myterraform"
	//	connConfig, err := pgx.ParseConfig(connString)
	//	if err != nil {
	//		// fmt.Println("ERR:", err)
	//		panic(err)
	//	}
	//	b.db.ConnConfig = connConfig
	//}
	//
	spew.Dump(b)
	time.Sleep(time.Second * 1) // Simulate some actions running ..

	iwf.AnyCommandCompletedRequest(
		iwf.NewTimerCommand("id", time.Now()),
	)
	// Get Signal from user on trial or non-trial
	// After 2 hours; abandon cart
	return iwf.EmptyCommandRequest(), nil
}

func (b Init_State) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	fmt.Println("INIT_DECIDE")
	spew.Dump(b)
	// Decision tree: Trial or no trial
	// Trial --> TRAILING
	// Non-Trial
	// Once WF starts; it Calculates Next Billing Anchor
	//	==> BillingAnchor - IdempotencyKey
	// Get Payment --> ACTIVE
	// If cannot collect payment within the hour; abandon it --> FAILED

	// Store history .. see if it appear or not ..
	return iwf.GracefulCompleteWorkflow(1), nil
}

func (b Init_State) GetStateOptions() *iwfidl.WorkflowStateOptions {

	//iwfidl.NewNullableWorkflowStateOptions(nil)
	return nil
}
