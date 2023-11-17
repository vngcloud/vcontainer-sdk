package v1

import (
	"github.com/vngcloud/vcontainer-sdk/client"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/portal/v1/obj"
)

// Get gets the under-user information from the portal information
func Get(pSc *client.ServiceClient, pProjectID string) (*obj.Portal, error) {
	portal := new(obj.Portal)
	_, err := pSc.Get(getURL(pSc, pProjectID), &client.RequestOpts{
		JSONResponse: portal,
		OkCodes:      []int{200},
	})

	return portal, err
}
