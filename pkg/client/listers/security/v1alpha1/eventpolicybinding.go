/*
Copyright 2020 Google LLC

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/google/knative-gcp/pkg/apis/security/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EventPolicyBindingLister helps list EventPolicyBindings.
type EventPolicyBindingLister interface {
	// List lists all EventPolicyBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.EventPolicyBinding, err error)
	// EventPolicyBindings returns an object that can list and get EventPolicyBindings.
	EventPolicyBindings(namespace string) EventPolicyBindingNamespaceLister
	EventPolicyBindingListerExpansion
}

// eventPolicyBindingLister implements the EventPolicyBindingLister interface.
type eventPolicyBindingLister struct {
	indexer cache.Indexer
}

// NewEventPolicyBindingLister returns a new EventPolicyBindingLister.
func NewEventPolicyBindingLister(indexer cache.Indexer) EventPolicyBindingLister {
	return &eventPolicyBindingLister{indexer: indexer}
}

// List lists all EventPolicyBindings in the indexer.
func (s *eventPolicyBindingLister) List(selector labels.Selector) (ret []*v1alpha1.EventPolicyBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventPolicyBinding))
	})
	return ret, err
}

// EventPolicyBindings returns an object that can list and get EventPolicyBindings.
func (s *eventPolicyBindingLister) EventPolicyBindings(namespace string) EventPolicyBindingNamespaceLister {
	return eventPolicyBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EventPolicyBindingNamespaceLister helps list and get EventPolicyBindings.
type EventPolicyBindingNamespaceLister interface {
	// List lists all EventPolicyBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.EventPolicyBinding, err error)
	// Get retrieves the EventPolicyBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.EventPolicyBinding, error)
	EventPolicyBindingNamespaceListerExpansion
}

// eventPolicyBindingNamespaceLister implements the EventPolicyBindingNamespaceLister
// interface.
type eventPolicyBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EventPolicyBindings in the indexer for a given namespace.
func (s eventPolicyBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EventPolicyBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventPolicyBinding))
	})
	return ret, err
}

// Get retrieves the EventPolicyBinding from the indexer for a given namespace and name.
func (s eventPolicyBindingNamespaceLister) Get(name string) (*v1alpha1.EventPolicyBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("eventpolicybinding"), name)
	}
	return obj.(*v1alpha1.EventPolicyBinding), nil
}
