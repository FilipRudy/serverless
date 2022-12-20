// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kyma-project/kyma/components/compass-runtime-agent/pkg/client/clientset/versioned/typed/compass/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCompassV1alpha1 struct {
	*testing.Fake
}

func (c *FakeCompassV1alpha1) CompassConnections() v1alpha1.CompassConnectionInterface {
	return &FakeCompassConnections{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCompassV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
