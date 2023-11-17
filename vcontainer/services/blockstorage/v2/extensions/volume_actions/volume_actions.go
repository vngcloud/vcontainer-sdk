package volume_actions

import (
	"github.com/vngcloud/vcontainer-sdk/client"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/blockstorage/v2/extensions/volume_actions/obj"
)

func Resize(pSc *client.ServiceClient, pOpts IResizeOptsBuilder) (*obj.ResizeVolume, error) {
	response := NewResizeResponse()
	_, err := pSc.Put(resizeURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		JSONBody:     pOpts.ToRequestBody(),
		OkCodes:      []int{202},
	})

	if err != nil {
		return nil, err
	}

	return response.ToResizeVolumeObject(), nil
}

func GetMappingVolume(pSc *client.ServiceClient, pOpts IMappingOptsBuilder) (*obj.MappingVolume, error) {
	response := NewBackendVolumeResponse()
	_, err := pSc.Get(mappingURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToMappingVolumeObject(), nil
}
