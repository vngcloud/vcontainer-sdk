package loadbalancer

func NewCreateResponse() ICreateResponse {
	return &CreateResponse{}
}

func NewGetResponse() IGetResponse {
	return &GetResponse{}
}

func NewListBySubnetIDResponse() []*ListBySubnetIDResponse {
	var response []*ListBySubnetIDResponse
	return response
}

func NewListResponse() IListResponse {
	return &ListResponse{}
}
