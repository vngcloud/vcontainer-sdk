package volume

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/blockstorage/v2/volume/obj"

// ******************************************** Response of CreateVolume API *******************************************

type ICreateResponse interface {
	ToVolumeObject() *obj.Volume
}

// ******************************************** Response of GetVolume API **********************************************

type IGetResponse interface {
	ToVolumeObject() *obj.Volume
}

// ******************************************** Response of ListVolume API *********************************************

type IListResponse interface {
	ToListVolumeObjects() []*obj.Volume
	ToVolumeObject(pIdx int) *obj.Volume
	NextPage() string
}
