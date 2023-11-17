package volume_type

import "github.com/vngcloud/vcontainer-sdk/client"

func getVolumeTypeURL(pSc *client.ServiceClient, pOpts IGetVolumeTypeOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.ExtractProjectID(),
		"volume_types",
		pOpts.ExtractVolumeTypeID())
}
