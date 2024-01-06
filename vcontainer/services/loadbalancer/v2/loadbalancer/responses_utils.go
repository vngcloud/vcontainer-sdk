package loadbalancer

func NewCreateResponse() ICreateResponse {
	return &CreateResponse{}
}

func NewGetResponse() IGetResponse {
	return &GetResponse{}
}

func NewListBySubnetIDResponse() IListBySubnetIDResponse {
	return &ListBySubnetIDResponse{}
}
