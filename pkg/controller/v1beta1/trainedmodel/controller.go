/*
Copyright 2020 kubeflow.org.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// +kubebuilder:rbac:groups=serving.kubeflow.org,resources=trainedmodels,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=serving.kubeflow.org,resources=trainedmodels/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=serving.knative.dev,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=serving.knative.dev,resources=services/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;update
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;update;patch;delete
package trainedmodel

import (
	"context"
	"github.com/go-logr/logr"
	v1beta1api "github.com/kubeflow/kfserving/pkg/apis/serving/v1beta1"
	"github.com/kubeflow/kfserving/pkg/controller/v1beta1/trainedmodel/reconcilers/modelconfig"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// TrainedModelReconciler reconciles a TrainedModel object
type TrainedModelReconciler struct {
	client.Client
	Log                   logr.Logger
	Scheme                *runtime.Scheme
	Recorder              record.EventRecorder
	ModelConfigReconciler *modelconfig.ModelConfigReconciler
}

func (r *TrainedModelReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	// Fetch the TrainedModel instance
	tm := &v1beta1api.TrainedModel{}
	if err := r.Get(context.TODO(), req.NamespacedName, tm); err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	// Use finalizer to handle TrainedModel deletion properly
	// When a TrainedModel object is being deleted it should
	// 1) Get its parent InferenceService
	// 2) Find its parent InferenceService model configmap
	// 3) Remove itself from the model configmap
	tmFinalizerName := "trainedmodel.finalizer"

	// examine DeletionTimestamp to determine if object is under deletion
	if tm.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent
		// registering our finalizer.
		if !containsString(tm.GetFinalizers(), tmFinalizerName) {
			tm.SetFinalizers(append(tm.GetFinalizers(), tmFinalizerName))
			if err := r.Update(context.Background(), tm); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else {
		// The object is being deleted
		if containsString(tm.GetFinalizers(), tmFinalizerName) {
			//reconcile configmap to remove the model
			if err := r.ModelConfigReconciler.Reconcile(req, tm); err != nil {
				return reconcile.Result{}, err
			}
			// remove our finalizer from the list and update it.
			tm.SetFinalizers(removeString(tm.GetFinalizers(), tmFinalizerName))
			if err := r.Update(context.Background(), tm); err != nil {
				return ctrl.Result{}, err
			}
		}

		// Stop reconciliation as the item is being deleted
		return ctrl.Result{}, nil
	}

	// Reconcile modelconfig to add this TrainedModel to its parent InferenceService's configmap
	if err := r.ModelConfigReconciler.Reconcile(req, tm); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *TrainedModelReconciler) updateStatus(desiredService *v1beta1api.TrainedModel) error {
	//TODO update TrainedModel status object, this will be done in a separate PR
	return nil
}

func (r *TrainedModelReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1api.TrainedModel{}).
		Complete(r)
}

// Helper functions to check and remove string from a slice of strings.
func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}
