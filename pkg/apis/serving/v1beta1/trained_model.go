package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TrainedModel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrainedModelSpec   `json:"spec,omitempty"`
	Status            TrainedModelStatus `json:"status,omitempty"`
}
type TrainedModelSpec struct {
	// Required field for parent inference service
	InferenceService string `json:"inferenceService"`
	// Predictor model spec
	PredictorModel ModelSpec `json:"predictorModel"`
	// Explainer model spec
	ExplainerModel ModelSpec `json:"explainerModel,omitempty"`
}
type ModelSpec struct {
	// Storage URI for the model repository
	StorageURI string `json:"storageUri"`
	// Machine Learning <framework name>:<git tag>
	// The values could be: "tensorflow:v2.2.0","pytorch:v1.5.1","sklearn:0.23.1","onnx:v1.7.0","xgboost:v1.1.1", "myawesomeinternalframework:1.1.0" etc.
	Framework string `json:"framework"`
	// Machine Learning model type, this field is used to match explainer type
	// +kubebuilder:webhooks:Enum={"tabular","text","image"}
	Type string `json:"type,omitempty"`
	//Defaults to requests and limits of 1CPU, 2Gb MEM.
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
}