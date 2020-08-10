// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	duckv1beta1 "knative.dev/pkg/apis/duck/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ABTestSpec) DeepCopyInto(out *ABTestSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ABTestSpec.
func (in *ABTestSpec) DeepCopy() *ABTestSpec {
	if in == nil {
		return nil
	}
	out := new(ABTestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlibiExplainerSpec) DeepCopyInto(out *AlibiExplainerSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlibiExplainerSpec.
func (in *AlibiExplainerSpec) DeepCopy() *AlibiExplainerSpec {
	if in == nil {
		return nil
	}
	out := new(AlibiExplainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Batcher) DeepCopyInto(out *Batcher) {
	*out = *in
	if in.MaxBatchSize != nil {
		in, out := &in.MaxBatchSize, &out.MaxBatchSize
		*out = new(int)
		**out = **in
	}
	if in.MaxLatency != nil {
		in, out := &in.MaxLatency, &out.MaxLatency
		*out = new(int)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Batcher.
func (in *Batcher) DeepCopy() *Batcher {
	if in == nil {
		return nil
	}
	out := new(Batcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentExtensionSpec) DeepCopyInto(out *ComponentExtensionSpec) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int)
		**out = **in
	}
	if in.LoggerSpec != nil {
		in, out := &in.LoggerSpec, &out.LoggerSpec
		*out = new(LoggerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Batcher != nil {
		in, out := &in.Batcher, &out.Batcher
		*out = new(Batcher)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentExtensionSpec.
func (in *ComponentExtensionSpec) DeepCopy() *ComponentExtensionSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentExtensionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentStatusSpec) DeepCopyInto(out *ComponentStatusSpec) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(duckv1.Addressable)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentStatusSpec.
func (in *ComponentStatusSpec) DeepCopy() *ComponentStatusSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentStatusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomPredictor) DeepCopyInto(out *CustomPredictor) {
	*out = *in
	in.PodTemplateSpec.DeepCopyInto(&out.PodTemplateSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomPredictor.
func (in *CustomPredictor) DeepCopy() *CustomPredictor {
	if in == nil {
		return nil
	}
	out := new(CustomPredictor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnsembleSpec) DeepCopyInto(out *EnsembleSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnsembleSpec.
func (in *EnsembleSpec) DeepCopy() *EnsembleSpec {
	if in == nil {
		return nil
	}
	out := new(EnsembleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExplainerSpec) DeepCopyInto(out *ExplainerSpec) {
	*out = *in
	if in.Alibi != nil {
		in, out := &in.Alibi, &out.Alibi
		*out = new(AlibiExplainerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PodTemplateSpec != nil {
		in, out := &in.PodTemplateSpec, &out.PodTemplateSpec
		*out = new(v1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ComponentExtensionSpec != nil {
		in, out := &in.ComponentExtensionSpec, &out.ComponentExtensionSpec
		*out = new(ComponentExtensionSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExplainerSpec.
func (in *ExplainerSpec) DeepCopy() *ExplainerSpec {
	if in == nil {
		return nil
	}
	out := new(ExplainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceRouter) DeepCopyInto(out *InferenceRouter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceRouter.
func (in *InferenceRouter) DeepCopy() *InferenceRouter {
	if in == nil {
		return nil
	}
	out := new(InferenceRouter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InferenceRouter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceRouterList) DeepCopyInto(out *InferenceRouterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InferenceRouter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceRouterList.
func (in *InferenceRouterList) DeepCopy() *InferenceRouterList {
	if in == nil {
		return nil
	}
	out := new(InferenceRouterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InferenceRouterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceRouterSpec) DeepCopyInto(out *InferenceRouterSpec) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]RouteSpec, len(*in))
		copy(*out, *in)
	}
	if in.Splitter != nil {
		in, out := &in.Splitter, &out.Splitter
		*out = new(SplitterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ABTest != nil {
		in, out := &in.ABTest, &out.ABTest
		*out = new(ABTestSpec)
		**out = **in
	}
	if in.MultiArmBandit != nil {
		in, out := &in.MultiArmBandit, &out.MultiArmBandit
		*out = new(MultiArmBanditSpec)
		**out = **in
	}
	if in.Ensemble != nil {
		in, out := &in.Ensemble, &out.Ensemble
		*out = new(EnsembleSpec)
		**out = **in
	}
	if in.Pipeline != nil {
		in, out := &in.Pipeline, &out.Pipeline
		*out = new(PipelineSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceRouterSpec.
func (in *InferenceRouterSpec) DeepCopy() *InferenceRouterSpec {
	if in == nil {
		return nil
	}
	out := new(InferenceRouterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceRouterStatus) DeepCopyInto(out *InferenceRouterStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(duckv1beta1.Addressable)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceRouterStatus.
func (in *InferenceRouterStatus) DeepCopy() *InferenceRouterStatus {
	if in == nil {
		return nil
	}
	out := new(InferenceRouterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceService) DeepCopyInto(out *InferenceService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceService.
func (in *InferenceService) DeepCopy() *InferenceService {
	if in == nil {
		return nil
	}
	out := new(InferenceService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InferenceService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceServiceList) DeepCopyInto(out *InferenceServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InferenceService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceServiceList.
func (in *InferenceServiceList) DeepCopy() *InferenceServiceList {
	if in == nil {
		return nil
	}
	out := new(InferenceServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InferenceServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceServiceSpec) DeepCopyInto(out *InferenceServiceSpec) {
	*out = *in
	in.Predictor.DeepCopyInto(&out.Predictor)
	if in.Explainer != nil {
		in, out := &in.Explainer, &out.Explainer
		*out = new(ExplainerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Transformer != nil {
		in, out := &in.Transformer, &out.Transformer
		*out = new(TransformerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceServiceSpec.
func (in *InferenceServiceSpec) DeepCopy() *InferenceServiceSpec {
	if in == nil {
		return nil
	}
	out := new(InferenceServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceServiceStatus) DeepCopyInto(out *InferenceServiceStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(duckv1.Addressable)
		(*in).DeepCopyInto(*out)
	}
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make(map[ComponentType]ComponentStatusSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceServiceStatus.
func (in *InferenceServiceStatus) DeepCopy() *InferenceServiceStatus {
	if in == nil {
		return nil
	}
	out := new(InferenceServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggerSpec) DeepCopyInto(out *LoggerSpec) {
	*out = *in
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggerSpec.
func (in *LoggerSpec) DeepCopy() *LoggerSpec {
	if in == nil {
		return nil
	}
	out := new(LoggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelSpec) DeepCopyInto(out *ModelSpec) {
	*out = *in
	out.Memory = in.Memory.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelSpec.
func (in *ModelSpec) DeepCopy() *ModelSpec {
	if in == nil {
		return nil
	}
	out := new(ModelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiArmBanditSpec) DeepCopyInto(out *MultiArmBanditSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiArmBanditSpec.
func (in *MultiArmBanditSpec) DeepCopy() *MultiArmBanditSpec {
	if in == nil {
		return nil
	}
	out := new(MultiArmBanditSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ONNXRuntimeSpec) DeepCopyInto(out *ONNXRuntimeSpec) {
	*out = *in
	in.PredictorExtensionSpec.DeepCopyInto(&out.PredictorExtensionSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ONNXRuntimeSpec.
func (in *ONNXRuntimeSpec) DeepCopy() *ONNXRuntimeSpec {
	if in == nil {
		return nil
	}
	out := new(ONNXRuntimeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpec) DeepCopyInto(out *PipelineSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpec.
func (in *PipelineSpec) DeepCopy() *PipelineSpec {
	if in == nil {
		return nil
	}
	out := new(PipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorExtensionSpec) DeepCopyInto(out *PredictorExtensionSpec) {
	*out = *in
	if in.StorageURI != nil {
		in, out := &in.StorageURI, &out.StorageURI
		*out = new(string)
		**out = **in
	}
	in.Container.DeepCopyInto(&out.Container)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorExtensionSpec.
func (in *PredictorExtensionSpec) DeepCopy() *PredictorExtensionSpec {
	if in == nil {
		return nil
	}
	out := new(PredictorExtensionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorSpec) DeepCopyInto(out *PredictorSpec) {
	*out = *in
	if in.SKLearn != nil {
		in, out := &in.SKLearn, &out.SKLearn
		*out = new(SKLearnSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.XGBoost != nil {
		in, out := &in.XGBoost, &out.XGBoost
		*out = new(XGBoostSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Tensorflow != nil {
		in, out := &in.Tensorflow, &out.Tensorflow
		*out = new(TensorflowSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PyTorch != nil {
		in, out := &in.PyTorch, &out.PyTorch
		*out = new(TorchServeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Triton != nil {
		in, out := &in.Triton, &out.Triton
		*out = new(TritonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ONNXRuntime != nil {
		in, out := &in.ONNXRuntime, &out.ONNXRuntime
		*out = new(ONNXRuntimeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomPredictor != nil {
		in, out := &in.CustomPredictor, &out.CustomPredictor
		*out = new(CustomPredictor)
		(*in).DeepCopyInto(*out)
	}
	in.ComponentExtensionSpec.DeepCopyInto(&out.ComponentExtensionSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorSpec.
func (in *PredictorSpec) DeepCopy() *PredictorSpec {
	if in == nil {
		return nil
	}
	out := new(PredictorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpec) DeepCopyInto(out *RouteSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpec.
func (in *RouteSpec) DeepCopy() *RouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SKLearnSpec) DeepCopyInto(out *SKLearnSpec) {
	*out = *in
	in.PredictorExtensionSpec.DeepCopyInto(&out.PredictorExtensionSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SKLearnSpec.
func (in *SKLearnSpec) DeepCopy() *SKLearnSpec {
	if in == nil {
		return nil
	}
	out := new(SKLearnSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SplitterSpec) DeepCopyInto(out *SplitterSpec) {
	*out = *in
	if in.Weights != nil {
		in, out := &in.Weights, &out.Weights
		*out = make([]*WeightsSpec, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(WeightsSpec)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SplitterSpec.
func (in *SplitterSpec) DeepCopy() *SplitterSpec {
	if in == nil {
		return nil
	}
	out := new(SplitterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TensorflowSpec) DeepCopyInto(out *TensorflowSpec) {
	*out = *in
	in.PredictorExtensionSpec.DeepCopyInto(&out.PredictorExtensionSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TensorflowSpec.
func (in *TensorflowSpec) DeepCopy() *TensorflowSpec {
	if in == nil {
		return nil
	}
	out := new(TensorflowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TorchServeSpec) DeepCopyInto(out *TorchServeSpec) {
	*out = *in
	in.PredictorExtensionSpec.DeepCopyInto(&out.PredictorExtensionSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TorchServeSpec.
func (in *TorchServeSpec) DeepCopy() *TorchServeSpec {
	if in == nil {
		return nil
	}
	out := new(TorchServeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainedModel) DeepCopyInto(out *TrainedModel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainedModel.
func (in *TrainedModel) DeepCopy() *TrainedModel {
	if in == nil {
		return nil
	}
	out := new(TrainedModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrainedModel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainedModelList) DeepCopyInto(out *TrainedModelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrainedModel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainedModelList.
func (in *TrainedModelList) DeepCopy() *TrainedModelList {
	if in == nil {
		return nil
	}
	out := new(TrainedModelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrainedModelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainedModelSpec) DeepCopyInto(out *TrainedModelSpec) {
	*out = *in
	in.Inference.DeepCopyInto(&out.Inference)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainedModelSpec.
func (in *TrainedModelSpec) DeepCopy() *TrainedModelSpec {
	if in == nil {
		return nil
	}
	out := new(TrainedModelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainedModelStatus) DeepCopyInto(out *TrainedModelStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(duckv1beta1.Addressable)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainedModelStatus.
func (in *TrainedModelStatus) DeepCopy() *TrainedModelStatus {
	if in == nil {
		return nil
	}
	out := new(TrainedModelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransformerSpec) DeepCopyInto(out *TransformerSpec) {
	*out = *in
	if in.PodTemplateSpec != nil {
		in, out := &in.PodTemplateSpec, &out.PodTemplateSpec
		*out = new(v1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ComponentExtensionSpec != nil {
		in, out := &in.ComponentExtensionSpec, &out.ComponentExtensionSpec
		*out = new(ComponentExtensionSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransformerSpec.
func (in *TransformerSpec) DeepCopy() *TransformerSpec {
	if in == nil {
		return nil
	}
	out := new(TransformerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TritonSpec) DeepCopyInto(out *TritonSpec) {
	*out = *in
	in.PredictorExtensionSpec.DeepCopyInto(&out.PredictorExtensionSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TritonSpec.
func (in *TritonSpec) DeepCopy() *TritonSpec {
	if in == nil {
		return nil
	}
	out := new(TritonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeightsSpec) DeepCopyInto(out *WeightsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeightsSpec.
func (in *WeightsSpec) DeepCopy() *WeightsSpec {
	if in == nil {
		return nil
	}
	out := new(WeightsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XGBoostSpec) DeepCopyInto(out *XGBoostSpec) {
	*out = *in
	in.PredictorExtensionSpec.DeepCopyInto(&out.PredictorExtensionSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XGBoostSpec.
func (in *XGBoostSpec) DeepCopy() *XGBoostSpec {
	if in == nil {
		return nil
	}
	out := new(XGBoostSpec)
	in.DeepCopyInto(out)
	return out
}
