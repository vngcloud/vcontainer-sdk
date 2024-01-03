package loadbalancer

type ICreateOptsBuilder interface {
	ToRequestBody() interface{}
	GetProjectID() string
}
