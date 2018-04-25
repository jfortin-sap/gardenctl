// This file was automatically generated by informer-gen

package internalversion

import (
	time "time"

	garden "github.com/gardener/gardener/pkg/apis/garden"
	clientset_internalversion "github.com/gardener/gardener/pkg/client/garden/clientset/internalversion"
	internalinterfaces "github.com/gardener/gardener/pkg/client/garden/informers/internalversion/internalinterfaces"
	internalversion "github.com/gardener/gardener/pkg/client/garden/listers/garden/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CrossSecretBindingInformer provides access to a shared informer and lister for
// CrossSecretBindings.
type CrossSecretBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.CrossSecretBindingLister
}

type crossSecretBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCrossSecretBindingInformer constructs a new informer for CrossSecretBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCrossSecretBindingInformer(client clientset_internalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCrossSecretBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCrossSecretBindingInformer constructs a new informer for CrossSecretBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCrossSecretBindingInformer(client clientset_internalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Garden().CrossSecretBindings(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Garden().CrossSecretBindings(namespace).Watch(options)
			},
		},
		&garden.CrossSecretBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *crossSecretBindingInformer) defaultInformer(client clientset_internalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCrossSecretBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *crossSecretBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&garden.CrossSecretBinding{}, f.defaultInformer)
}

func (f *crossSecretBindingInformer) Lister() internalversion.CrossSecretBindingLister {
	return internalversion.NewCrossSecretBindingLister(f.Informer().GetIndexer())
}