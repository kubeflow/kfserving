// +build !ignore_autogenerated

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

// Code generated by main. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha2

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.AlibiExplainerSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "AlibiExplainerSpec defines the arguments for configuring an Alibi Explanation Server",
					Properties: map[string]spec.Schema{
						"type": {
							SchemaProps: spec.SchemaProps{
								Description: "The type of Alibi explainer",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"storageUri": {
							SchemaProps: spec.SchemaProps{
								Description: "The location of a trained explanation model",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"runtimeVersion": {
							SchemaProps: spec.SchemaProps{
								Description: "Defaults to latest Alibi Version.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"resources": {
							SchemaProps: spec.SchemaProps{
								Description: "Defaults to requests and limits of 1CPU, 2Gb MEM.",
								Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
							},
						},
						"config": {
							SchemaProps: spec.SchemaProps{
								Description: "Inline custom parameter settings for explainer",
								Type:        []string{"object"},
								AdditionalProperties: &spec.SchemaOrBool{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Type:   []string{"string"},
											Format: "",
										},
									},
								},
							},
						},
					},
					Required: []string{"type"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.ResourceRequirements"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.CustomSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "CustomSpec provides a hook for arbitrary container configuration.",
					Properties: map[string]spec.Schema{
						"container": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/api/core/v1.Container"),
							},
						},
					},
					Required: []string{"container"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.Container"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.DeploymentSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "DeploymentSpec defines the configuration for a given KFService service component",
					Properties: map[string]spec.Schema{
						"serviceAccountName": {
							SchemaProps: spec.SchemaProps{
								Description: "ServiceAccountName is the name of the ServiceAccount to use to run the service",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"minReplicas": {
							SchemaProps: spec.SchemaProps{
								Description: "Minimum number of replicas, pods won't scale down to 0 in case of no traffic",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
						"maxReplicas": {
							SchemaProps: spec.SchemaProps{
								Description: "This is the up bound for autoscaler to scale to",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
					},
				},
			},
			Dependencies: []string{},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.EndpointSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"predictor": {
							SchemaProps: spec.SchemaProps{
								Description: "Predictor defines the model serving spec",
								Ref:         ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PredictorSpec"),
							},
						},
						"explainer": {
							SchemaProps: spec.SchemaProps{
								Description: "Explainer defines the model explanation service spec explainer service calls to transformer or predictor service",
								Ref:         ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ExplainerSpec"),
							},
						},
						"transformer": {
							SchemaProps: spec.SchemaProps{
								Description: "Transformer defines the transformer service spec for pre/post processing transformer service calls to predictor service",
								Ref:         ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TransformerSpec"),
							},
						},
					},
					Required: []string{"predictor"},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ExplainerSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PredictorSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TransformerSpec"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ExplainerConfig": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"image": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
					},
					Required: []string{"image"},
				},
			},
			Dependencies: []string{},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ExplainerSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "ExplainerSpec defines the arguments for a model explanation server",
					Properties: map[string]spec.Schema{
						"alibi": {
							SchemaProps: spec.SchemaProps{
								Description: "The following fields follow a \"1-of\" semantic. Users must specify exactly one spec.",
								Ref:         ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.AlibiExplainerSpec"),
							},
						},
						"custom": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.CustomSpec"),
							},
						},
						"serviceAccountName": {
							SchemaProps: spec.SchemaProps{
								Description: "ServiceAccountName is the name of the ServiceAccount to use to run the service",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"minReplicas": {
							SchemaProps: spec.SchemaProps{
								Description: "Minimum number of replicas, pods won't scale down to 0 in case of no traffic",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
						"maxReplicas": {
							SchemaProps: spec.SchemaProps{
								Description: "This is the up bound for autoscaler to scale to",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
					},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.AlibiExplainerSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.CustomSpec"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ExplainersConfig": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"alibi": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ExplainerConfig"),
							},
						},
					},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ExplainerConfig"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.FeastTransformerSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "FeastTransformerSpec defines arguments for configuring a Transformer to call Feast",
					Properties: map[string]spec.Schema{
						"feastUrl": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"dataType": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"entityIds": {
							SchemaProps: spec.SchemaProps{
								Type: []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Type:   []string{"string"},
											Format: "",
										},
									},
								},
							},
						},
						"featureIds": {
							SchemaProps: spec.SchemaProps{
								Type: []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Type:   []string{"string"},
											Format: "",
										},
									},
								},
							},
						},
					},
				},
			},
			Dependencies: []string{},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFService": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "KFService is the Schema for the services API",
					Properties: map[string]spec.Schema{
						"kind": {
							SchemaProps: spec.SchemaProps{
								Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"apiVersion": {
							SchemaProps: spec.SchemaProps{
								Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"metadata": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
							},
						},
						"spec": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFServiceSpec"),
							},
						},
						"status": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFServiceStatus"),
							},
						},
					},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFServiceSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFServiceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFServiceList": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "KFServiceList contains a list of Service",
					Properties: map[string]spec.Schema{
						"kind": {
							SchemaProps: spec.SchemaProps{
								Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"apiVersion": {
							SchemaProps: spec.SchemaProps{
								Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"metadata": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
							},
						},
						"items": {
							SchemaProps: spec.SchemaProps{
								Type: []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFService"),
										},
									},
								},
							},
						},
					},
					Required: []string{"items"},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFService", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFServiceSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "KFServiceSpec defines the desired state of KFService",
					Properties: map[string]spec.Schema{
						"default": {
							SchemaProps: spec.SchemaProps{
								Description: "Default defines default KFService endpoints",
								Ref:         ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.EndpointSpec"),
							},
						},
						"canary": {
							SchemaProps: spec.SchemaProps{
								Description: "Canary defines an alternate endpoints to route a percentage of traffic.",
								Ref:         ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.EndpointSpec"),
							},
						},
						"canaryTrafficPercent": {
							SchemaProps: spec.SchemaProps{
								Description: "CanaryTrafficPercent defines the percentage of traffic going to canary KFService endpoints",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
					},
					Required: []string{"default"},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.EndpointSpec"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFServiceStatus": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "KFServiceStatus defines the observed state of KFService",
					Properties: map[string]spec.Schema{
						"observedGeneration": {
							SchemaProps: spec.SchemaProps{
								Description: "ObservedGeneration is the 'Generation' of the Service that was last processed by the controller.",
								Type:        []string{"integer"},
								Format:      "int64",
							},
						},
						"conditions": {
							VendorExtensible: spec.VendorExtensible{
								Extensions: spec.Extensions{
									"x-kubernetes-patch-merge-key": "type",
									"x-kubernetes-patch-strategy":  "merge",
								},
							},
							SchemaProps: spec.SchemaProps{
								Description: "Conditions the latest available observations of a resource's current state.",
								Type:        []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Ref: ref("knative.dev/pkg/apis.Condition"),
										},
									},
								},
							},
						},
						"url": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"default": {
							SchemaProps: spec.SchemaProps{
								Type: []string{"object"},
								AdditionalProperties: &spec.SchemaOrBool{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.StatusConfigurationSpec"),
										},
									},
								},
							},
						},
						"canary": {
							SchemaProps: spec.SchemaProps{
								Type: []string{"object"},
								AdditionalProperties: &spec.SchemaOrBool{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.StatusConfigurationSpec"),
										},
									},
								},
							},
						},
					},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.StatusConfigurationSpec", "knative.dev/pkg/apis.Condition"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ONNXSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "ONNXSpec defines arguments for configuring ONNX model serving.",
					Properties: map[string]spec.Schema{
						"storageUri": {
							SchemaProps: spec.SchemaProps{
								Description: "The location of the trained model",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"runtimeVersion": {
							SchemaProps: spec.SchemaProps{
								Description: "Defaults to latest ONNX Version.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"resources": {
							SchemaProps: spec.SchemaProps{
								Description: "Defaults to requests and limits of 1CPU, 2Gb MEM.",
								Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
							},
						},
					},
					Required: []string{"storageUri"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.ResourceRequirements"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PredictorConfig": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"image": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
					},
					Required: []string{"image"},
				},
			},
			Dependencies: []string{},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PredictorSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "PredictorSpec defines the configuration to route traffic to a predictor.",
					Properties: map[string]spec.Schema{
						"custom": {
							SchemaProps: spec.SchemaProps{
								Description: "The following fields follow a \"1-of\" semantic. Users must specify exactly one spec.",
								Ref:         ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.CustomSpec"),
							},
						},
						"tensorflow": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TensorflowSpec"),
							},
						},
						"tensorrt": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TensorRTSpec"),
							},
						},
						"xgboost": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.XGBoostSpec"),
							},
						},
						"sklearn": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.SKLearnSpec"),
							},
						},
						"onnx": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ONNXSpec"),
							},
						},
						"pytorch": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PyTorchSpec"),
							},
						},
						"serviceAccountName": {
							SchemaProps: spec.SchemaProps{
								Description: "ServiceAccountName is the name of the ServiceAccount to use to run the service",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"minReplicas": {
							SchemaProps: spec.SchemaProps{
								Description: "Minimum number of replicas, pods won't scale down to 0 in case of no traffic",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
						"maxReplicas": {
							SchemaProps: spec.SchemaProps{
								Description: "This is the up bound for autoscaler to scale to",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
					},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.CustomSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ONNXSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PyTorchSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.SKLearnSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TensorRTSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TensorflowSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.XGBoostSpec"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PredictorsConfig": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"tensorflow": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PredictorConfig"),
							},
						},
						"tensorrt": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PredictorConfig"),
							},
						},
						"xgboost": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PredictorConfig"),
							},
						},
						"sklearn": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PredictorConfig"),
							},
						},
						"pytorch": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PredictorConfig"),
							},
						},
						"onnx": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PredictorConfig"),
							},
						},
					},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PredictorConfig"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PyTorchSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "PyTorchSpec defines arguments for configuring PyTorch model serving.",
					Properties: map[string]spec.Schema{
						"storageUri": {
							SchemaProps: spec.SchemaProps{
								Description: "The location of the trained model",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"modelClassName": {
							SchemaProps: spec.SchemaProps{
								Description: "Defaults PyTorch model class name to 'PyTorchModel'",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"runtimeVersion": {
							SchemaProps: spec.SchemaProps{
								Description: "Defaults to latest PyTorch Version",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"resources": {
							SchemaProps: spec.SchemaProps{
								Description: "Defaults to requests and limits of 1CPU, 2Gb MEM.",
								Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
							},
						},
					},
					Required: []string{"storageUri"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.ResourceRequirements"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.SKLearnSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "SKLearnSpec defines arguments for configuring SKLearn model serving.",
					Properties: map[string]spec.Schema{
						"storageUri": {
							SchemaProps: spec.SchemaProps{
								Description: "The location of the trained model",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"runtimeVersion": {
							SchemaProps: spec.SchemaProps{
								Description: "Defaults to latest SKLearn Version.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"resources": {
							SchemaProps: spec.SchemaProps{
								Description: "Defaults to requests and limits of 1CPU, 2Gb MEM.",
								Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
							},
						},
					},
					Required: []string{"storageUri"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.ResourceRequirements"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.StatusConfigurationSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "StatusConfigurationSpec describes the state of the configuration receiving traffic.",
					Properties: map[string]spec.Schema{
						"name": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"host": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"replicas": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"integer"},
								Format: "int32",
							},
						},
						"traffic": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"integer"},
								Format: "int32",
							},
						},
					},
				},
			},
			Dependencies: []string{},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TensorRTSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "TensorRTSpec defines arguments for configuring TensorRT model serving.",
					Properties: map[string]spec.Schema{
						"storageUri": {
							SchemaProps: spec.SchemaProps{
								Description: "The location of the trained model",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"runtimeVersion": {
							SchemaProps: spec.SchemaProps{
								Description: "Defaults to latest TensorRT Version.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"resources": {
							SchemaProps: spec.SchemaProps{
								Description: "Defaults to requests and limits of 1CPU, 2Gb MEM.",
								Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
							},
						},
					},
					Required: []string{"storageUri"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.ResourceRequirements"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TensorflowSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "TensorflowSpec defines arguments for configuring Tensorflow model serving.",
					Properties: map[string]spec.Schema{
						"storageUri": {
							SchemaProps: spec.SchemaProps{
								Description: "The location of the trained model",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"runtimeVersion": {
							SchemaProps: spec.SchemaProps{
								Description: "Defaults to latest TF Version.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"resources": {
							SchemaProps: spec.SchemaProps{
								Description: "Defaults to requests and limits of 1CPU, 2Gb MEM.",
								Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
							},
						},
					},
					Required: []string{"storageUri"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.ResourceRequirements"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TransformerSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "TransformerSpec defines transformer service for pre/post processing",
					Properties: map[string]spec.Schema{
						"custom": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.CustomSpec"),
							},
						},
						"feast": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.FeastTransformerSpec"),
							},
						},
						"serviceAccountName": {
							SchemaProps: spec.SchemaProps{
								Description: "ServiceAccountName is the name of the ServiceAccount to use to run the service",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"minReplicas": {
							SchemaProps: spec.SchemaProps{
								Description: "Minimum number of replicas, pods won't scale down to 0 in case of no traffic",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
						"maxReplicas": {
							SchemaProps: spec.SchemaProps{
								Description: "This is the up bound for autoscaler to scale to",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
					},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.CustomSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.FeastTransformerSpec"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.XGBoostSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "XGBoostSpec defines arguments for configuring XGBoost model serving.",
					Properties: map[string]spec.Schema{
						"storageUri": {
							SchemaProps: spec.SchemaProps{
								Description: "The location of the trained model",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"runtimeVersion": {
							SchemaProps: spec.SchemaProps{
								Description: "Defaults to latest XGBoost Version.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"resources": {
							SchemaProps: spec.SchemaProps{
								Description: "Defaults to requests and limits of 1CPU, 2Gb MEM.",
								Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
							},
						},
					},
					Required: []string{"storageUri"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.ResourceRequirements"},
		},
		"knative.dev/pkg/apis.Condition": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "Conditions defines a readiness condition for a Knative resource. See: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#typical-status-properties",
					Properties: map[string]spec.Schema{
						"type": {
							SchemaProps: spec.SchemaProps{
								Description: "Type of condition.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"status": {
							SchemaProps: spec.SchemaProps{
								Description: "Status of the condition, one of True, False, Unknown.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"severity": {
							SchemaProps: spec.SchemaProps{
								Description: "Severity with which to treat failures of this type of condition. When this is not specified, it defaults to Error.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"lastTransitionTime": {
							SchemaProps: spec.SchemaProps{
								Description: "LastTransitionTime is the last time the condition transitioned from one status to another. We use VolatileTime in place of metav1.Time to exclude this from creating equality.Semantic differences (all other things held constant).",
								Ref:         ref("knative.dev/pkg/apis.VolatileTime"),
							},
						},
						"reason": {
							SchemaProps: spec.SchemaProps{
								Description: "The reason for the condition's last transition.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"message": {
							SchemaProps: spec.SchemaProps{
								Description: "A human readable message indicating details about the transition.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
					},
					Required: []string{"type", "status"},
				},
			},
			Dependencies: []string{
				"knative.dev/pkg/apis.VolatileTime"},
		},
		"knative.dev/pkg/apis.VolatileTime": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "VolatileTime wraps metav1.Time",
					Properties: map[string]spec.Schema{
						"Inner": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
							},
						},
					},
					Required: []string{"Inner"},
				},
			},
			Dependencies: []string{
				"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
		},
	}
}
