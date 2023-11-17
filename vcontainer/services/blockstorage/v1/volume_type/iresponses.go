package volume_type

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/blockstorage/v1/volume_type/obj"

type IGetVolumeTypeResponse interface {
	ToVolumeTypeObject() (*obj.VolumeType, error)
}
