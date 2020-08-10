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

package v1beta1

import (
	"context"
	"encoding/json"
	"fmt"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	"github.com/kubeflow/kfserving/pkg/constants"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ConfigMap Keys
const (
	PredictorConfigKeyName   = "predictors"
	TransformerConfigKeyName = "transformers"
	ExplainerConfigKeyName   = "explainers"
)

// +kubebuilder:object:generate=false
type ExplainerConfig struct {
	// explainer docker image name
	ContainerImage string `json:"image"`
	// default explainer docker image version
	DefaultImageVersion string `json:"defaultImageVersion"`
}

// +kubebuilder:object:generate=false
type ExplainersConfig struct {
	AlibiExplainer ExplainerConfig `json:"alibi,omitempty"`
}

// +kubebuilder:object:generate=false
type PredictorConfig struct {
	// predictor docker image name
	ContainerImage string `json:"image"`
	// default predictor docker image version on cpu
	DefaultImageVersion string `json:"defaultImageVersion"`
	// default predictor docker image version on gpu
	DefaultGpuImageVersion string `json:"defaultGpuImageVersion"`
}

// +kubebuilder:object:generate=false
type PredictorsConfig struct {
	Tensorflow PredictorConfig `json:"tensorflow,omitempty"`
	Triton     PredictorConfig `json:"triton,omitempty"`
	XGBoost    PredictorConfig `json:"xgboost,omitempty"`
	SKlearn    PredictorConfig `json:"sklearn,omitempty"`
	PyTorch    PredictorConfig `json:"pytorch,omitempty"`
	ONNX       PredictorConfig `json:"onnx,omitempty"`
}

// +kubebuilder:object:generate=false
type TransformerConfig struct {
	// transformer docker image name
	ContainerImage string `json:"image"`
	// default transformer docker image version
	DefaultImageVersion string `json:"defaultImageVersion"`
}

// +kubebuilder:object:generate=false
type TransformersConfig struct {
	Feast TransformerConfig `json:"feast,omitempty"`
}

// +kubebuilder:object:generate=false
type InferenceServicesConfig struct {
	// Transformer configurations
	Transformers TransformersConfig `json:"transformers"`
	// Predictor configurations
	Predictors PredictorsConfig `json:"predictors"`
	// Explainer configurations
	Explainers ExplainersConfig `json:"explainers"`
}

func NewInferenceServicesConfig() (*InferenceServicesConfig, error) {
	cli, err := client.New(config.GetConfigOrDie(), client.Options{})
	if err != nil {
		return nil, err
	}
	configMap := &v1.ConfigMap{}
	err = cli.Get(context.TODO(), types.NamespacedName{Name: constants.InferenceServiceConfigMapName, Namespace: constants.KFServingNamespace}, configMap)
	if err != nil {
		return nil, err
	}

	predictorsConfig, err := getPredictorsConfigs(configMap)
	if err != nil {
		return nil, err
	}
	transformersConfig, err := getTransformersConfigs(configMap)
	if err != nil {
		return nil, err
	}
	explainersConfig, err := getExplainersConfigs(configMap)
	if err != nil {
		return nil, err
	}
	return &InferenceServicesConfig{
		Predictors:   *predictorsConfig,
		Transformers: *transformersConfig,
		Explainers:   *explainersConfig,
	}, nil
}

func getPredictorsConfigs(configMap *v1.ConfigMap) (*PredictorsConfig, error) {
	predictorConfig := &PredictorsConfig{}
	if data, ok := configMap.Data[PredictorConfigKeyName]; ok {
		err := json.Unmarshal([]byte(data), &predictorConfig)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshall %v json string due to %v ", PredictorConfigKeyName, err)
		}
	}
	return predictorConfig, nil
}

func getTransformersConfigs(configMap *v1.ConfigMap) (*TransformersConfig, error) {
	transformerConfig := &TransformersConfig{}
	if data, ok := configMap.Data[TransformerConfigKeyName]; ok {
		err := json.Unmarshal([]byte(data), &transformerConfig)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshall %v json string due to %v ", TransformerConfigKeyName, err)
		}
	}
	return transformerConfig, nil
}

func getExplainersConfigs(configMap *v1.ConfigMap) (*ExplainersConfig, error) {
	explainerConfig := &ExplainersConfig{}
	if data, ok := configMap.Data[ExplainerConfigKeyName]; ok {
		err := json.Unmarshal([]byte(data), &explainerConfig)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshall %v json string due to %v ", ExplainerConfigKeyName, err)
		}
	}
	return explainerConfig, nil
}
