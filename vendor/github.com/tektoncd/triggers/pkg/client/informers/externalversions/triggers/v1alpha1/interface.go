/*
Copyright 2019 The Tekton Authors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/tektoncd/triggers/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterInterceptors returns a ClusterInterceptorInformer.
	ClusterInterceptors() ClusterInterceptorInformer
	// ClusterTriggerBindings returns a ClusterTriggerBindingInformer.
	ClusterTriggerBindings() ClusterTriggerBindingInformer
	// EventListeners returns a EventListenerInformer.
	EventListeners() EventListenerInformer
	// Interceptors returns a InterceptorInformer.
	Interceptors() InterceptorInformer
	// Triggers returns a TriggerInformer.
	Triggers() TriggerInformer
	// TriggerBindings returns a TriggerBindingInformer.
	TriggerBindings() TriggerBindingInformer
	// TriggerTemplates returns a TriggerTemplateInformer.
	TriggerTemplates() TriggerTemplateInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterInterceptors returns a ClusterInterceptorInformer.
func (v *version) ClusterInterceptors() ClusterInterceptorInformer {
	return &clusterInterceptorInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterTriggerBindings returns a ClusterTriggerBindingInformer.
func (v *version) ClusterTriggerBindings() ClusterTriggerBindingInformer {
	return &clusterTriggerBindingInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// EventListeners returns a EventListenerInformer.
func (v *version) EventListeners() EventListenerInformer {
	return &eventListenerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Interceptors returns a InterceptorInformer.
func (v *version) Interceptors() InterceptorInformer {
	return &interceptorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Triggers returns a TriggerInformer.
func (v *version) Triggers() TriggerInformer {
	return &triggerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TriggerBindings returns a TriggerBindingInformer.
func (v *version) TriggerBindings() TriggerBindingInformer {
	return &triggerBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TriggerTemplates returns a TriggerTemplateInformer.
func (v *version) TriggerTemplates() TriggerTemplateInformer {
	return &triggerTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
