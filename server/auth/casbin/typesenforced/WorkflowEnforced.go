package typesenforced

import (
	"context"
	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned/typed/workflow/v1alpha1"
	"github.com/argoproj/argo-workflows/v3/server/auth/casbin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
)

const (
	Workflows = "workflows"
)

type WorkflowEnforced struct {
	delegate  v1alpha1.WorkflowInterface
	namespace string
	enforcer *casbin.CustomEnforcer
}

func (w WorkflowEnforcedInterface) Workflows(namespace string) v1alpha1.WorkflowInterface {
	return &WorkflowEnforced{w.delegate.ArgoprojV1alpha1().Workflows(namespace), namespace, casbin.GetCustomEnforcerInstance()}
}

func (w WorkflowEnforced) Create(ctx context.Context, workflow *wfv1.Workflow, opts metav1.CreateOptions) (*wfv1.Workflow, error) {
	if err := w.enforcer.Enforce(ctx, Workflows, w.namespace, workflow.GetGenerateName(), casbin.ActionCreate); err != nil {
		return nil, err
	}
	return w.delegate.Create(ctx, workflow, opts)
}

func (w WorkflowEnforced) Update(ctx context.Context, workflow *wfv1.Workflow, opts metav1.UpdateOptions) (*wfv1.Workflow, error) {
	if err := w.enforcer.Enforce(ctx, Workflows, w.namespace, workflow.GetName(), casbin.ActionUpdate); err != nil {
		return nil, err
	}
	return w.delegate.Update(ctx, workflow, opts)
}

func (w WorkflowEnforced) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	if err := w.enforcer.Enforce(ctx, Workflows, w.namespace, name, casbin.ActionDelete); err != nil {
		return err
	}
	return w.delegate.Delete(ctx, name, opts)
}

func (w WorkflowEnforced) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	if err := w.enforcer.Enforce(ctx, Workflows, w.namespace, "*", casbin.ActionDeleteCollection); err != nil {
		return err
	}
	return w.delegate.DeleteCollection(ctx, opts, listOpts)
}

func (w WorkflowEnforced) Get(ctx context.Context, name string, opts metav1.GetOptions) (*wfv1.Workflow, error) {
	if err := w.enforcer.Enforce(ctx, Workflows, w.namespace, name, casbin.ActionGet); err != nil {
		return nil, err
	}
	return w.delegate.Get(ctx, name, opts)
}

func (w WorkflowEnforced) List(ctx context.Context, opts metav1.ListOptions) (*wfv1.WorkflowList, error) {
	if err := w.enforcer.Enforce(ctx, Workflows, w.namespace, "*", casbin.ActionList); err != nil {
		return nil, err
	}
	return w.delegate.List(ctx, opts)
}

func (w WorkflowEnforced) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	if err := w.enforcer.Enforce(ctx, Workflows, w.namespace, "*", casbin.ActionWatch); err != nil {
		return nil, err
	}
	return w.delegate.Watch(ctx, opts)
}

func (w WorkflowEnforced) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *wfv1.Workflow, err error) {
	if err := w.enforcer.Enforce(ctx, Workflows, w.namespace, name, casbin.ActionPatch); err != nil {
		return nil, err
	}
	return w.delegate.Patch(ctx, name, pt, data, opts, subresources...)
}