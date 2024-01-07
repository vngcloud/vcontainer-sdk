package secgroup

import (
	"github.com/vngcloud/vcontainer-sdk/client"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/network/v2/secgroup/obj"
)

func Create(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) (*obj.Secgroup, error) {
	response := NewCreateResponse()
	body := pOpts.ToRequestBody()
	_, err := pSc.Post(createURL(pSc, pOpts), &client.RequestOpts{
		JSONBody:     body,
		JSONResponse: response,
		OkCodes:      []int{201},
	})

	if err != nil {
		return nil, err
	}

	return response.ToSecgroupObject(), nil
}

func Delete(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) error {
	_, err := pSc.Delete(deleteURL(pSc, pOpts), &client.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}

func Get(pSc *client.ServiceClient, pOpts IGetOptsBuilder) (*obj.Secgroup, error) {
	response := NewGetResponse()
	_, err := pSc.Get(getURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToSecgroupObject(), nil
}
