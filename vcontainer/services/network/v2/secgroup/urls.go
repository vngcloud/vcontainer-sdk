package secgroup

import "github.com/vngcloud/vcontainer-sdk/client"

func createURL(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) string {
	return pSc.ServiceURL(pOpts.GetProjectID(), "secgroups")
}

func deleteURL(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) string {
	return pSc.ServiceURL(pOpts.GetProjectID(), "secgroups", pOpts.GetSecgroupUUID())
}

func getURL(pSc *client.ServiceClient, pOpts IGetOptsBuilder) string {
	return pSc.ServiceURL(pOpts.GetProjectID(), "secgroups", pOpts.GetSecgroupUUID())
}
