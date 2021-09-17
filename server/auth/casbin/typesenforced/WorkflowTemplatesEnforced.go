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

type WorkflowTemplatesEnforced struct {
	delegate v1alpha1.WorkflowTemplateInterface
	namespace string
	enforcer *casbin.CustomEnforcer
}

const (
	WorkflowsTemplates = "workflowtemplates"
)


func (w WorkflowEnforcedInterface) WorkflowTemplates(namespace string) v1alpha1.WorkflowTemplateInterface {
	return &WorkflowTemplatesEnforced{w.delegate.ArgoprojV1alpha1().WorkflowTemplates(namespace), namespace, casbin.GetCustomEnforcerInstance()}
}

func (w WorkflowTemplatesEnforced) Create(ctx context.Context, workflowTemplate *wfv1.WorkflowTemplate, opts metav1.CreateOptions) (*wfv1.WorkflowTemplate, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowsTemplates, w.namespace, workflowTemplate.GetName(), casbin.ActionCreate); err != nil {
		return nil, err
	}
	return w.delegate.Create(ctx, workflowTemplate, opts)
}

func (w WorkflowTemplatesEnforced) Update(ctx context.Context, workflowTemplate *wfv1.WorkflowTemplate, opts metav1.UpdateOptions) (*wfv1.WorkflowTemplate, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowsTemplates, w.namespace, workflowTemplate.GetName(), casbin.ActionUpdate); err != nil {
		return nil, err
	}
	return w.delegate.Update(ctx, workflowTemplate, opts)
}

func (w WorkflowTemplatesEnforced) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	if err := w.enforcer.Enforce(ctx, WorkflowsTemplates, w.namespace, name, casbin.ActionDelete); err != nil {
		return nil
	}
	return w.delegate.Delete(ctx, name, opts)
}

func (w WorkflowTemplatesEnforced) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	if err := w.enforcer.Enforce(ctx, WorkflowsTemplates, w.namespace, "*", casbin.ActionDeleteCollection); err != nil {
		return nil
	}
	return w.delegate.DeleteCollection(ctx, opts, listOpts)
}

func (w WorkflowTemplatesEnforced) Get(ctx context.Context, name string, opts metav1.GetOptions) (*wfv1.WorkflowTemplate, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowsTemplates, w.namespace, name, casbin.ActionGet); err != nil {
		return nil, err
	}
	return w.delegate.Get(ctx, name, opts)
}

func (w WorkflowTemplatesEnforced) List(ctx context.Context, opts metav1.ListOptions) (*wfv1.WorkflowTemplateList, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowsTemplates, w.namespace, "*", casbin.ActionList); err != nil {
		return nil, err
	}
	return w.delegate.List(ctx, opts)
}

func (w WorkflowTemplatesEnforced) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	if err := w.enforcer.Enforce(ctx, WorkflowsTemplates, w.namespace, "*", casbin.ActionWatch); err != nil {
		return nil, err
	}
	return w.delegate.Watch(ctx, opts)
}

func (w WorkflowTemplatesEnforced) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *wfv1.WorkflowTemplate, err error) {
	if err := w.enforcer.Enforce(ctx, WorkflowsTemplates, w.namespace, name, casbin.ActionWatch); err != nil {
		return nil, err
	}
	return w.delegate.Patch(ctx, name, pt, data, opts, subresources...)
}