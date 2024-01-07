package secgroup_rule

import (
	"encoding/json"
	"github.com/vngcloud/vcontainer-sdk/client"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/network/v2/extensions/secgroup_rule/obj"
	"k8s.io/klog/v2"
)

func Create(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) (*obj.SecgroupRule, error) {
	response := NewCreateResponse()
	body := pOpts.ToRequestBody()
	reqRes, err := pSc.Post(createURL(pSc, pOpts), &client.RequestOpts{
		JSONBody:     body,
		JSONResponse: response,
		OkCodes:      []int{201},
	})

	if err != nil {
		result := make(map[string]interface{})
		err2 := json.Unmarshal(reqRes.Bytes(), &result)
		if err2 == nil {
			message, _ := result["message"].(string)
			klog.Errorf("Create secgroup rule failed: %+v", message)
		}

		return nil, err
	}

	return response.ToSecgroupRuleObject(), nil
}
