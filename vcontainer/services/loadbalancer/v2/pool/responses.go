package pool

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2/pool/obj"

type CreateResponse struct {
}

func (p *CreateResponse) ToPoolObject() *obj.Pool {
	return nil
}
