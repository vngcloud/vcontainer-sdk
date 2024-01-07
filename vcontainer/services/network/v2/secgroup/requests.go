package secgroup

import (
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/common"
	lSecgroupCommonV2 "github.com/vngcloud/vcontainer-sdk/vcontainer/services/network/v2"
)

type CreateOpts struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	common.CommonOpts
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}

type DeleteOpts struct {
	lSecgroupCommonV2.SecgroupV2Common
	common.CommonOpts
}
