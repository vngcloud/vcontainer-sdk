package snapshot

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/blockstorage/v2/snapshot/obj"

type ICreateResponse interface {
	ToSnapshotObject() *obj.Snapshot
}
