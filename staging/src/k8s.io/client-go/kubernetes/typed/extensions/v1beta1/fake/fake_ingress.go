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

	v1beta1 "k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	extensionsv1beta1 "k8s.io/client-go/applyconfigurations/extensions/v1beta1"
	testing "k8s.io/client-go/testing"
)

// FakeIngresses implements IngressInterface
type FakeIngresses struct {
	Fake *FakeExtensionsV1beta1
	ns   string
}

var ingressesResource = schema.GroupVersionResource{Group: "extensions", Version: "v1beta1", Resource: "ingresses"}

var ingressesKind = schema.GroupVersionKind{Group: "extensions", Version: "v1beta1", Kind: "Ingress"}

// Get takes name of the ingress, and returns the corresponding ingress object, and an error if there is any.
func (c *FakeIngresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Ingress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ingressesResource, c.ns, name), &v1beta1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ingress), err
}

// List takes label and field selectors, and returns the list of Ingresses that match those selectors.
func (c *FakeIngresses) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IngressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ingressesResource, ingressesKind, c.ns, opts), &v1beta1.IngressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IngressList{ListMeta: obj.(*v1beta1.IngressList).ListMeta}
	for _, item := range obj.(*v1beta1.IngressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ingresses.
func (c *FakeIngresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ingressesResource, c.ns, opts))

}

// Create takes the representation of a ingress and creates it.  Returns the server's representation of the ingress, and an error, if there is any.
func (c *FakeIngresses) Create(ctx context.Context, ingress *v1beta1.Ingress, opts v1.CreateOptions) (result *v1beta1.Ingress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ingressesResource, c.ns, ingress), &v1beta1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ingress), err
}

// Update takes the representation of a ingress and updates it. Returns the server's representation of the ingress, and an error, if there is any.
func (c *FakeIngresses) Update(ctx context.Context, ingress *v1beta1.Ingress, opts v1.UpdateOptions) (result *v1beta1.Ingress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ingressesResource, c.ns, ingress), &v1beta1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ingress), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIngresses) UpdateStatus(ctx context.Context, ingress *v1beta1.Ingress, opts v1.UpdateOptions) (*v1beta1.Ingress, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ingressesResource, "status", c.ns, ingress), &v1beta1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ingress), err
}

// Delete takes name of the ingress and deletes it. Returns an error if one occurs.
func (c *FakeIngresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ingressesResource, c.ns, name, &opts), &v1beta1.Ingress{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIngresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ingressesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.IngressList{})
	return err
}

// Patch applies the patch and returns the patched ingress.
func (c *FakeIngresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Ingress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ingressesResource, c.ns, name, pt, data, subresources...), &v1beta1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ingress), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied ingress.
func (c *FakeIngresses) Apply(ctx context.Context, ingress *extensionsv1beta1.IngressApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Ingress, err error) {
	if ingress == nil {
		return nil, fmt.Errorf("ingress provided to Apply must not be nil")
	}
	data, err := json.Marshal(ingress)
	if err != nil {
		return nil, err
	}
	name := ingress.Name
	if name == nil {
		return nil, fmt.Errorf("ingress.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ingressesResource, c.ns, *name, types.ApplyPatchType, data), &v1beta1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ingress), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeIngresses) ApplyStatus(ctx context.Context, ingress *extensionsv1beta1.IngressApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Ingress, err error) {
	if ingress == nil {
		return nil, fmt.Errorf("ingress provided to Apply must not be nil")
	}
	data, err := json.Marshal(ingress)
	if err != nil {
		return nil, err
	}
	name := ingress.Name
	if name == nil {
		return nil, fmt.Errorf("ingress.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ingressesResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1beta1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ingress), err
}
