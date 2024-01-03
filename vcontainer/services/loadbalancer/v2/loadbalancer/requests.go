package loadbalancer

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/common"

// ************************************************* CreateOptsBuilder *************************************************

const (
	// For internet facing
	CreateOptsSchemeOptInternet CreateOptsSchemeOpt = "Internet"
	CreateOptsSchemeOptInternal CreateOptsSchemeOpt = "Internal"

	// For lb type
	CreateOptsTypeOptLayer4 CreateOptsTypeOpt = "Layer 4"
	CreateOptsTypeOptLayer7 CreateOptsTypeOpt = "Layer 7"
)

type (
	CreateOptsSchemeOpt string
	CreateOptsTypeOpt   string
)

type CreateOpts struct {
	Name      string              `json:"name"`
	PackageID string              `json:"packageId"`
	Scheme    CreateOptsSchemeOpt `json:"scheme"`
	SubnetID  string              `json:"subnetId"`
	Type      CreateOptsTypeOpt   `json:"type"`
	common.CommonOpts
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}
