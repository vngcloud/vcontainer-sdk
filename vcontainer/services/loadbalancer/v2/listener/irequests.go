package listener

type ICreateOptsBuilder interface {
	GetLoadBalancerID() string
	GetProjectID() string
	ToRequestBody() interface{}
}
