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

package v1alpha2

import (
	"github.com/kubeflow/kfserving/pkg/constants"
	"k8s.io/apimachinery/pkg/runtime"
	"knative.dev/pkg/apis/duck/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIXExplainerSpec) DeepCopyInto(out *AIXExplainerSpec) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIXExplainerSpec.
func (in *AIXExplainerSpec) DeepCopy() *AIXExplainerSpec {
	if in == nil {
		return nil
	}
	out := new(AIXExplainerSpec)
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
func (in *CustomSpec) DeepCopyInto(out *CustomSpec) {
	*out = *in
	in.Container.DeepCopyInto(&out.Container)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomSpec.
func (in *CustomSpec) DeepCopy() *CustomSpec {
	if in == nil {
		return nil
	}
	out := new(CustomSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentSpec) DeepCopyInto(out *DeploymentSpec) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int)
		**out = **in
	}
	if in.Logger != nil {
		in, out := &in.Logger, &out.Logger
		*out = new(Logger)
		(*in).DeepCopyInto(*out)
	}
	if in.Batcher != nil {
		in, out := &in.Batcher, &out.Batcher
		*out = new(Batcher)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSpec.
func (in *DeploymentSpec) DeepCopy() *DeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSpec) DeepCopyInto(out *EndpointSpec) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSpec.
func (in *EndpointSpec) DeepCopy() *EndpointSpec {
	if in == nil {
		return nil
	}
	out := new(EndpointSpec)
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
	if in.AIX != nil {
		in, out := &in.AIX, &out.AIX
		*out = new(AIXExplainerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = new(CustomSpec)
		(*in).DeepCopyInto(*out)
	}
	in.DeploymentSpec.DeepCopyInto(&out.DeploymentSpec)
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
	in.Default.DeepCopyInto(&out.Default)
	if in.Canary != nil {
		in, out := &in.Canary, &out.Canary
		*out = new(EndpointSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CanaryTrafficPercent != nil {
		in, out := &in.CanaryTrafficPercent, &out.CanaryTrafficPercent
		*out = new(int)
		**out = **in
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
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(map[constants.InferenceServiceComponent]StatusConfigurationSpec)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[constants.InferenceServiceComponent]StatusConfigurationSpec, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Canary != nil {
		in, out := &in.Canary, &out.Canary
		*out = new(map[constants.InferenceServiceComponent]StatusConfigurationSpec)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[constants.InferenceServiceComponent]StatusConfigurationSpec, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(v1beta1.Addressable)
		(*in).DeepCopyInto(*out)
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
func (in *LightGBMSpec) DeepCopyInto(out *LightGBMSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LightGBMSpec.
func (in *LightGBMSpec) DeepCopy() *LightGBMSpec {
	if in == nil {
		return nil
	}
	out := new(LightGBMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Logger) DeepCopyInto(out *Logger) {
	*out = *in
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Logger.
func (in *Logger) DeepCopy() *Logger {
	if in == nil {
		return nil
	}
	out := new(Logger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ONNXSpec) DeepCopyInto(out *ONNXSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ONNXSpec.
func (in *ONNXSpec) DeepCopy() *ONNXSpec {
	if in == nil {
		return nil
	}
	out := new(ONNXSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PMMLSpec) DeepCopyInto(out *PMMLSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PMMLSpec.
func (in *PMMLSpec) DeepCopy() *PMMLSpec {
	if in == nil {
		return nil
	}
	out := new(PMMLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorSpec) DeepCopyInto(out *PredictorSpec) {
	*out = *in
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = new(CustomSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Tensorflow != nil {
		in, out := &in.Tensorflow, &out.Tensorflow
		*out = new(TensorflowSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Triton != nil {
		in, out := &in.Triton, &out.Triton
		*out = new(TritonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.XGBoost != nil {
		in, out := &in.XGBoost, &out.XGBoost
		*out = new(XGBoostSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.LightGBM != nil {
		in, out := &in.LightGBM, &out.LightGBM
		*out = new(LightGBMSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.SKLearn != nil {
		in, out := &in.SKLearn, &out.SKLearn
		*out = new(SKLearnSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ONNX != nil {
		in, out := &in.ONNX, &out.ONNX
		*out = new(ONNXSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PyTorch != nil {
		in, out := &in.PyTorch, &out.PyTorch
		*out = new(PyTorchSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PMML != nil {
		in, out := &in.PMML, &out.PMML
		*out = new(PMMLSpec)
		(*in).DeepCopyInto(*out)
	}
	in.DeploymentSpec.DeepCopyInto(&out.DeploymentSpec)
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
func (in *PyTorchSpec) DeepCopyInto(out *PyTorchSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PyTorchSpec.
func (in *PyTorchSpec) DeepCopy() *PyTorchSpec {
	if in == nil {
		return nil
	}
	out := new(PyTorchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SKLearnSpec) DeepCopyInto(out *SKLearnSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
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
func (in *StatusConfigurationSpec) DeepCopyInto(out *StatusConfigurationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusConfigurationSpec.
func (in *StatusConfigurationSpec) DeepCopy() *StatusConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(StatusConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TensorflowSpec) DeepCopyInto(out *TensorflowSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
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
func (in *TransformerSpec) DeepCopyInto(out *TransformerSpec) {
	*out = *in
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = new(CustomSpec)
		(*in).DeepCopyInto(*out)
	}
	in.DeploymentSpec.DeepCopyInto(&out.DeploymentSpec)
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
	in.Resources.DeepCopyInto(&out.Resources)
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
func (in *VirtualServiceStatus) DeepCopyInto(out *VirtualServiceStatus) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(v1beta1.Addressable)
		(*in).DeepCopyInto(*out)
	}
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServiceStatus.
func (in *VirtualServiceStatus) DeepCopy() *VirtualServiceStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XGBoostSpec) DeepCopyInto(out *XGBoostSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
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
