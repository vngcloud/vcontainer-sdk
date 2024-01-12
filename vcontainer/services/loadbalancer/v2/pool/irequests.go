package pool

type ICreateOptsBuilder interface {
	GetLoadBalancerID() string
	GetProjectID() string
	ToRequestBody() interface{}
}

type IListPoolsBasedLoadBalancerOptsBuilder interface {
	GetLoadBalancerID() string
	GetProjectID() string
}

type IDeleteOptsBuilder interface {
	GetLoadBalancerID() string
	GetProjectID() string
	GetPoolID() string
}

type IUpdatePoolMembersOptsBuilder interface {
	GetLoadBalancerID() string
	GetProjectID() string
	GetPoolID() string
	ToRequestBody() interface{}
}
