package listener

func NewCreateResponse() ICreateResponse {
	return &CreateResponse{}
}

func NewGetResponse() IGetResponse {
	return &GetResponse{}
}
