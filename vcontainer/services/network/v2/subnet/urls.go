package subnet

import "github.com/vngcloud/vcontainer-sdk/client"

func getURL(pSc *client.ServiceClient, pOpts IGetOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"networks",
		pOpts.GetNetworkUUID(),
		"subnets",
		pOpts.GetSubnetUUID())
}
