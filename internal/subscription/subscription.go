package subscription

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
