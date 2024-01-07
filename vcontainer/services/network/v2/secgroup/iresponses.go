package secgroup

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/network/v2/secgroup/obj"

type ICreateResponse interface {
	ToSecgroupObject() *obj.Secgroup
}
