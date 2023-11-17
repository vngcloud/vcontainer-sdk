package volume

type IGetVolumeOptsBuilder interface{}

type IGetOptsBuilder interface {
	GetVolumeID() string
	GetProjectID() string
}

type IDeleteOptsBuilder interface {
	GetProjectID() string
	GetVolumeID() string
}

type ICreateOptsBuilder interface {
	ToRequestBody() interface{}
	GetProjectID() string
}

type IListOptsBuilder interface {
	ToListQuery() (string, error)
	ToListQueryWithParams(*map[string]interface{}) (string, error)
	GetProjectID() string
}
