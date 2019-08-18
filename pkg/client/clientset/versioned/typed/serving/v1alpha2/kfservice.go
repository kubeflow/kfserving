/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2"
	scheme "github.com/kubeflow/kfserving/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KFServicesGetter has a method to return a KFServiceInterface.
// A group's client should implement this interface.
type KFServicesGetter interface {
	KFServices(namespace string) KFServiceInterface
}

// KFServiceInterface has methods to work with KFService resources.
type KFServiceInterface interface {
	Create(*v1alpha2.KFService) (*v1alpha2.KFService, error)
	Update(*v1alpha2.KFService) (*v1alpha2.KFService, error)
	UpdateStatus(*v1alpha2.KFService) (*v1alpha2.KFService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.KFService, error)
	List(opts v1.ListOptions) (*v1alpha2.KFServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.KFService, err error)
	KFServiceExpansion
}

// kFServices implements KFServiceInterface
type kFServices struct {
	client rest.Interface
	ns     string
}

// newKFServices returns a KFServices
func newKFServices(c *ServingV1alpha2Client, namespace string) *kFServices {
	return &kFServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kFService, and returns the corresponding kFService object, and an error if there is any.
func (c *kFServices) Get(name string, options v1.GetOptions) (result *v1alpha2.KFService, err error) {
	result = &v1alpha2.KFService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kfservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KFServices that match those selectors.
func (c *kFServices) List(opts v1.ListOptions) (result *v1alpha2.KFServiceList, err error) {
	result = &v1alpha2.KFServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kfservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kFServices.
func (c *kFServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kfservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a kFService and creates it.  Returns the server's representation of the kFService, and an error, if there is any.
func (c *kFServices) Create(kFService *v1alpha2.KFService) (result *v1alpha2.KFService, err error) {
	result = &v1alpha2.KFService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kfservices").
		Body(kFService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kFService and updates it. Returns the server's representation of the kFService, and an error, if there is any.
func (c *kFServices) Update(kFService *v1alpha2.KFService) (result *v1alpha2.KFService, err error) {
	result = &v1alpha2.KFService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kfservices").
		Name(kFService.Name).
		Body(kFService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kFServices) UpdateStatus(kFService *v1alpha2.KFService) (result *v1alpha2.KFService, err error) {
	result = &v1alpha2.KFService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kfservices").
		Name(kFService.Name).
		SubResource("status").
		Body(kFService).
		Do().
		Into(result)
	return
}

// Delete takes name of the kFService and deletes it. Returns an error if one occurs.
func (c *kFServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kfservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kFServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kfservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kFService.
func (c *kFServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.KFService, err error) {
	result = &v1alpha2.KFService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kfservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
