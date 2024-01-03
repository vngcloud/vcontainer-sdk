package loadbalancer

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2/loadbalancer/obj"

// ******************************************** Response of CreateVolume API *******************************************

type ICreateResponse interface {
	ToLoadBalancerObject() *obj.LoadBalancer
}
