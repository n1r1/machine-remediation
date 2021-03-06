/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2019 Red Hat, Inc.
 *
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubevirt.io/machine-remediation/pkg/apis/machineremediation/v1alpha1"
)

// FakeMachineRemediations implements MachineRemediationInterface
type FakeMachineRemediations struct {
	Fake *FakeMachineremediationV1alpha1
	ns   string
}

var machineremediationsResource = schema.GroupVersionResource{Group: "machineremediation.kubevirt.io", Version: "v1alpha1", Resource: "machineremediations"}

var machineremediationsKind = schema.GroupVersionKind{Group: "machineremediation.kubevirt.io", Version: "v1alpha1", Kind: "MachineRemediation"}

// Get takes name of the machineRemediation, and returns the corresponding machineRemediation object, and an error if there is any.
func (c *FakeMachineRemediations) Get(name string, options v1.GetOptions) (result *v1alpha1.MachineRemediation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(machineremediationsResource, c.ns, name), &v1alpha1.MachineRemediation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineRemediation), err
}

// List takes label and field selectors, and returns the list of MachineRemediations that match those selectors.
func (c *FakeMachineRemediations) List(opts v1.ListOptions) (result *v1alpha1.MachineRemediationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(machineremediationsResource, machineremediationsKind, c.ns, opts), &v1alpha1.MachineRemediationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MachineRemediationList{ListMeta: obj.(*v1alpha1.MachineRemediationList).ListMeta}
	for _, item := range obj.(*v1alpha1.MachineRemediationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineRemediations.
func (c *FakeMachineRemediations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(machineremediationsResource, c.ns, opts))

}

// Create takes the representation of a machineRemediation and creates it.  Returns the server's representation of the machineRemediation, and an error, if there is any.
func (c *FakeMachineRemediations) Create(machineRemediation *v1alpha1.MachineRemediation) (result *v1alpha1.MachineRemediation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(machineremediationsResource, c.ns, machineRemediation), &v1alpha1.MachineRemediation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineRemediation), err
}

// Update takes the representation of a machineRemediation and updates it. Returns the server's representation of the machineRemediation, and an error, if there is any.
func (c *FakeMachineRemediations) Update(machineRemediation *v1alpha1.MachineRemediation) (result *v1alpha1.MachineRemediation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(machineremediationsResource, c.ns, machineRemediation), &v1alpha1.MachineRemediation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineRemediation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMachineRemediations) UpdateStatus(machineRemediation *v1alpha1.MachineRemediation) (*v1alpha1.MachineRemediation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(machineremediationsResource, "status", c.ns, machineRemediation), &v1alpha1.MachineRemediation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineRemediation), err
}

// Delete takes name of the machineRemediation and deletes it. Returns an error if one occurs.
func (c *FakeMachineRemediations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(machineremediationsResource, c.ns, name), &v1alpha1.MachineRemediation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineRemediations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(machineremediationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MachineRemediationList{})
	return err
}

// Patch applies the patch and returns the patched machineRemediation.
func (c *FakeMachineRemediations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MachineRemediation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(machineremediationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MachineRemediation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineRemediation), err
}
