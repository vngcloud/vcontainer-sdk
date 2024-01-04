package pool

type ICreateOptsBuilder interface {
	GetLoadBalancerID() string
	GetProjectID() string
	ToRequestBody() interface{}
}
