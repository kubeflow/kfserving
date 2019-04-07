package knservice

import (
	"context"
	"fmt"
	"github.com/knative/pkg/kmp"
	knservingv1alpha1 "github.com/knative/serving/pkg/apis/serving/v1alpha1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("ServiceReconciler")

type ServiceReconciler struct {
	client client.Client
}

func NewServiceReconcile(client client.Client) *ServiceReconciler {
	return &ServiceReconciler{
		client: client,
	}
}

// Reconcile compares the actual state with the desired, and attempts to
// converge the two. It then updates the Status block of the Service resource
// with the current status of the resource.
func (c *ServiceReconciler) Reconcile(ctx context.Context, desiredService *knservingv1alpha1.Service) (*knservingv1alpha1.Service, error) {
	service := &knservingv1alpha1.Service{}
	err := c.client.Get(context.TODO(), types.NamespacedName{Name: desiredService.Name, Namespace: desiredService.Namespace}, service)
	if err != nil && errors.IsNotFound(err) {
		log.Info("Creating service", "namespace", service.Namespace, "name", service.Name)
		err = c.client.Create(context.TODO(), desiredService)
		return nil, err
	} else if err != nil {
		return nil, err
	}

	if serviceSemanticEquals(desiredService, service) {
		// No differences to reconcile.
		return service, nil
	}
	if service.Spec.Release != nil {
		service.Spec.Release.Configuration.DeprecatedGeneration = desiredService.Spec.Release.Configuration.DeprecatedGeneration
	}
	diff, err := kmp.SafeDiff(desiredService.Spec, service.Spec)
	if err != nil {
		return service, fmt.Errorf("failed to diff service: %v", err)
	}
	log.Info("Reconciling service diff (-desired, +observed): %s", diff)

	service.Spec = desiredService.Spec
	log.Info("Updating service", "namespace", service.Namespace, "name", service.Name)
	err = c.client.Update(context.TODO(), service)
	if err != nil {
		return service, err
	}
	return service, nil
}

func serviceSemanticEquals(desiredService, service *knservingv1alpha1.Service) bool {
	return equality.Semantic.DeepEqual(desiredService.Spec, service.Spec) &&
		equality.Semantic.DeepEqual(desiredService.ObjectMeta.Labels, service.ObjectMeta.Labels)
}
