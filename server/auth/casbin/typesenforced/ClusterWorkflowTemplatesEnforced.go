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
	ClusterWorkflowTemplates = "clusterworkflowtemplates"
)

type ClusterWorkflowTemplatesEnforced struct {
	delegate v1alpha1.ClusterWorkflowTemplateInterface
	enforcer *casbin.CustomEnforcer
}

func (w WorkflowEnforcedInterface) ClusterWorkflowTemplates() v1alpha1.ClusterWorkflowTemplateInterface {
	return &ClusterWorkflowTemplatesEnforced{w.delegate.ArgoprojV1alpha1().ClusterWorkflowTemplates(), casbin.GetCustomEnforcerInstance()}
}

func (w ClusterWorkflowTemplatesEnforced) Create(ctx context.Context, clusterWorkflowTemplate *wfv1.ClusterWorkflowTemplate, opts metav1.CreateOptions) (*wfv1.ClusterWorkflowTemplate, error) {
	if err := w.enforcer.Enforce(ctx, ClusterWorkflowTemplates, "*", clusterWorkflowTemplate.GetName(), casbin.ActionCreate); err != nil {
		return nil, err
	}
	return w.delegate.Create(ctx, clusterWorkflowTemplate, opts)
}

func (w ClusterWorkflowTemplatesEnforced) Update(ctx context.Context, clusterWorkflowTemplate *wfv1.ClusterWorkflowTemplate, opts metav1.UpdateOptions) (*wfv1.ClusterWorkflowTemplate, error) {
	if err := w.enforcer.Enforce(ctx, ClusterWorkflowTemplates, "*", clusterWorkflowTemplate.GetName(), casbin.ActionUpdate); err != nil {
		return nil, err
	}
	return w.delegate.Update(ctx, clusterWorkflowTemplate, opts)
}

func (w ClusterWorkflowTemplatesEnforced) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	if err := w.enforcer.Enforce(ctx, ClusterWorkflowTemplates, "*", name, casbin.ActionDelete); err != nil {
		return nil
	}
	return w.delegate.Delete(ctx, name, opts)
}

func (w ClusterWorkflowTemplatesEnforced) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	if err := w.enforcer.Enforce(ctx, ClusterWorkflowTemplates, "*", "*", casbin.ActionDeleteCollection); err != nil {
		return nil
	}
	return w.delegate.DeleteCollection(ctx, opts, listOpts)
}

func (w ClusterWorkflowTemplatesEnforced) Get(ctx context.Context, name string, opts metav1.GetOptions) (*wfv1.ClusterWorkflowTemplate, error) {
	if err := w.enforcer.Enforce(ctx, ClusterWorkflowTemplates, "*", name, casbin.ActionGet); err != nil {
		return nil, err
	}
	return w.delegate.Get(ctx, name, opts)
}

func (w ClusterWorkflowTemplatesEnforced) List(ctx context.Context, opts metav1.ListOptions) (*wfv1.ClusterWorkflowTemplateList, error) {
	if err := w.enforcer.Enforce(ctx, ClusterWorkflowTemplates, "*", opts.FieldSelector, casbin.ActionList); err != nil {
		return nil, err
	}
	return w.delegate.List(ctx, opts)
}

func (w ClusterWorkflowTemplatesEnforced) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	if err := w.enforcer.Enforce(ctx, ClusterWorkflowTemplates, "*", opts.FieldSelector, casbin.ActionWatch); err != nil {
		return nil, err
	}
	return w.delegate.Watch(ctx, opts)
}

func (w ClusterWorkflowTemplatesEnforced) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *wfv1.ClusterWorkflowTemplate, err error) {
	if err := w.enforcer.Enforce(ctx, ClusterWorkflowTemplates, "*", name, casbin.ActionPatch); err != nil {
		return nil, err
	}
	return w.delegate.Patch(ctx, name, pt, data, opts, subresources...)
}