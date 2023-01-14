package workflow

import "go.temporal.io/sdk/workflow"

type PaymentWorkflowReq struct {
	ID string // Billing Anchor
}

type PaymentActivities struct {
}

func PaymentWorkflow(ctx workflow.Context) {
	// Calculate the Billing Anchor; which is the WorkflowID
	// STarts out as Draft ..

	// If ACH; check back in 4 days for any returns ..

}

func (pa PaymentActivities) CollectSubscription(idemKey string) error {
	// Use idempotencyKey to ensure no duplicate collections are made ..
	return nil
}
