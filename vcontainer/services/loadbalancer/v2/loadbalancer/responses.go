package loadbalancer

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2/loadbalancer/obj"

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

func (s *CreateResponse) ToLoadBalancerObject() *obj.LoadBalancer {
	if s == nil {
		return nil
	}

	return &obj.LoadBalancer{
		UUID: s.UUID,
	}
}

// ********************************************* Response of GetVolume API *********************************************

type GetResponse struct {
	Data ResponseData `json:"data"`
}

func (s *GetResponse) ToLoadBalancerObject() *obj.LoadBalancer {
	if s == nil {
		return nil
	}

	return &obj.LoadBalancer{
		UUID:     s.Data.UUID,
		Status:   s.Data.DisplayStatus,
		Address:  s.Data.Address,
		Name:     s.Data.Name,
		SubnetID: s.Data.PrivateSubnetID,
	}
}

// ******************************************* Response of ListBySubnetID API ******************************************

type ListBySubnetIDResponse struct {
	Data []ResponseData `json:"data"`
}

func (s *ListBySubnetIDResponse) ToListLoadBalancerObjects() []*obj.LoadBalancer {
	if s == nil {
		return nil
	}

	result := make([]*obj.LoadBalancer, len(s.Data))
	for i := range s.Data {
		result[i] = s.ToLoadBalancerObjectAt(i)
	}

	return result
}

func (s *ListBySubnetIDResponse) ToLoadBalancerObjectAt(i int) *obj.LoadBalancer {
	if s == nil {
		return nil
	}

	return &obj.LoadBalancer{
		UUID:     s.Data[i].UUID,
		Status:   s.Data[i].DisplayStatus,
		Address:  s.Data[i].Address,
		Name:     s.Data[i].Name,
		SubnetID: s.Data[i].PrivateSubnetID,
	}
}

// *********************************************** Response of List API ************************************************

type ListResponse struct {
	ListData  []ResponseData `json:"listData"`
	Page      int            `json:"page"`
	PageSize  int            `json:"pageSize"`
	TotalPage int            `json:"totalPage"`
	TotalItem int            `json:"totalItem"`
}

func (s *ListResponse) ToListLoadBalancerObjects() []*obj.LoadBalancer {
	if s == nil {
		return nil
	}

	result := make([]*obj.LoadBalancer, len(s.ListData))
	for i := range s.ListData {
		result[i] = s.ToLoadBalancerObjectAt(i)
	}

	return result
}

func (s *ListResponse) ToLoadBalancerObjectAt(i int) *obj.LoadBalancer {
	if s == nil {
		return nil
	}

	return &obj.LoadBalancer{
		UUID:     s.ListData[i].UUID,
		Status:   s.ListData[i].DisplayStatus,
		Address:  s.ListData[i].Address,
		Name:     s.ListData[i].Name,
		SubnetID: s.ListData[i].PrivateSubnetID,
	}
}
