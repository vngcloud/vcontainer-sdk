package cluster

import (
	"github.com/vngcloud/vcontainer-sdk/client"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/coe/v2/cluster/obj"
)

func Get(pSc *client.ServiceClient, pOpts IGetOptsBuilder) (*obj.Cluster, error) {
	response := NewGetResponse()
	_, err := pSc.Get(getURL(pSc, pOpts), &client.RequestOpts{
		OkCodes:      []int{200},
		JSONResponse: response,
	})

	if err != nil {
		return nil, err
	}

	return response.ToClusterObject(), nil
}
