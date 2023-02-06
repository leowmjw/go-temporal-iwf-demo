package subscription

import (
	"app/internal/subscription/workflow"
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"github.com/stripe/stripe-go/v74"
	"strings"

	"github.com/stripe/stripe-go/v74/paymentintent"
)

// ========================================
// FSM -https://youtu.be/uNwbdQyLpns?t=449
// =========================================
// Summary of FSM
// StartState: INIT (setup stuff)
// TerminalState(s):
//	- FAILED - no trial, 1st Payment Fail
//	- TRIALING - with trial
//	- ACTIVE - no trial, 1st Payment Succeed,
//		Payment Succeed Next Cycle, Payment Retry Succeed
//	- PAST_DUE - Payment failed next billing cycle, Trial Ends + 1st Payment Failed
//	- CANCELED - auto-cancel after Payment Retry Fails (after 23 hours no recover),
//		Consumer Cancel Subscription (from ACTIVE + TRIALING)

// ******************************************************
// Cadence Workflow - https://youtu.be/uNwbdQyLpns?t=693
// ******************************************************
// Summary of Cadence Workflow
// Once WF starts; it Calculates Next Billing Anchor
//	==> BillingAnchor - IdempotencyKey
// DTP will spawn off a recovery in 23 hours; failure means SubsExpire
// Events(s):
//	- CancelSubscription, SubscriptionExpired
// In TRIALING - 30 days, with reminder send 7 days before expiry
//	Draft Invoice Created + next Billing Cycle Starts w. Billing Anchor

// Start + Stop Events - https://youtu.be/uNwbdQyLpns?t=630
// WorkflowID - uniqueID per customer (can have MMs)
// Create Subscription
// Cancel Subscription

// DTP (Direct-to-Pay?)

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
	// Override the SubsID ..
	if v, ok := input.(string); ok {
		if strings.Contains(v, "sub_") {
			wfID = v
		}
	}
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
		// If already running; not fatal, just no-op
		if !iwf.IsWorkflowAlreadyStartedError(err) {
			spew.Dump(err)
			return "", err
		}
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

func ConfirmPaymentIntent(piID string) {
	stripe.Key = "sk_test_51MVCg1JJLlTnVKtUb4jx0jYfuRDjIk1wz2oArUcHGyef0d5uoPIcjGgLVOrJIDk6JHQnvsZTvt8psVuNytNcWCwu00xO72IcIC"

	piID = "pi_3MWjv8JJLlTnVKtU1fcufqKq"
	// Todo .. match client secret?

	// To create a PaymentIntent for confirmation, see our guide at: https://stripe.com/docs/payments/payment-intents/creating-payment-intents#creating-for-automatic
	//params := &stripe.PaymentIntentConfirmParams{
	//	PaymentMethod: stripe.String("pm_card_visa"),
	//	OffSession:    stripe.Bool(true),
	//}
	// Below example for ACH - pm_usBankAccount_success
	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_usBankAccount_success"),
		OffSession:    stripe.Bool(true),
		ReturnURL:     stripe.String("http://localhost:8803/sim/webhook"),
	}
	pi, err := paymentintent.Confirm(
		piID,
		params,
	)
	if err != nil {
		spew.Dump(err)
	} else {
		spew.Dump(pi)
	}

}

func CreatePaymentIntent(idemKey string) {

	// Set your secret key. Remember to switch to your live secret key in production.
	// See your keys here: https://dashboard.stripe.com/apikeys
	stripe.Key = "sk_test_51MVCg1JJLlTnVKtUb4jx0jYfuRDjIk1wz2oArUcHGyef0d5uoPIcjGgLVOrJIDk6JHQnvsZTvt8psVuNytNcWCwu00xO72IcIC"

	params := &stripe.PaymentIntentParams{
		//PaymentMethodTypes: stripe.StringSlice([]string{
		//	"card",
		//	"us_bank_account",
		//}),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
		Amount:   stripe.Int64(99),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Customer: stripe.String("cus_NHDbqmuqNVHF0R"),
		//PaymentMethod: stripe.String("{{CARD_ID}}"),
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		spew.Dump(err)
	} else {
		spew.Dump(pi)
	}
}
