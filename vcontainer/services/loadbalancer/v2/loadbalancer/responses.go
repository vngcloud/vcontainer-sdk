package loadbalancer

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2/loadbalancer/obj"

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
	Data struct {
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
	} `json:"data"`
}

func (s *GetResponse) ToLoadBalancerObject() *obj.LoadBalancer {
	if s == nil {
		return nil
	}

	return &obj.LoadBalancer{
		UUID:   s.Data.UUID,
		Status: s.Data.DisplayStatus,
	}
}
