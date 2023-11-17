package volume_attach

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/common"

// ******************************************** Create Attach Volume Request *******************************************

type CreateOpts struct {
	AttachCommonOpts
	common.CommonOpts
}

// ***************************************** Delete Volume Attachment Requests *****************************************

type DeleteOpts struct {
	AttachCommonOpts
	common.CommonOpts
}
