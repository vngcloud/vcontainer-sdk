package secgroup

import (
	"encoding/json"
	"github.com/vngcloud/vcontainer-sdk/client"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/objects"
	"strings"
)

func Create(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) (*objects.Secgroup, error) {
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
			if message, _ := result["message"].(string); strings.Contains(strings.ToLower(strings.TrimSpace(message)), "name of security group already exist") {
				return nil, NewErrNameDuplicate("", "name is already used")
			}
		}

		return nil, err
	}

	return response.ToSecgroupObject(), nil
}

func Delete(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) error {
	_, err := pSc.Delete(deleteURL(pSc, pOpts), &client.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}

func Get(pSc *client.ServiceClient, pOpts IGetOptsBuilder) (*objects.Secgroup, error) {
	response := NewGetResponse()
	_, err := pSc.Get(getURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToSecgroupObject(), nil
}
