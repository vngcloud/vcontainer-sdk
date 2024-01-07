package secgroup

type ICreateOptsBuilder interface {
	GetProjectID() string
	ToRequestBody() interface{}
}
