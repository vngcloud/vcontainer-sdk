package volume_attach

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/compute/v2/extensions/volume_attach/obj"

// ***************************************** Create Volume Attachment Response *****************************************

type CreateResponse struct {
}

func (s *CreateResponse) ToVolumeAttachObject() *obj.VolumeAttach {
	return nil
}

// ***************************************** Delete Volume Attachment Response *****************************************

type DeleteResponse struct{}

func (s *DeleteResponse) ToVolumeAttachObject() *obj.VolumeAttach {
	return nil
}
