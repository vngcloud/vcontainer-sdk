package listener

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2/listener/obj"

type CreateResponse struct {
}

func (p *CreateResponse) ToListenerObject() *obj.Listener {
	return nil
}
