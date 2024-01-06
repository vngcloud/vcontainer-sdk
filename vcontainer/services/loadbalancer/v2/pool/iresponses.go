package pool

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2/pool/obj"

type ICreateResponse interface {
	ToPoolObject() *obj.Pool
}

type IListPoolsBasedLoadBalancerResponse interface {
	ToPoolObjectAt(i int) *obj.Pool
	ToListPoolObjects() []*obj.Pool
}
