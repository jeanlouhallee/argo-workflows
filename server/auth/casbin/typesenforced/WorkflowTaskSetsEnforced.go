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

type WorkflowTaskSetsEnforced struct {
	delegate v1alpha1.WorkflowTaskSetInterface
	namespace string
	enforcer *casbin.CustomEnforcer
}

var (
	WorkflowTaskSets = "workflowtasksets"
)

func (w WorkflowEnforcedInterface) WorkflowTaskSets(namespace string) v1alpha1.WorkflowTaskSetInterface {
	return &WorkflowTaskSetsEnforced{w.delegate.ArgoprojV1alpha1().WorkflowTaskSets(namespace), namespace, casbin.GetCustomEnforcerInstance()}
}

func (w WorkflowTaskSetsEnforced) Create(ctx context.Context, workflowTaskSet *wfv1.WorkflowTaskSet, opts metav1.CreateOptions) (*wfv1.WorkflowTaskSet, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowTaskSets, w.namespace, workflowTaskSet.GetName(), casbin.ActionCreate); err != nil {
		return nil, err
	}
	return w.delegate.Create(ctx, workflowTaskSet, opts)
}

func (w WorkflowTaskSetsEnforced) Update(ctx context.Context, workflowTaskSet *wfv1.WorkflowTaskSet, opts metav1.UpdateOptions) (*wfv1.WorkflowTaskSet, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowTaskSets, w.namespace, workflowTaskSet.GetName(), casbin.ActionUpdate); err != nil {
		return nil, err
	}
	return w.delegate.Update(ctx, workflowTaskSet, opts)
}

func (w WorkflowTaskSetsEnforced) UpdateStatus(ctx context.Context, workflowTaskSet *wfv1.WorkflowTaskSet, opts metav1.UpdateOptions) (*wfv1.WorkflowTaskSet, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowTaskSets, w.namespace, workflowTaskSet.GetName(), "UpdateStatus"); err != nil {
		return nil, err
	}
	return w.delegate.UpdateStatus(ctx, workflowTaskSet, opts)
}

func (w WorkflowTaskSetsEnforced) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	if err := w.enforcer.Enforce(ctx, WorkflowTaskSets, w.namespace, name, casbin.ActionDelete); err != nil {
		return nil
	}
	return w.delegate.Delete(ctx, name, opts)
}

func (w WorkflowTaskSetsEnforced) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	if err := w.enforcer.Enforce(ctx, WorkflowTaskSets, w.namespace, "*", casbin.ActionDeleteCollection); err != nil {
		return nil
	}
	return w.delegate.DeleteCollection(ctx, opts, listOpts)
}

func (w WorkflowTaskSetsEnforced) Get(ctx context.Context, name string, opts metav1.GetOptions) (*wfv1.WorkflowTaskSet, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowTaskSets, w.namespace,name, casbin.ActionGet); err != nil {
		return nil, err
	}
	return w.delegate.Get(ctx, name, opts)
}

func (w WorkflowTaskSetsEnforced) List(ctx context.Context, opts metav1.ListOptions) (*wfv1.WorkflowTaskSetList, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowTaskSets, w.namespace, "*", casbin.ActionList); err != nil {
		return nil, err
	}
	return w.delegate.List(ctx, opts)
}

func (w WorkflowTaskSetsEnforced) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowTaskSets, w.namespace, "*", casbin.ActionWatch); err != nil {
		return nil, err
	}
	return w.delegate.Watch(ctx, opts)
}

func (w WorkflowTaskSetsEnforced) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *wfv1.WorkflowTaskSet, err error) {
	if err := w.enforcer.Enforce(ctx, WorkflowTaskSets, w.namespace, name, casbin.ActionPatch); err != nil {
		return nil, err
	}
	return w.delegate.Patch(ctx, name, pt, data, opts, subresources...)
}