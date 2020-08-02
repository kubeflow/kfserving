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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	"knative.dev/pkg/apis/duck/v1"
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
func (in *EpsilonGreedySpec) DeepCopyInto(out *EpsilonGreedySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EpsilonGreedySpec.
func (in *EpsilonGreedySpec) DeepCopy() *EpsilonGreedySpec {
	if in == nil {
		return nil
	}
	out := new(EpsilonGreedySpec)
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
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
		(*in).DeepCopyInto(*out)
	}
	if in.Ensemble != nil {
		in, out := &in.Ensemble, &out.Ensemble
		*out = new(EnsembleSpec)
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
		*out = new(v1.Addressable)
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
func (in *MultiArmBanditSpec) DeepCopyInto(out *MultiArmBanditSpec) {
	*out = *in
	if in.EpsilonGreedy != nil {
		in, out := &in.EpsilonGreedy, &out.EpsilonGreedy
		*out = new(EpsilonGreedySpec)
		**out = **in
	}
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
func (in *RouteSpec) DeepCopyInto(out *RouteSpec) {
	*out = *in
	in.Destination.DeepCopyInto(&out.Destination)
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]*StringMatch, len(*in))
		for key, val := range *in {
			var outVal *StringMatch
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(StringMatch)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
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
func (in *StringMatch) DeepCopyInto(out *StringMatch) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringMatch.
func (in *StringMatch) DeepCopy() *StringMatch {
	if in == nil {
		return nil
	}
	out := new(StringMatch)
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
