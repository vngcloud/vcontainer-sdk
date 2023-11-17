package server

import (
	"github.com/vngcloud/vcontainer-sdk/client"
	serverObj "github.com/vngcloud/vcontainer-sdk/vcontainer/services/compute/v2/server/obj"
)

func Get(sc *client.ServiceClient, opts GetOptsBuilder) (*serverObj.Server, error) {
	var response *serverObj.Server
	_, err := sc.Get(getServerURL(sc, opts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}
