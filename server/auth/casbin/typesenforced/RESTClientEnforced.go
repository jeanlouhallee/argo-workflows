package typesenforced

import (
	"k8s.io/client-go/rest"
)

type RESTClientEnforced struct {
	delegate rest.Interface
}

func (w WorkflowEnforcedInterface) RESTClient() rest.Interface {
	panic("RESTClient not supported")
}
