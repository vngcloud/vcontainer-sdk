package volume_type

import (
	"fmt"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/blockstorage/v1/volume_type/obj"
)

type GetVolumeTypeResponse struct {
	ErrorMsg    *string     `json:"errorMsg"`
	ErrorCode   *string     `json:"errorCode"`
	Success     bool        `json:"success"`
	Extra       interface{} `json:"extra"`
	VolumeTypes []struct {
		Iops       int    `json:"iops"`
		Name       string `json:"name"`
		ThroughPut int    `json:"throughPut"`
		Id         string `json:"id"`
		MinSize    int    `json:"minSize"`
		MaxSize    int    `json:"maxSize"`
		ZoneID     string `json:"zoneId"`
	} `json:"volumeTypes"`
}

func (s *GetVolumeTypeResponse) ToVolumeTypeObject() (*obj.VolumeType, error) {
	if len(s.VolumeTypes) > 0 {
		firstItem := s.VolumeTypes[0]
		return &obj.VolumeType{
			VolumeTypeID: firstItem.Id,
			MaxSize:      int64(firstItem.MaxSize),
			MinSize:      int64(firstItem.MinSize),
		}, nil
	}
	return nil, fmt.Errorf("no volume type found")
}
