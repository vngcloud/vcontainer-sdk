package loadbalancer

import (
	"github.com/vngcloud/vcontainer-sdk/vcontainer/objects"
)

// ******************************************************* MISC ********************************************************

type ResponseData struct {
	UUID               string `json:"uuid"`
	Name               string `json:"name"`
	DisplayStatus      string `json:"displayStatus"`
	Address            string `json:"address"`
	PrivateSubnetID    string `json:"privateSubnetId"`
	PrivateSubnetCidr  string `json:"privateSubnetCidr"`
	Type               string `json:"type"`
	DisplayType        string `json:"displayType"`
	LoadBalancerSchema string `json:"loadBalancerSchema"`
	PackageID          string `json:"packageId"`
	Description        string `json:"description"`
	Location           string `json:"location"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
	PackageInfo        struct {
		PackageID        string `json:"packageId"`
		ConnectionNumber int    `json:"connectionNumber"`
		DataTransfer     int    `json:"dataTransfer"`
		Name             string `json:"name"`
	} `json:"packageInfo"`
	ProgressStatus string `json:"progressStatus"`
}

// ******************************************* Response of CreateVolume API ********************************************

type CreateResponse struct {
	UUID string `json:"uuid"`
}

func (s *CreateResponse) ToLoadBalancerObject() *objects.LoadBalancer {
	if s == nil {
		return nil
	}

	return &objects.LoadBalancer{
		UUID: s.UUID,
	}
}

// ********************************************* Response of GetVolume API *********************************************

type GetResponse struct {
	Data ResponseData `json:"data"`
}

func (s *GetResponse) ToLoadBalancerObject() *objects.LoadBalancer {
	if s == nil {
		return nil
	}

	return &objects.LoadBalancer{
		UUID:     s.Data.UUID,
		Status:   s.Data.DisplayStatus,
		Address:  s.Data.Address,
		Name:     s.Data.Name,
		SubnetID: s.Data.PrivateSubnetID,
	}
}

// ******************************************* Response of ListBySubnetID API ******************************************

type ListBySubnetIDResponse struct {
	Data []struct {
		Address            string `json:"address"`
		CreatedAt          string `json:"createdAt"`
		Description        string `json:"description"`
		DisplayStatus      string `json:"displayStatus"`
		DisplayType        string `json:"displayType"`
		LoadBalancerSchema string `json:"loadBalancerSchema"`
		Name               string `json:"name"`
		PackageId          string `json:"packageId"`
		PrivateSubnetCidr  string `json:"privateSubnetCidr"`
		PrivateSubnetId    string `json:"privateSubnetId"`
		ProgressStatus     string `json:"progressStatus"`
		Type               string `json:"type"`
		UUID               string `json:"uuid"`
		UpdatedAt          string `json:"updatedAt"`
	}
}

func (s *ListBySubnetIDResponse) ToListLoadBalancerObjects() []*objects.LoadBalancer {
	if s == nil || s.Data == nil || len(s.Data) < 1 {
		return nil
	}

	var result []*objects.LoadBalancer
	for _, itemLb := range s.Data {
		result = append(result, &objects.LoadBalancer{
			UUID:     itemLb.UUID,
			Status:   itemLb.DisplayStatus,
			Address:  itemLb.Address,
			Name:     itemLb.Name,
			SubnetID: itemLb.PrivateSubnetId,
		})
	}

	return result
}

// *********************************************** Response of List API ************************************************

type ListResponse struct {
	ListData  []ResponseData `json:"listData"`
	Page      int            `json:"page"`
	PageSize  int            `json:"pageSize"`
	TotalPage int            `json:"totalPage"`
	TotalItem int            `json:"totalItem"`
}

func (s *ListResponse) ToListLoadBalancerObjects() []*objects.LoadBalancer {
	if s == nil {
		return nil
	}

	result := make([]*objects.LoadBalancer, len(s.ListData))
	for i := range s.ListData {
		result[i] = s.ToLoadBalancerObjectAt(i)
	}

	return result
}

func (s *ListResponse) ToLoadBalancerObjectAt(i int) *objects.LoadBalancer {
	if s == nil {
		return nil
	}

	return &objects.LoadBalancer{
		UUID:     s.ListData[i].UUID,
		Status:   s.ListData[i].DisplayStatus,
		Address:  s.ListData[i].Address,
		Name:     s.ListData[i].Name,
		SubnetID: s.ListData[i].PrivateSubnetID,
	}
}
