package volume_attach

import (
	"fmt"
	"github.com/vngcloud/vcontainer-sdk/client"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/compute/v2/extensions/volume_attach/obj"
)

// Create creates a volume attachment.
func Create(sc *client.ServiceClient, opts ICreateOptsBuilder) (*obj.VolumeAttach, error) {
	response := NewCreateResponse()
	_, err := sc.Put(createURL(sc, opts), &client.RequestOpts{
		OkCodes:      []int{202},
		JSONResponse: response,
		JSONBody:     opts.ToRequestBody(),
	})

	if err != nil {
		if err.Error() == "This volume is available" {
			return nil, NewErrAttachNotFound(fmt.Sprintf("Volume %s is available", opts.GetVolumeID()))
		}
		return nil, err
	}

	return response.ToVolumeAttachObject(), nil
}

// Delete deletes a volume attachment.
func Delete(sc *client.ServiceClient, opts IDeleteOptsBuilder) (*obj.VolumeAttach, error) {
	response := NewDeleteResponse()
	_, err := sc.Put(deleteURL(sc, opts), &client.RequestOpts{
		OkCodes:      []int{202},
		JSONResponse: response,
		JSONBody:     opts.ToRequestBody(),
	})

	if err != nil {
		return nil, err
	}

	return response.ToVolumeAttachObject(), nil
}
