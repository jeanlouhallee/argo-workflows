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

type WorkflowEventBindingsEnforced struct {
	delegate v1alpha1.WorkflowEventBindingInterface
	namespace string
	enforcer *casbin.CustomEnforcer
}

var (
	WorkflowEventBindings = "workfloweventbinding"
)

func (w WorkflowEnforcedInterface) WorkflowEventBindings(namespace string) v1alpha1.WorkflowEventBindingInterface {
	return &WorkflowEventBindingsEnforced{w.delegate.ArgoprojV1alpha1().WorkflowEventBindings(namespace), namespace, casbin.GetCustomEnforcerInstance()}
}

func (w WorkflowEventBindingsEnforced) Create(ctx context.Context, workflowEventBinding *wfv1.WorkflowEventBinding, opts metav1.CreateOptions) (*wfv1.WorkflowEventBinding, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowEventBindings, w.namespace, workflowEventBinding.GetGenerateName(), casbin.ActionCreate); err != nil {
		return nil, err
	}
	return w.delegate.Create(ctx, workflowEventBinding, opts)
}

func (w WorkflowEventBindingsEnforced) Update(ctx context.Context, workflowEventBinding *wfv1.WorkflowEventBinding, opts metav1.UpdateOptions) (*wfv1.WorkflowEventBinding, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowEventBindings, w.namespace, workflowEventBinding.GetName(), casbin.ActionUpdate); err != nil {
		return nil, err
	}
	return w.delegate.Update(ctx, workflowEventBinding, opts)
}

func (w WorkflowEventBindingsEnforced) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	if err := w.enforcer.Enforce(ctx, WorkflowEventBindings, w.namespace, name, casbin.ActionDelete); err != nil {
		return nil
	}
	return w.delegate.Delete(ctx, name, opts)
}

func (w WorkflowEventBindingsEnforced) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	if err := w.enforcer.Enforce(ctx, WorkflowEventBindings, w.namespace, "*", casbin.ActionDeleteCollection); err != nil {
		return nil
	}
	return w.delegate.DeleteCollection(ctx, opts, listOpts)
}

func (w WorkflowEventBindingsEnforced) Get(ctx context.Context, name string, opts metav1.GetOptions) (*wfv1.WorkflowEventBinding, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowEventBindings, w.namespace, name, casbin.ActionGet); err != nil {
		return nil, err
	}
	return w.delegate.Get(ctx, name, opts)
}

func (w WorkflowEventBindingsEnforced) List(ctx context.Context, opts metav1.ListOptions) (*wfv1.WorkflowEventBindingList, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowEventBindings, w.namespace, "*", casbin.ActionList); err != nil {
		return nil, err
	}
	return w.delegate.List(ctx, opts)
}

func (w WorkflowEventBindingsEnforced) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowEventBindings, w.namespace, "*", casbin.ActionWatch); err != nil {
		return nil, err
	}
	return w.delegate.Watch(ctx, opts)
}

func (w WorkflowEventBindingsEnforced) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *wfv1.WorkflowEventBinding, err error) {
	if err := w.enforcer.Enforce(ctx, WorkflowEventBindings, w.namespace, name, casbin.ActionPatch); err != nil {
		return nil, err
	}
	return w.delegate.Patch(ctx, name, pt, data, opts, subresources...)
}