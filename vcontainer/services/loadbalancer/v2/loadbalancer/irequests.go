package loadbalancer

type ICreateOptsBuilder interface {
	ToRequestBody() interface{}
	GetProjectID() string
}

type IGetOptsBuilder interface {
	GetLoadBalancerID() string
	GetProjectID() string
}
