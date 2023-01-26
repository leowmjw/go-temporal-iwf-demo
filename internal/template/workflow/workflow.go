package workflow

import "github.com/indeedeng/iwf-golang-sdk/iwf"

var registry = iwf.NewRegistry()

func init() {
	err := registry.AddWorkflows(
		// Generally Workflow should be pointer; otherwise it does not stick :(
		&__REPLACE__Workflow{},
	)
	if err != nil {
		panic(err)
	}
}

func GetRegistry() iwf.Registry {
	return registry
}

type __REPLACE__Workflow struct {
	iwf.EmptyCommunicationSchema
	iwf.EmptyPersistenceSchema
	iwf.DefaultWorkflowType
}

//func setupStates() (SteadyState, DiffState) {
//	// Can it be shared? Risk for address will be multi-node is not by copy ..
//	// likely better by copy, and one per state; no sharing ..
//	db := &PostgresDB{}
//	return SteadyState{db: db}, DiffState{db: db}
//}

func (b __REPLACE__Workflow) GetStates() []iwf.StateDef {
	//ss, ds := setupStates()
	return []iwf.StateDef{
		//iwf.StartingStateDef(&stead{}),
		//iwf.NonStartingStateDef(&basicWorkflowState2{}),
		//iwf.StartingStateDef(ss),
		//iwf.NonStartingStateDef(ApprovalState{}),
		//iwf.NonStartingStateDef(ds),
	}
}

// All below no need with above annotations
//func (b __REPLACE__Workflow) GetPersistenceSchema() []iwf.PersistenceFieldDef {
//	psc := []iwf.PersistenceFieldDef{
//		iwf.DataObjectDef("TrackedTables"),
//	}
//	return psc
//}
//
//func (b __REPLACE__Workflow) GetCommunicationSchema() []iwf.CommunicationMethodDef {
//	cmd := []iwf.CommunicationMethodDef{
//		iwf.SignalChannelDef(SignalName),
//		iwf.InterstateChannelDef(SignalName),
//	}
//	// DEBUG
//	//spew.Dump(cmd)
//	return cmd
//}
//
//func (b __REPLACE__Workflow) GetWorkflowType() string {
//	return ""
//}
