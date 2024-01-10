package loadbalancer

import (
	"github.com/vngcloud/vcontainer-sdk/vcontainer/objects"
)

// ***************************************** Response of CreateLoadBalancer API ****************************************

type ICreateResponse interface {
	ToLoadBalancerObject() *objects.LoadBalancer
}

// ****************************************** Response of GetLoadBalancer API ******************************************

type IGetResponse interface {
	ToLoadBalancerObject() *objects.LoadBalancer
}

// ****************************************** Response of ListBySubnetID API *******************************************

type IListBySubnetIDResponse interface {
	ToListLoadBalancerObjects() []*objects.LoadBalancer
}

// *********************************************** Response if List API ************************************************

type IListResponse interface {
	ToListLoadBalancerObjects() []*objects.LoadBalancer
	ToLoadBalancerObjectAt(i int) *objects.LoadBalancer
}
