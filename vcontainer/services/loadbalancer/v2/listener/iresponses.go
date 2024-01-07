package listener

import (
	"github.com/vngcloud/vcontainer-sdk/vcontainer/objects"
)

type ICreateResponse interface {
	ToListenerObject() *objects.Listener
}

type IGetBasedLoadBalancerResponse interface {
	ToListListenerObject() []*objects.Listener
}
