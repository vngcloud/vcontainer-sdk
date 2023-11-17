package volume

import (
	"fmt"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/pagination"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/blockstorage/v2/volume/obj"
)

// ********************************************* List Volume API Response **********************************************

type (
	ListResponse struct {
		Page      int64    `json:"page"`
		PageSize  int64    `json:"pageSize"`
		TotalPage int64    `json:"totalPage"`
		TotalItem int64    `json:"totalItem"`
		ListData  []Common `json:"listData"`

		pagination.IPage
	}

	Common struct {
		UUID               string   `json:"uuid"`
		ProjectID          string   `json:"projectId"`
		Name               string   `json:"name"`
		Size               uint64   `json:"size"`
		Status             string   `json:"status"`
		VolumeTypeID       string   `json:"volumeTypeId"`
		VolumeTypeZoneName string   `json:"volumeTypeZoneName"`
		IOPS               string   `json:"iops"`
		ServerID           *string  `json:"serverId"`
		CreatedAt          string   `json:"createdAt"`
		UpdatedAt          *string  `json:"updatedAt"`
		Bootable           bool     `json:"bootable"`
		EncryptionType     *string  `json:"encryptionType"`
		BootIndex          int      `json:"bootIndex"`
		MultiAttach        bool     `json:"multiAttach"`
		ServerIDList       []string `json:"serverIdList"`
		Location           *string  `json:"location"`
		Product            string   `json:"product"`
		PersistentVolume   bool     `json:"persistentVolume"`
	}
)

func (s *ListResponse) IsEmpty() (bool, error) {
	if len(s.ListData) < 1 {
		return true, nil
	}

	return false, nil
}

func (s *ListResponse) NextPageURL(pOpts pagination.IPageOpts) (string, error) {
	currentPage := s.Page
	totalPages := s.TotalPage
	defaultOpts := pOpts.(*ListOpts)

	if totalPages > currentPage {
		query, err := pOpts.ToListQueryWithParams(&map[string]interface{}{
			"page": defaultOpts.Page + 1,
			"size": defaultOpts.Size,
			"name": defaultOpts.Name,
		})
		if err != nil {
			return "", err
		}

		return query, nil
	}

	return "", nil
}

func (s *ListResponse) GetBody() interface{} {
	return s
}

func (s *ListResponse) ToListVolumeObjects() []*obj.Volume {
	var volumes []*obj.Volume
	for i, _ := range s.ListData {
		vol := s.ToVolumeObject(i)
		if vol != nil {
			volumes = append(volumes, vol)
		}
	}

	return volumes
}

func (s *ListResponse) NextPage() string {
	currentPage := s.Page
	totalPages := s.TotalPage
	if totalPages > currentPage {
		return fmt.Sprintf("%d", currentPage+1)
	}

	return ""
}

func (s *ListResponse) ToVolumeObject(pIdx int) *obj.Volume {
	if s == nil {
		return nil
	}

	if pIdx >= 0 && pIdx < len(s.ListData) {
		vol := s.ListData[pIdx]
		return &obj.Volume{
			VolumeId:  vol.UUID,
			Name:      vol.Name,
			Size:      vol.Size,
			Status:    vol.Status,
			CreatedAt: vol.CreatedAt,
			UpdatedAt: vol.UpdatedAt,
			VmId:      vol.ServerID,
		}
	}

	return nil
}

// ********************************************* Create Volume API Response ********************************************

type (
	CreateResponse struct {
		Data Common `json:"data"`
	}
)

func (s *CreateResponse) ToVolumeObject() *obj.Volume {
	if s == nil {
		return nil
	}

	return &obj.Volume{
		VolumeId:  s.Data.UUID,
		Name:      s.Data.Name,
		Size:      s.Data.Size,
		Status:    s.Data.Status,
		CreatedAt: s.Data.CreatedAt,
		UpdatedAt: s.Data.UpdatedAt,
		VmId:      s.Data.ServerID,
	}
}

// ********************************************** Get Volume API Response **********************************************

type (
	GetResponse struct {
		Data Common `json:"data"`
	}
)

func (s *GetResponse) ToVolumeObject() *obj.Volume {
	if s == nil {
		return nil
	}

	return &obj.Volume{
		VolumeId:         s.Data.UUID,
		VolumeTypeID:     s.Data.VolumeTypeID,
		Name:             s.Data.Name,
		Size:             s.Data.Size,
		Status:           s.Data.Status,
		CreatedAt:        s.Data.CreatedAt,
		UpdatedAt:        s.Data.UpdatedAt,
		VmId:             s.Data.ServerID,
		PersistentVolume: s.Data.PersistentVolume,
	}
}
