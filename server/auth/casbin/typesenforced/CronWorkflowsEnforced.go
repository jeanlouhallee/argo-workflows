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

type CronWorkflowsEnforced struct {
	delegate v1alpha1.CronWorkflowInterface
	namespace string
	enforcer *casbin.CustomEnforcer
}

const (
	CronWorkflows = "cronworkflows"
)

func (w WorkflowEnforcedInterface) CronWorkflows(namespace string) v1alpha1.CronWorkflowInterface {
	return &CronWorkflowsEnforced{w.delegate.ArgoprojV1alpha1().CronWorkflows(namespace), namespace, casbin.GetCustomEnforcerInstance()}
}

func (w CronWorkflowsEnforced) Create(ctx context.Context, cronWorkflow *wfv1.CronWorkflow, opts metav1.CreateOptions) (*wfv1.CronWorkflow, error) {
	if err := w.enforcer.Enforce(ctx, CronWorkflows, w.namespace, cronWorkflow.GetGenerateName(), casbin.ActionCreate); err != nil {
		return nil, err
	}
	return w.delegate.Create(ctx, cronWorkflow, opts)
}

func (w CronWorkflowsEnforced) Update(ctx context.Context, cronWorkflow *wfv1.CronWorkflow, opts metav1.UpdateOptions) (*wfv1.CronWorkflow, error) {
	if err := w.enforcer.Enforce(ctx, CronWorkflows, w.namespace, cronWorkflow.GetName(), casbin.ActionUpdate); err != nil {
		return nil, err
	}
	return w.delegate.Update(ctx, cronWorkflow, opts)
}

func (w CronWorkflowsEnforced) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	if err := w.enforcer.Enforce(ctx, CronWorkflows, w.namespace, name, casbin.ActionDelete); err != nil {
		return err
	}
	return w.delegate.Delete(ctx, name, opts)
}

func (w CronWorkflowsEnforced) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	if err := w.enforcer.Enforce(ctx, CronWorkflows, w.namespace, "*", casbin.ActionDeleteCollection); err != nil {
		return err
	}
	return w.delegate.DeleteCollection(ctx, opts, listOpts)
}

func (w CronWorkflowsEnforced) Get(ctx context.Context, name string, opts metav1.GetOptions) (*wfv1.CronWorkflow, error) {
	if err := w.enforcer.Enforce(ctx, CronWorkflows, w.namespace, name, casbin.ActionGet); err != nil {
		return nil, err
	}
	return w.delegate.Get(ctx, name, opts)
}

func (w CronWorkflowsEnforced) List(ctx context.Context, opts metav1.ListOptions) (*wfv1.CronWorkflowList, error) {
	if err := w.enforcer.Enforce(ctx, CronWorkflows, w.namespace, "*", casbin.ActionList); err != nil {
		return nil, err
	}
	return w.delegate.List(ctx, opts)
}

func (w CronWorkflowsEnforced) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	if err := w.enforcer.Enforce(ctx, CronWorkflows, w.namespace, "*", casbin.ActionWatch); err != nil {
		return nil, err
	}
	return w.delegate.Watch(ctx, opts)
}

func (w CronWorkflowsEnforced) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *wfv1.CronWorkflow, err error) {
	if err := w.enforcer.Enforce(ctx, CronWorkflows, w.namespace, name, casbin.ActionPatch); err != nil {
		return nil, err
	}
	return w.delegate.Patch(ctx, name, pt, data, opts, subresources...)
}