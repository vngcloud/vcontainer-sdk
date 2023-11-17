package v1

import "github.com/vngcloud/vcontainer-sdk/client"

func getURL(pSc *client.ServiceClient, pProjectID string) string {
	return pSc.ServiceURL("projects", pProjectID, "detail")
}
