package volume_actions

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/blockstorage/v2/extensions/volume_actions/obj"

type IResizeResponse interface {
	ToResizeVolumeObject() *obj.ResizeVolume
}

type IMappingResponse interface {
	ToMappingVolumeObject() *obj.MappingVolume
}
