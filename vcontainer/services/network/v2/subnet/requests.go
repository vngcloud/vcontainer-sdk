package subnet

import (
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/common"
	lNetworkCommonV2 "github.com/vngcloud/vcontainer-sdk/vcontainer/services/network/v2"
)

// ****************************************************** GetOpts ******************************************************

type GetOpts struct {
	SubnetCommonOpts
	common.CommonOpts
	lNetworkCommonV2.NetworkV2Common
}
