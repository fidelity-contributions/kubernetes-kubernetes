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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	apisapiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	clientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	internalinterfaces "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions/internalinterfaces"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/client/listers/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CustomResourceDefinitionInformer provides access to a shared informer and lister for
// CustomResourceDefinitions.
type CustomResourceDefinitionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() apiextensionsv1.CustomResourceDefinitionLister
}

type customResourceDefinitionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCustomResourceDefinitionInformer constructs a new informer for CustomResourceDefinition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCustomResourceDefinitionInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCustomResourceDefinitionInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCustomResourceDefinitionInformer constructs a new informer for CustomResourceDefinition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCustomResourceDefinitionInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApiextensionsV1().CustomResourceDefinitions().List(context.Background(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApiextensionsV1().CustomResourceDefinitions().Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApiextensionsV1().CustomResourceDefinitions().List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApiextensionsV1().CustomResourceDefinitions().Watch(ctx, options)
			},
		},
		&apisapiextensionsv1.CustomResourceDefinition{},
		resyncPeriod,
		indexers,
	)
}

func (f *customResourceDefinitionInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCustomResourceDefinitionInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *customResourceDefinitionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisapiextensionsv1.CustomResourceDefinition{}, f.defaultInformer)
}

func (f *customResourceDefinitionInformer) Lister() apiextensionsv1.CustomResourceDefinitionLister {
	return apiextensionsv1.NewCustomResourceDefinitionLister(f.Informer().GetIndexer())
}
