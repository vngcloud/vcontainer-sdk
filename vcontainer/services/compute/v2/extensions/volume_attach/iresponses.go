package volume_attach

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/compute/v2/extensions/volume_attach/obj"

// ***************************************** Response of Create Attachment API *****************************************

type ICreateResponse interface {
	ToVolumeAttachObject() *obj.VolumeAttach
}

type IDeleteResponse interface {
	ToVolumeAttachObject() *obj.VolumeAttach
}
