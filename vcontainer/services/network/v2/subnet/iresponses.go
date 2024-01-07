package subnet

import (
	"github.com/vngcloud/vcontainer-sdk/vcontainer/objects"
)

type IGetResponse interface {
	ToSubnetObject() *objects.Subnet
}
