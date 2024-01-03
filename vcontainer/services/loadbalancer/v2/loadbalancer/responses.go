package loadbalancer

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2/loadbalancer/obj"

// ******************************************* Response of CreateVolume API ********************************************

type CreateResponse struct {
	UUID string `json:"uuid"`
}

func (s *CreateResponse) ToLoadBalancerObject() *obj.LoadBalancer {
	if s == nil {
		return nil
	}

	return &obj.LoadBalancer{
		UUID: s.UUID,
	}
}
