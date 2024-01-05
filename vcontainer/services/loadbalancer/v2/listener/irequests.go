package listener

type ICreateOptsBuilder interface {
	GetLoadBalancerID() string
	GetProjectID() string
	ToRequestBody() interface{}
}

type IGetBasedLoadBalancerOptsBuilder interface {
	GetLoadBalancerID() string
	GetProjectID() string
}
