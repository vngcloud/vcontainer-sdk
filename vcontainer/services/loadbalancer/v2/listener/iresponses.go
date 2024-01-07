package listener

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2/listener/obj"

type ICreateResponse interface {
	ToListenerObject() *obj.Listener
}

type IGetBasedLoadBalancerResponse interface {
	ToListListenerObject() []*obj.Listener
}
