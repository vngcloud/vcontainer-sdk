package secgroup

import (
	"github.com/vngcloud/vcontainer-sdk/vcontainer/objects"
)

type ICreateResponse interface {
	ToSecgroupObject() *objects.Secgroup
}

type IGetResponse interface {
	ToSecgroupObject() *objects.Secgroup
}

type IListResponse interface {
	ToListSecgroupObjects() []*objects.Secgroup
}
