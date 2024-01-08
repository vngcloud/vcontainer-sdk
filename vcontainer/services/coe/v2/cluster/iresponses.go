package cluster

import (
	"github.com/vngcloud/vcontainer-sdk/vcontainer/objects"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/coe/v2/cluster/obj"
)

// ******************************************** Response of GetCluster API *********************************************

type IGetResponse interface {
	ToClusterObject() *obj.Cluster
}

// *************************************** Response of Update Secgroup rule API ****************************************

type IUpdateSecgroupResponse interface {
	ToListClusterSecgroupRuleObjects() []*objects.ClusterSecgroup
}
