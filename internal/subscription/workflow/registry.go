package workflow

import "github.com/indeedeng/iwf-golang-sdk/iwf"

var registry = iwf.NewRegistry()

func init() {
	err := registry.AddWorkflows(
		// Generally Workflow should be pointer; otherwise it does not stick :(
		&SubscriptionWorkflow{},
	)
	if err != nil {
		panic(err)
	}
}

func GetRegistry() iwf.Registry {
	return registry
}
