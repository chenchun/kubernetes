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

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1beta1 "k8s.io/api/discovery/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	discoveryv1beta1 "k8s.io/client-go/applyconfigurations/discovery/v1beta1"
	testing "k8s.io/client-go/testing"
)

// FakeEndpointSlices implements EndpointSliceInterface
type FakeEndpointSlices struct {
	Fake *FakeDiscoveryV1beta1
	ns   string
}

var endpointslicesResource = schema.GroupVersionResource{Group: "discovery.k8s.io", Version: "v1beta1", Resource: "endpointslices"}

var endpointslicesKind = schema.GroupVersionKind{Group: "discovery.k8s.io", Version: "v1beta1", Kind: "EndpointSlice"}

// Get takes name of the endpointSlice, and returns the corresponding endpointSlice object, and an error if there is any.
func (c *FakeEndpointSlices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.EndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(endpointslicesResource, c.ns, name), &v1beta1.EndpointSlice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EndpointSlice), err
}

// List takes label and field selectors, and returns the list of EndpointSlices that match those selectors.
func (c *FakeEndpointSlices) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.EndpointSliceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(endpointslicesResource, endpointslicesKind, c.ns, opts), &v1beta1.EndpointSliceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.EndpointSliceList{ListMeta: obj.(*v1beta1.EndpointSliceList).ListMeta}
	for _, item := range obj.(*v1beta1.EndpointSliceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested endpointSlices.
func (c *FakeEndpointSlices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(endpointslicesResource, c.ns, opts))

}

// Create takes the representation of a endpointSlice and creates it.  Returns the server's representation of the endpointSlice, and an error, if there is any.
func (c *FakeEndpointSlices) Create(ctx context.Context, endpointSlice *v1beta1.EndpointSlice, opts v1.CreateOptions) (result *v1beta1.EndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(endpointslicesResource, c.ns, endpointSlice), &v1beta1.EndpointSlice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EndpointSlice), err
}

// Update takes the representation of a endpointSlice and updates it. Returns the server's representation of the endpointSlice, and an error, if there is any.
func (c *FakeEndpointSlices) Update(ctx context.Context, endpointSlice *v1beta1.EndpointSlice, opts v1.UpdateOptions) (result *v1beta1.EndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(endpointslicesResource, c.ns, endpointSlice), &v1beta1.EndpointSlice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EndpointSlice), err
}

// Delete takes name of the endpointSlice and deletes it. Returns an error if one occurs.
func (c *FakeEndpointSlices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(endpointslicesResource, c.ns, name, &opts), &v1beta1.EndpointSlice{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEndpointSlices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(endpointslicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.EndpointSliceList{})
	return err
}

// Patch applies the patch and returns the patched endpointSlice.
func (c *FakeEndpointSlices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.EndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(endpointslicesResource, c.ns, name, pt, data, subresources...), &v1beta1.EndpointSlice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EndpointSlice), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied endpointSlice.
func (c *FakeEndpointSlices) Apply(ctx context.Context, endpointSlice *discoveryv1beta1.EndpointSliceApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.EndpointSlice, err error) {
	if endpointSlice == nil {
		return nil, fmt.Errorf("endpointSlice provided to Apply must not be nil")
	}
	data, err := json.Marshal(endpointSlice)
	if err != nil {
		return nil, err
	}
	name := endpointSlice.Name
	if name == nil {
		return nil, fmt.Errorf("endpointSlice.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(endpointslicesResource, c.ns, *name, types.ApplyPatchType, data), &v1beta1.EndpointSlice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EndpointSlice), err
}
