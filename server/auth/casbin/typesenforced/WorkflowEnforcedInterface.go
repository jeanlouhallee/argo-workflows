package typesenforced

import (
	workflow "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
	"github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned/typed/workflow/v1alpha1"
	"k8s.io/client-go/discovery"
)

type WorkflowEnforcedInterface struct {
	delegate workflow.Interface
}

func (w WorkflowEnforcedInterface) Discovery() discovery.DiscoveryInterface {
	panic("Discovery not supported")
}

func (w WorkflowEnforcedInterface) ArgoprojV1alpha1() v1alpha1.ArgoprojV1alpha1Interface {
	return w
}

func WrapWorkflowInterface(w workflow.Interface) workflow.Interface {
	return &WorkflowEnforcedInterface{w}
}
