package snapshot

import "github.com/vngcloud/vcontainer-sdk/client"

// listURL generates the URL to get all the Snapshot.
func listURL(pSc *client.ServiceClient, pOpts IListOptsBuilder) string {
	return pSc.ServiceURL("snapshot-volumes")
}
