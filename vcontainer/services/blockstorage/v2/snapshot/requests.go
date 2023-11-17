package snapshot

import (
	bsCm "github.com/vngcloud/vcontainer-sdk/vcontainer/services/blockstorage/v2"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/common"
)

type CreateOpts struct {
	Name        string
	Description string
	Permanently string
	RetainedDay uint64

	common.CommonOpts
	bsCm.BlockStorageV2Common
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}

// ******************************************************* Delete ******************************************************

type DeleteOpts struct {
	SnapshotID string
	common.CommonOpts
	bsCm.BlockStorageV2Common
}

func (s *DeleteOpts) GetSnapshotID() string {
	return s.SnapshotID
}
