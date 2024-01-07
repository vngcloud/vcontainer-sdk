package secgroup_rule

import (
	"github.com/vngcloud/vcontainer-sdk/client"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/network/v2/extensions/secgroup_rule/obj"
	"k8s.io/klog/v2"
)

func Create(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) (*obj.SecgroupRule, error) {
	response := NewCreateResponse()
	body := pOpts.ToRequestBody()
	respReq, err := pSc.Post(createURL(pSc, pOpts), &client.RequestOpts{
		JSONBody:     body,
		JSONResponse: response,
		OkCodes:      []int{201},
	})

	if err != nil {
		klog.Errorf("Create secgroup rule failed: %v", respReq)
		return nil, err
	}

	return response.ToSecgroupRuleObject(), nil
}
