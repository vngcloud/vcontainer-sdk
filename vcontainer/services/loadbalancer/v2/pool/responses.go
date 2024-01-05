package pool

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2/pool/obj"

type CreateResponse struct {
	UUID string `json:"uuid"`
}

func (p *CreateResponse) ToPoolObject() *obj.Pool {
	return &obj.Pool{
		UUID: p.UUID,
	}
}
