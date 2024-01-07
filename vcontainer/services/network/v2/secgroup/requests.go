package secgroup

import (
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/common"
)

type CreateOpts struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	common.CommonOpts
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}
