package workflow

import "github.com/indeedeng/iwf-golang-sdk/iwf"

type SubscriptionWorkflow struct{}

//func setupStates() (SteadyState, DiffState) {
//	// Can it be shared? Risk for address will be multi-node is not by copy ..
//	// likely better by copy, and one per state; no sharing ..
//	db := &PostgresDB{}
//	return SteadyState{db: db}, DiffState{db: db}
//}

func (b SubscriptionWorkflow) GetStates() []iwf.StateDef {
	//ss, ds := setupStates()
	return []iwf.StateDef{
		//iwf.NewStartingState(&stead{}),
		//iwf.NewNonStartingState(&basicWorkflowState2{}),
		//iwf.NewStartingState(ss),
		//iwf.NewNonStartingState(ApprovalState{}),
		//iwf.NewNonStartingState(ds),
		iwf.NewStartingState(Init_State{}),
	}
}

func (b SubscriptionWorkflow) GetPersistenceSchema() []iwf.PersistenceFieldDef {
	psc := []iwf.PersistenceFieldDef{
		//iwf.NewDataObjectDef("TrackedTables"),
	}
	return psc
}

func (b SubscriptionWorkflow) GetCommunicationSchema() []iwf.CommunicationMethodDef {
	cmd := []iwf.CommunicationMethodDef{
		//iwf.NewSignalChannelDef(SignalName),
		//iwf.NewInterstateChannelDef(SignalName),
	}
	// DEBUG
	//spew.Dump(cmd)
	return cmd
}

func (b SubscriptionWorkflow) GetWorkflowType() string {
	return ""
}
