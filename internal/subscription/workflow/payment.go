package workflow

import (
	"fmt"
	"go.temporal.io/sdk/workflow"
)

type PaymentWorkflowReq struct {
	ID string // Billing Anchor
}

type PaymentActivities struct {
}

// PaymentWorkflow dummy
func PaymentWorkflow(ctx workflow.Context) {
	fmt.Println("WF PaymentWorkflow ..")
	// Calculate the Billing Anchor; which is the WorkflowID
	// STarts out as Draft ..

	// If ACH; check back in 4 days for any returns ..
	pact := PaymentActivities{}
	pact.CollectSubscription("FOO")

}

func (pa PaymentActivities) CollectSubscription(idemKey string) error {
	// Use idempotencyKey to ensure no duplicate collections are made ..
	fmt.Println("Inside .. CollectSubscription")
	return nil
}
