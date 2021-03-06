/*
Copyright 2017 The Kubernetes Authors.

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

package internalversion

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	testgroup "k8s.io/kubernetes/cmd/libs/go2idl/client-gen/test_apis/testgroup"
	api "k8s.io/kubernetes/pkg/api"
)

// TestTypesGetter has a method to return a TestTypeInterface.
// A group's client should implement this interface.
type TestTypesGetter interface {
	TestTypes(namespace string) TestTypeInterface
}

// TestTypeInterface has methods to work with TestType resources.
type TestTypeInterface interface {
	Create(*testgroup.TestType) (*testgroup.TestType, error)
	Update(*testgroup.TestType) (*testgroup.TestType, error)
	UpdateStatus(*testgroup.TestType) (*testgroup.TestType, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*testgroup.TestType, error)
	List(opts v1.ListOptions) (*testgroup.TestTypeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *testgroup.TestType, err error)
	TestTypeExpansion
}

// testTypes implements TestTypeInterface
type testTypes struct {
	client rest.Interface
	ns     string
}

// newTestTypes returns a TestTypes
func newTestTypes(c *TestgroupClient, namespace string) *testTypes {
	return &testTypes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a testType and creates it.  Returns the server's representation of the testType, and an error, if there is any.
func (c *testTypes) Create(testType *testgroup.TestType) (result *testgroup.TestType, err error) {
	result = &testgroup.TestType{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("testtypes").
		Body(testType).
		Do().
		Into(result)
	return
}

// Update takes the representation of a testType and updates it. Returns the server's representation of the testType, and an error, if there is any.
func (c *testTypes) Update(testType *testgroup.TestType) (result *testgroup.TestType, err error) {
	result = &testgroup.TestType{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("testtypes").
		Name(testType.Name).
		Body(testType).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclientstatus=false comment above the type to avoid generating UpdateStatus().

func (c *testTypes) UpdateStatus(testType *testgroup.TestType) (result *testgroup.TestType, err error) {
	result = &testgroup.TestType{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("testtypes").
		Name(testType.Name).
		SubResource("status").
		Body(testType).
		Do().
		Into(result)
	return
}

// Delete takes name of the testType and deletes it. Returns an error if one occurs.
func (c *testTypes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("testtypes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *testTypes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("testtypes").
		VersionedParams(&listOptions, api.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the testType, and returns the corresponding testType object, and an error if there is any.
func (c *testTypes) Get(name string, options v1.GetOptions) (result *testgroup.TestType, err error) {
	result = &testgroup.TestType{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("testtypes").
		Name(name).
		VersionedParams(&options, api.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TestTypes that match those selectors.
func (c *testTypes) List(opts v1.ListOptions) (result *testgroup.TestTypeList, err error) {
	result = &testgroup.TestTypeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("testtypes").
		VersionedParams(&opts, api.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested testTypes.
func (c *testTypes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("testtypes").
		VersionedParams(&opts, api.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched testType.
func (c *testTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *testgroup.TestType, err error) {
	result = &testgroup.TestType{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("testtypes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
