package cluster

type IGetClusterOptsBuilder interface{}

type IGetOptsBuilder interface {
	GetClusterID() string
	GetProjectID() string
}
