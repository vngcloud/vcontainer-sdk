package pool

import (
	"github.com/vngcloud/vcontainer-sdk/vcontainer/objects"
)

type ICreateResponse interface {
	ToPoolObject() *objects.Pool
}

type IListPoolsBasedLoadBalancerResponse interface {
	ToListPoolObjects() []*objects.Pool
}
