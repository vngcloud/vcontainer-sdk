package volume_type

import (
	"github.com/vngcloud/vcontainer-sdk/client"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/blockstorage/v1/volume_type/obj"
)

func Get(pSc *client.ServiceClient, pOpts IGetVolumeTypeOptsBuilder) (*obj.VolumeType, error) {
	response := NewGetVolumeTypeResponses()
	_, err := pSc.Get(
		getVolumeTypeURL(pSc, pOpts),
		&client.RequestOpts{
			JSONResponse: response,
			OkCodes:      []int{200}})

	if err != nil {
		return nil, err
	}

	return response.ToVolumeTypeObject()
}
