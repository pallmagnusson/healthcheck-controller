// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/mbellgb/healthcheck-controller/pkg/apis/health/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HealthCheckLister helps list HealthChecks.
type HealthCheckLister interface {
	// List lists all HealthChecks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.HealthCheck, err error)
	// HealthChecks returns an object that can list and get HealthChecks.
	HealthChecks(namespace string) HealthCheckNamespaceLister
	HealthCheckListerExpansion
}

// healthCheckLister implements the HealthCheckLister interface.
type healthCheckLister struct {
	indexer cache.Indexer
}

// NewHealthCheckLister returns a new HealthCheckLister.
func NewHealthCheckLister(indexer cache.Indexer) HealthCheckLister {
	return &healthCheckLister{indexer: indexer}
}

// List lists all HealthChecks in the indexer.
func (s *healthCheckLister) List(selector labels.Selector) (ret []*v1alpha1.HealthCheck, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HealthCheck))
	})
	return ret, err
}

// HealthChecks returns an object that can list and get HealthChecks.
func (s *healthCheckLister) HealthChecks(namespace string) HealthCheckNamespaceLister {
	return healthCheckNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HealthCheckNamespaceLister helps list and get HealthChecks.
type HealthCheckNamespaceLister interface {
	// List lists all HealthChecks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.HealthCheck, err error)
	// Get retrieves the HealthCheck from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.HealthCheck, error)
	HealthCheckNamespaceListerExpansion
}

// healthCheckNamespaceLister implements the HealthCheckNamespaceLister
// interface.
type healthCheckNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HealthChecks in the indexer for a given namespace.
func (s healthCheckNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HealthCheck, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HealthCheck))
	})
	return ret, err
}

// Get retrieves the HealthCheck from the indexer for a given namespace and name.
func (s healthCheckNamespaceLister) Get(name string) (*v1alpha1.HealthCheck, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("healthcheck"), name)
	}
	return obj.(*v1alpha1.HealthCheck), nil
}
