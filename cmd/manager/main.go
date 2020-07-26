/*
Copyright 2019 kubeflow.org.

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

package main

import (
	"flag"
	"github.com/kubeflow/kfserving/pkg/webhook/admission/pod"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	"github.com/kubeflow/kfserving/pkg/apis/serving/v1beta1"
	controller "github.com/kubeflow/kfserving/pkg/controller/v1beta1/inferenceservice"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	knservingv1 "knative.dev/serving/pkg/apis/serving/v1"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	flag.Parse()
	logf.SetLogger(logf.ZapLogger(false))
	log := logf.Log.WithName("entrypoint")

	// Create a new Cmd to provide shared dependencies and start components
	log.Info("Setting up manager")
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		LeaderElection:     enableLeaderElection,
		Port:               9443,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	log.Info("Registering Components.")

	// Setup Scheme for all resources
	log.Info("Setting up KFServing scheme")
	if err := v1beta1.AddToScheme(mgr.GetScheme()); err != nil {
		log.Error(err, "unable to add KFServing v1beta1 api to scheme")
		os.Exit(1)
	}

	// Setup Scheme for all resources
	log.Info("Setting up Knative scheme")
	if err := knservingv1.AddToScheme(mgr.GetScheme()); err != nil {
		log.Error(err, "unable to add Knative APIs to scheme")
		os.Exit(1)
	}

	log.Info("Setting up Istio schemes")
	if err := v1alpha3.AddToScheme(mgr.GetScheme()); err != nil {
		log.Error(err, "unable to add Istio v1alpha3 APIs to scheme")
		os.Exit(1)
	}

	// Setup all Controllers
	setupLog.Info("Setting up inference service controller")
	eventBroadcaster := record.NewBroadcaster()
	clientSet, err := kubernetes.NewForConfig(mgr.GetConfig())
	if err != nil {
		setupLog.Error(err, "unable to create clientSet")
		os.Exit(1)
	}
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: clientSet.CoreV1().Events("")})
	if err = (&controller.InferenceServiceReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controllers").WithName("InferenceService"),
		Scheme: mgr.GetScheme(),
		Recorder: eventBroadcaster.NewRecorder(
			mgr.GetScheme(), v1.EventSource{Component: "InferenceServiceControllers"}),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "Controller", "InferenceService")
		os.Exit(1)
	}

	log.Info("setting up webhook server")
	hookServer := mgr.GetWebhookServer()

	log.Info("registering webhooks to the webhook server")
	hookServer.Register("/mutate-pods", &webhook.Admission{Handler: &pod.Mutator{}})
	if err = (&v1beta1.InferenceService{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "v1beta1")
		os.Exit(1)
	}

	// Start the Cmd
	log.Info("Starting the Cmd.")
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		log.Error(err, "unable to run the manager")
		os.Exit(1)
	}
}
