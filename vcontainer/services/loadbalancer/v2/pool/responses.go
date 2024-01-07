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

// **************************************** ListPoolsBasedLoadBalancerResponse *****************************************

type ListPoolsBasedLoadBalancerResponse struct {
	Data []struct {
		UUID              string  `json:"uuid"`
		Name              string  `json:"name"`
		Protocol          string  `json:"protocol"`
		Description       *string `json:"description"`
		LoadBalanceMethod string  `json:"loadBalanceMethod"`
		DisplayStatus     string  `json:"displayStatus"`
		CreatedAt         string  `json:"createdAt"`
		UpdatedAt         string  `json:"updatedAt"`
		Stickiness        bool    `json:"stickiness"`
		TLSEncryption     bool    `json:"tlsEncryption"`
		ProgressStatus    string  `json:"progressStatus"`
	} `json:"data"`
}

func (s *ListPoolsBasedLoadBalancerResponse) ToListPoolObjects() []*obj.Pool {
	var pools []*obj.Pool
	for _, item := range s.Data {
		pools = append(pools, &obj.Pool{
			UUID:              item.UUID,
			Name:              item.Name,
			Protocol:          item.Protocol,
			Description:       *item.Description,
			LoadBalanceMethod: item.LoadBalanceMethod,
			Status:            item.DisplayStatus,
			Stickiness:        item.Stickiness,
			TLSEncryption:     item.TLSEncryption,
		})
	}
	return pools
}
