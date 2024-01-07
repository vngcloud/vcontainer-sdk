package secgroup_rule

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/network/v2/extensions/secgroup_rule/obj"

type ICreateResponse interface {
	ToSecgroupRuleObject() *obj.SecgroupRule
}
