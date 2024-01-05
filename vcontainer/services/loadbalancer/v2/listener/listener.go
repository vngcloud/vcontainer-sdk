package listener

import (
	"github.com/vngcloud/vcontainer-sdk/client"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2/listener/obj"
)

func Create(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) (*obj.Listener, error) {
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

	return response.ToListenerObject(), nil
}

func GetBasedLoadBalancer(pSc *client.ServiceClient, pOpts IGetOptsBuilder) ([]*obj.Listener, error) {
	response := NewGetResponse()
	_, err := pSc.Get(getBasedLoadBalancerURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToListListenerObject(), nil
}
