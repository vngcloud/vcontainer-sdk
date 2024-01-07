package secgroup

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/network/v2/secgroup/obj"

type CreateResponse struct {
	Data struct {
		ID           int     `json:"id"`
		UUID         string  `json:"uuid"`
		CreatedAt    string  `json:"createdAt"`
		DeletedAt    *string `json:"deletedAt,omitempty"`
		Status       string  `json:"status"`
		SecgroupId   int     `json:"secgroupId"`
		SecgroupName string  `json:"secgroupName"`
		ProjectUuid  string  `json:"projectUuid"`
		Description  string  `json:"description"`
		UpdatedAt    *string `json:"updatedAt,omitempty"`
		IsSystem     bool    `json:"isSystem"`
		Type         string  `json:"type"`
	} `json:"data"`
}

func (s *CreateResponse) ToSecgroupObject() *obj.Secgroup {
	if s == nil {
		return nil
	}

	return &obj.Secgroup{
		UUID:        s.Data.UUID,
		Name:        s.Data.SecgroupName,
		Description: s.Data.Description,
		Status:      s.Data.Status,
	}
}
