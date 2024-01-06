package pool

import (
	"github.com/vngcloud/vcontainer-sdk/client"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2/pool/obj"
)

func Create(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) (*obj.Pool, error) {
	response := NewCreateResponse()
	body := pOpts.ToRequestBody()
	_, err := pSc.Post(createURL(pSc, pOpts), &client.RequestOpts{
		JSONBody:     body,
		JSONResponse: response,
		OkCodes:      []int{202},
	})

	if err != nil {
		return nil, err
	}

	return response.ToPoolObject(), nil
}

func ListPoolsBasedLoadBalancer(pSc *client.ServiceClient, pOpts IListPoolsBasedLoadBalancerOptsBuilder) ([]*obj.Pool, error) {
	response := NewListPoolsBasedLoadBalancerResponse()
	_, err := pSc.Get(listPoolsBasedLoadBalancerURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToListPoolObjects(), nil
}
