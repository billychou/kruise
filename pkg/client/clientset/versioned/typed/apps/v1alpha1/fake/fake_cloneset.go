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
	v1alpha1 "github.com/openkruise/kruise/pkg/apis/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloneSets implements CloneSetInterface
type FakeCloneSets struct {
	Fake *FakeAppsV1alpha1
	ns   string
}

var clonesetsResource = schema.GroupVersionResource{Group: "apps.kruise.io", Version: "v1alpha1", Resource: "clonesets"}

var clonesetsKind = schema.GroupVersionKind{Group: "apps.kruise.io", Version: "v1alpha1", Kind: "CloneSet"}

// Get takes name of the cloneSet, and returns the corresponding cloneSet object, and an error if there is any.
func (c *FakeCloneSets) Get(name string, options v1.GetOptions) (result *v1alpha1.CloneSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clonesetsResource, c.ns, name), &v1alpha1.CloneSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloneSet), err
}

// List takes label and field selectors, and returns the list of CloneSets that match those selectors.
func (c *FakeCloneSets) List(opts v1.ListOptions) (result *v1alpha1.CloneSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clonesetsResource, clonesetsKind, c.ns, opts), &v1alpha1.CloneSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloneSetList{ListMeta: obj.(*v1alpha1.CloneSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloneSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloneSets.
func (c *FakeCloneSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clonesetsResource, c.ns, opts))

}

// Create takes the representation of a cloneSet and creates it.  Returns the server's representation of the cloneSet, and an error, if there is any.
func (c *FakeCloneSets) Create(cloneSet *v1alpha1.CloneSet) (result *v1alpha1.CloneSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clonesetsResource, c.ns, cloneSet), &v1alpha1.CloneSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloneSet), err
}

// Update takes the representation of a cloneSet and updates it. Returns the server's representation of the cloneSet, and an error, if there is any.
func (c *FakeCloneSets) Update(cloneSet *v1alpha1.CloneSet) (result *v1alpha1.CloneSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clonesetsResource, c.ns, cloneSet), &v1alpha1.CloneSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloneSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloneSets) UpdateStatus(cloneSet *v1alpha1.CloneSet) (*v1alpha1.CloneSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clonesetsResource, "status", c.ns, cloneSet), &v1alpha1.CloneSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloneSet), err
}

// Delete takes name of the cloneSet and deletes it. Returns an error if one occurs.
func (c *FakeCloneSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clonesetsResource, c.ns, name), &v1alpha1.CloneSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloneSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clonesetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloneSetList{})
	return err
}

// Patch applies the patch and returns the patched cloneSet.
func (c *FakeCloneSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloneSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clonesetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloneSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloneSet), err
}
