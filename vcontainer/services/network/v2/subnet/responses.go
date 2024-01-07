package subnet

import "github.com/vngcloud/vcontainer-sdk/vcontainer/objects"

type GetResponse struct {
	Subnet                 string  `json:"subnet"`
	CreatedAt              string  `json:"createdAt"`
	UpdatedAt              *string `json:"updatedAt,omitempty"`
	DeletedAt              *string `json:"deletedAt,omitempty"`
	Status                 string  `json:"status"`
	Cidr                   string  `json:"cidr"`
	NetworkUUID            string  `json:"networkUuid"`
	RouteTableUUID         *string `json:"routeTableUuid,omitempty"`
	Name                   string  `json:"name"`
	InterfaceAclPolicyUUID *string `json:"interfaceAclPolicyUuid,omitempty"`
}

func (s *GetResponse) ToSubnetObject() *objects.Subnet {
	if s == nil {
		return nil
	}

	return &objects.Subnet{
		UUID:      s.Subnet,
		NetworkID: s.NetworkUUID,
		CIRD:      s.Cidr,
	}
}
