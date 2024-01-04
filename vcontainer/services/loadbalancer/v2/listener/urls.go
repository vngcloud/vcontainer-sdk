package listener

import "github.com/vngcloud/vcontainer-sdk/client"

func createURL(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		pOpts.GetLoadBalancerID(),
		"listeners")
}
