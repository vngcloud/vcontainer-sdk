package snapshot

import (
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/blockstorage/v1/snapshot/obj"
)

// ******************************************* Response of List Snapshot API *******************************************

type IListResponse interface {
	ToListSnapshotObjects() []*obj.Snapshot
	ToSnapshotObject(pIdx int) *obj.Snapshot
	NextPage() string
}
