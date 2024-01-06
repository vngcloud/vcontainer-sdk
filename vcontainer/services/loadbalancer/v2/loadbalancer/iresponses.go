package loadbalancer

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2/loadbalancer/obj"

// ***************************************** Response of CreateLoadBalancer API ****************************************

type ICreateResponse interface {
	ToLoadBalancerObject() *obj.LoadBalancer
}

// ****************************************** Response of GetLoadBalancer API ******************************************

type IGetResponse interface {
	ToLoadBalancerObject() *obj.LoadBalancer
}

// ****************************************** Response of ListBySubnetID API *******************************************

type IListBySubnetIDResponse interface {
	ToListLoadBalancerObjects() []*obj.LoadBalancer
	ToLoadBalancerObjectAt(i int) *obj.LoadBalancer
}
