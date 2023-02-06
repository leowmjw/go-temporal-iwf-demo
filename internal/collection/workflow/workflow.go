package workflow

import "github.com/indeedeng/iwf-golang-sdk/iwf"

var registry = iwf.NewRegistry()

func init() {
	err := registry.AddWorkflows(
		// Generally Workflow should be pointer; otherwise it does not stick :(
		&CollectionWorkflow{},
	)
	if err != nil {
		panic(err)
	}
}

func GetRegistry() iwf.Registry {
	return registry
}

type CollectionWorkflow struct {
	iwf.EmptyCommunicationSchema
	//iwf.EmptyPersistenceSchema
	iwf.DefaultWorkflowType
}

//func setupStates() (SteadyState, DiffState) {
//	// Can it be shared? Risk for address will be multi-node is not by copy ..
//	// likely better by copy, and one per state; no sharing ..
//	db := &PostgresDB{}
//	return SteadyState{db: db}, DiffState{db: db}
//}

func (b CollectionWorkflow) GetStates() []iwf.StateDef {
	//ss, ds := setupStates()
	return []iwf.StateDef{
		//iwf.StartingStateDef(&stead{}),
		//iwf.NonStartingStateDef(&basicWorkflowState2{}),
		//iwf.StartingStateDef(ss),
		//iwf.NonStartingStateDef(ApprovalState{}),
		//iwf.NonStartingStateDef(ds),
		iwf.StartingStateDef(&MockPaymentState{}),
	}
}

// All below no need with above annotations
// Not needed for statelocal ..

func (b CollectionWorkflow) GetPersistenceSchema() []iwf.PersistenceFieldDef {
	psc := []iwf.PersistenceFieldDef{
		iwf.DataObjectDef("init"),
		iwf.DataObjectDef("mock"),
		iwf.DataObjectDef("mocker"),
	}
	return psc
}

//
//func (b CollectionWorkflow) GetCommunicationSchema() []iwf.CommunicationMethodDef {
//	cmd := []iwf.CommunicationMethodDef{
//		iwf.SignalChannelDef(SignalName),
//		iwf.InterstateChannelDef(SignalName),
//	}
//	// DEBUG
//	//spew.Dump(cmd)
//	return cmd
//}
//
//func (b CollectionWorkflow) GetWorkflowType() string {
//	return ""
//}
